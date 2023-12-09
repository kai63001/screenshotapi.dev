package api

import (
	module "backend/module"
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func TakeScreenshotByAPI(c echo.Context, db dbx.Builder, mongo *mongo.Collection, rdb *redis.Client) error {
	//get all query params
	url := c.QueryString()

	access_key := c.QueryParam("access_key")
	if access_key == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "access_key is required",
		})
	}

	//get user_id from access_key
	userData := module.UserForKey{}
	errAccessKey := db.Select("user_id").From("access_keys").Where(dbx.NewExp("access_key = {:access_key}", dbx.Params{"access_key": access_key})).One(&userData)
	if errAccessKey != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "access_key is invalid",
		})
	}

	custom := c.QueryParam("custom")
	customData := module.CustomSet{}
	if custom != "" {
		errCustom := db.Select("id", "name", "user_id", "css", "javascript", "cookies", "localStorage", "user_agent", "headers", "bucket_endpoint", "bucket_default", "bucket_access_key", "bucket_secret_key").From("custom_sets").Where(dbx.NewExp("name = {:name} and user_id = {:user_id}", dbx.Params{"name": custom, "user_id": userData.UserId})).One(&customData)
		if errCustom != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status":  "error",
				"message": "custom is invalid",
			})
		}
	}

	quotaData := module.GetQuotaScreenshot{}
	errCheckQuota := db.
		Select("screenshot_usage.screenshots_taken", "subscription_plans.name", "subscription_plans.included_screenshots", "subscription_plans.rate_limit_per_minute").
		From("screenshot_usage").
		InnerJoin("users", dbx.NewExp("users.id = screenshot_usage.user_id")).
		InnerJoin("subscription_plans", dbx.NewExp("subscription_plans.id = users.subscription_plan")).
		Where(dbx.NewExp("screenshot_usage.user_id = {:user_id}", dbx.Params{"user_id": userData.UserId})).
		One(&quotaData)
	if errCheckQuota != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "access_key is invalid",
		})
	}

	//check rate limit per minute with redis
	// Key to store the rate limit counter
	key := "rate_limit:" + userData.UserId

	// Increment the counter by 1
	count, err := rdb.Incr(context.Background(), key).Result()
	if err != nil {
		// Handle Redis error
		log.Println("err", err)
	}

	// Set the expiration time to 1 minute
	if count == 1 {
		_, err := rdb.Expire(context.Background(), key, time.Minute).Result()
		if err != nil {
			// Handle Redis error
			log.Println("err", err)
		}
	}

	// Check if the counter exceeds the rate limit
	if count > quotaData.SubscriptionPlans.RateLimitPerMinute {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Rate limit exceeded",
		})
	}

	if quotaData.ScreenshotUsage.ScreenshotTaken >= quotaData.SubscriptionPlans.IncludedScreenshots && quotaData.SubscriptionPlans.Name == "Free" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "You have reached your quota",
		})
	}

	//list of links api to take screenshot
	listApi := []string{
		"http://screenshot:1323/v1/screenshot",
	}

	indexOfApi, err := getNextApiIndex(rdb, context.Background(), listApi)
	if err != nil {
		return err
	}

	apiLink := listApi[indexOfApi]

	body := []byte(`{}`)
	if (customData != module.CustomSet{}) {
		body = []byte(`{
			"custom": {
				"id": "` + customData.Id + `",
				"name": "` + customData.Name + `",
				"user_id": "` + customData.UserId + `",
				"css": "` + customData.CSS + `",
				"javascript": "` + customData.JavaScript + `",
				"cookies": "` + customData.Cookies + `",
				"localStorage": "` + customData.LocalStorage + `",
				"user_agent": "` + customData.UserAgent + `",
				"headers": "` + customData.Headers + `",
				"bucket_endpoint": "` + customData.BucketEndpoint + `",
				"bucket_default": "` + customData.BucketDefault + `",
				"bucket_access_key": "` + customData.BucketAccessKey + `",
				"bucket_secret_key": "` + customData.BucketSecretKey + `"
			}
		}`)
	}

	//request to api
	fullURL := apiLink + "?" + url
	resp, errRes := http.Post(fullURL, "application/json", bytes.NewBuffer(body))
	if errRes != nil {
		// Handle the connection error gracefully
		return c.JSON(200, map[string]interface{}{
			"status":  "error",
			"message": errRes.Error(),
		})
	}
	defer resp.Body.Close()

	//check if resp is application/json
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		//convert response to blob
		body, errBody := ioutil.ReadAll(resp.Body)
		if errBody != nil {
			return c.JSON(200, map[string]interface{}{
				"status":  "error",
				"message": errBody.Error(),
			})
		}
		return c.Blob(200, contentType, body)
	} else {
		log.Println("json")
		//get body response to json
		body, errBody := ioutil.ReadAll(resp.Body)
		if errBody != nil {
			log.Println("errBody", errBody)
			return c.JSON(200, map[string]interface{}{
				"status":  "error",
				"message": errBody.Error(),
			})
		}
		log.Println("body", string(body))

		//convert to json
		return c.JSON(200, body)
	}
}

func getNextApiIndex(rdb *redis.Client, ctx context.Context, list []string) (int, error) {
	key := "apiIndex"
	length := len(list)
	index, err := rdb.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return int(index-1) % length, nil
}
