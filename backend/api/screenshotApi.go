package api

import (
	lib "backend/lib"
	module "backend/module"
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

// TakeScreenshotByAPI takes a screenshot using the provided API.
// It requires the following parameters:
// - c: the echo.Context object
// - db: the dbx.Builder object for database operations
// - mongo: the *mongo.Collection object for MongoDB operations
// - rdb: the *redis.Client object for Redis operations
// It returns an error if any operation fails.
func TakeScreenshotByAPI(c echo.Context, db dbx.Builder, mongo *mongo.Collection, rdb *redis.Client) error {
	//get all query params
	// url := c.QueryString()
	access_key := ""
	saveToS3 := false
	pathFileName := lib.GenerateRandomString(10)
	custom := ""

	//optional
	url := ""
	width := int64(1280)
	height := int64(1024)
	fullScreen := false
	scrollDelay := int64(1)
	noAds := false
	noCookie := false
	delay := int64(0)
	blockTracker := false
	timeout := int64(60)
	element := "body"
	imageQuality := 100
	imageFormat := "png"

	//check method
	if c.Request().Method != "GET" {
		//get it from body
		body, errBody := ioutil.ReadAll(c.Request().Body)
		if errBody != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status":  "error",
				"message": errBody.Error(),
			})
		}

		//convert body to json
		jsonBody := map[string]interface{}{}
		errJsonBody := json.Unmarshal(body, &jsonBody)
		if errJsonBody != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status":  "error",
				"message": errJsonBody.Error(),
			})
		}

		//access_key
		access_key_raw, ok := jsonBody["access_key"].(string)
		if !ok {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status":  "error",
				"message": "access_key is required",
			})
		}
		access_key = access_key_raw

		//save_to_s3
		saveToS3Str, ok := jsonBody["save_to_s3"].(string)
		if ok {
			saveToS3_raw, err := strconv.ParseBool(saveToS3Str)
			if err != nil {
				saveToS3_raw = false
			}
			saveToS3 = saveToS3_raw
		}

		//path_file_name
		pathFileName_raw, ok := jsonBody["path_file_name"].(string)
		if ok {
			pathFileName = pathFileName_raw
		} else {
			//randomString
			pathFileName = lib.GenerateRandomString(10)
		}

		//custom
		custom_raw, ok := jsonBody["custom"].(string)
		if ok {
			custom = custom_raw
		}

		//* ---------------------- OPTIONAL ---------------------- *//
		//url
		url_raw, ok := jsonBody["url"].(string)
		if ok {
			url = url_raw
		}

		//v_width
		width_raw, ok := jsonBody["v_width"].(float64)
		if ok {
			width = int64(width_raw)
		}

		//v_height
		height_raw, ok := jsonBody["v_height"].(float64)
		if ok {
			height = int64(height_raw)
		}

		//full_screen
		fullScreen_raw, ok := jsonBody["full_screen"].(bool)
		if ok {
			fullScreen = fullScreen_raw
		}

		//scroll_delay
		scrollDelay_raw, ok := jsonBody["scroll_delay"].(float64)
		if ok {
			scrollDelay = int64(scrollDelay_raw)
		}

		//no_ads
		noAds_raw, ok := jsonBody["no_ads"].(bool)
		if ok {
			noAds = noAds_raw
		}

		//no_cookie
		noCookie_raw, ok := jsonBody["no_cookie"].(bool)
		if ok {
			noCookie = noCookie_raw
		}

		//delay
		delay_raw, ok := jsonBody["delay"].(float64)
		if ok {
			delay = int64(delay_raw)
		}

		//block_tracker
		blockTracker_raw, ok := jsonBody["block_tracker"].(bool)
		if ok {
			blockTracker = blockTracker_raw
		}

		//timeout
		timeout_raw, ok := jsonBody["timeout"].(float64)
		if ok {
			timeout = int64(timeout_raw)
		}

		//element
		element_raw, ok := jsonBody["element"].(string)
		if ok {
			element = element_raw
		}

		//image_quality
		imageQuality_raw, ok := jsonBody["quality"].(float64)
		if ok {
			imageQuality = int(imageQuality_raw)
		}

		//image_format
		imageFormat_raw, ok := jsonBody["format"].(string)
		if ok {
			imageFormat = imageFormat_raw
		}
		//* ---------------------- OPTIONAL ---------------------- *//

	} else {
		access_key_raw := c.QueryParam("access_key")
		if access_key_raw == "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status":  "error",
				"message": "access_key is required",
			})
		}
		access_key = access_key_raw

		saveToS3Str := c.QueryParam("save_to_s3")
		saveToS3_raw, err := strconv.ParseBool(saveToS3Str)
		if err != nil {
			saveToS3_raw = false
		}
		saveToS3 = saveToS3_raw

		pathFileName = c.QueryParam("path_file_name")
		if pathFileName == "" {
			//randomString
			pathFileName = lib.GenerateRandomString(10)
		}

		custom = c.QueryParam("custom")

		//* ---------------------- OPTIONAL ---------------------- *//
		url_raw := c.QueryParam("url")
		if url_raw != "" {
			url = url_raw
		}

		width_raw := c.QueryParam("v_width")
		if width_raw != "" {
			width, err = strconv.ParseInt(width_raw, 10, 64)
			if err != nil {
				width = 1280
			}
		}

		height_raw := c.QueryParam("v_height")
		if height_raw != "" {
			height, err = strconv.ParseInt(height_raw, 10, 64)
			if err != nil {
				height = 1024
			}
		}

		fullScreen_raw := c.QueryParam("full_screen")
		if fullScreen_raw != "" {
			fullScreen, err = strconv.ParseBool(fullScreen_raw)
			if err != nil {
				fullScreen = false
			}
		}

		scrollDelay_raw := c.QueryParam("scroll_delay")
		if scrollDelay_raw != "" {
			scrollDelay, err = strconv.ParseInt(scrollDelay_raw, 10, 64)
			if err != nil {
				scrollDelay = 1
			}
		}

		noAds_raw := c.QueryParam("no_ads")
		if noAds_raw != "" {
			noAds, err = strconv.ParseBool(noAds_raw)
			if err != nil {
				noAds = false
			}
		}

		noCookie_raw := c.QueryParam("no_cookie")
		if noCookie_raw != "" {
			noCookie, err = strconv.ParseBool(noCookie_raw)
			if err != nil {
				noCookie = false
			}
		}

		delay_raw := c.QueryParam("delay")
		if delay_raw != "" {
			delay, err = strconv.ParseInt(delay_raw, 10, 64)
			if err != nil {
				delay = 0
			}
		}

		blockTracker_raw := c.QueryParam("block_tracker")
		if blockTracker_raw != "" {
			blockTracker, err = strconv.ParseBool(blockTracker_raw)
			if err != nil {
				blockTracker = false
			}
		}

		timeout_raw := c.QueryParam("timeout")
		if timeout_raw != "" {
			timeout, err = strconv.ParseInt(timeout_raw, 10, 64)
			if err != nil {
				timeout = 60
			}
		}

		element_raw := c.QueryParam("element")
		if element_raw != "" {
			element = element_raw
		}

		imageQuality_raw := c.QueryParam("quality")
		if imageQuality_raw != "" {
			imageQuality, err = strconv.Atoi(imageQuality_raw)
			if err != nil {
				imageQuality = 100
			}
		}

		imageFormat_raw := c.QueryParam("format")
		if imageFormat_raw != "" {
			imageFormat = imageFormat_raw
		}
		//* ---------------------- OPTIONAL ---------------------- *//
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

	customData := module.CustomSet{}
	if custom != "" {
		errCustom := db.Select("id", "name", "user_id", "css", "javascript", "cookies", "localStorage", "user_agent", "headers", "bucket_endpoint", "bucket_default", "bucket_access_key", "bucket_secret_key").
			From("custom_sets").
			Where(dbx.NewExp("name = {:name} and user_id = {:user_id}", dbx.Params{"name": custom, "user_id": userData.UserId})).
			One(&customData)
		if errCustom != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status":  "error",
				"message": "custom is invalid",
			})
		}
	}

	// * ---------------------- QUOTA ---------------------- * //
	quotaData := module.GetQuotaScreenshot{}
	errCheckQuota := db.
		Select("screenshot_usage.screenshots_taken", "screenshot_usage.disable_extra", "subscription_plans.name", "subscription_plans.included_screenshots", "subscription_plans.rate_limit_per_minute").
		From("screenshot_usage").
		InnerJoin("users", dbx.NewExp("users.id = screenshot_usage.user_id")).
		InnerJoin("subscription_plans", dbx.NewExp("subscription_plans.id = users.subscription_plan")).
		Where(dbx.NewExp("screenshot_usage.user_id = {:user_id}", dbx.Params{"user_id": userData.UserId})).
		One(&quotaData)
	if errCheckQuota != nil {
		log.Println("errCheckQuota", errCheckQuota)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "user_id is invalid quota",
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
	} else if quotaData.ScreenshotUsage.ScreenshotTaken >= quotaData.SubscriptionPlans.IncludedScreenshots && quotaData.SubscriptionPlans.Name != "Free" {
		if quotaData.ScreenshotUsage.DisableExtra {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status":  "error",
				"message": "You have reached your quota",
			})
		}
		// ? update screenshot_usage table increase extra_screenshots_taken
		_, errorExtraScreenshot := db.NewQuery(`
			UPDATE screenshot_usage
			SET extra_screenshots_taken = extra_screenshots_taken + 1
			WHERE user_id = {:user_id}
		`).Bind(dbx.Params{
			"user_id": userData.UserId,
		}).Execute()
		if errorExtraScreenshot != nil {
			log.Println("errorExtraScreenshot", errorExtraScreenshot)
		}
	}
	// * ---------------------- QUOTA ---------------------- * //

	//list of links api to take screenshot
	listApi := []string{
		"http://screenshot:1323/v1/screenshot",
	}

	indexOfApi, err := getNextApiIndex(rdb, context.Background(), listApi)
	if err != nil {
		return err
	}

	apiLink := listApi[indexOfApi]

	body, err := json.Marshal(map[string]interface{}{})
	jsonData, err := json.Marshal(map[string]interface{}{
		"url":           url,
		"v_width":       strconv.FormatInt(width, 10),
		"v_height":      strconv.FormatInt(height, 10),
		"full_screen":   strconv.FormatBool(fullScreen),
		"scroll_delay":  strconv.FormatInt(scrollDelay, 10),
		"no_ads":        strconv.FormatBool(noAds),
		"no_cookie":     strconv.FormatBool(noCookie),
		"delay":         strconv.FormatInt(delay, 10),
		"block_tracker": strconv.FormatBool(blockTracker),
		"timeout":       strconv.FormatInt(timeout, 10),
		"element":       element,
		"quality":       strconv.Itoa(imageQuality),
		"format":        imageFormat,
	})
	if err != nil {
		log.Println("err", err)
	}

	if (customData != module.CustomSet{}) {
		custom := map[string]string{
			"id":                customData.Id,
			"name":              customData.Name,
			"user_id":           customData.UserId,
			"css":               customData.CSS,
			"javascript":        customData.JavaScript,
			"cookies":           customData.Cookies,
			"localStorage":      customData.LocalStorage,
			"user_agent":        customData.UserAgent,
			"headers":           customData.Headers,
			"bucket_endpoint":   customData.BucketEndpoint,
			"bucket_default":    customData.BucketDefault,
			"bucket_access_key": customData.BucketAccessKey,
			"bucket_secret_key": customData.BucketSecretKey,
		}

		jsonDataCustom, err := json.Marshal(map[string]interface{}{
			"custom":        custom,
			"url":           url,
			"v_width":       strconv.FormatInt(width, 10),
			"v_height":      strconv.FormatInt(height, 10),
			"full_screen":   strconv.FormatBool(fullScreen),
			"scroll_delay":  strconv.FormatInt(scrollDelay, 10),
			"no_ads":        strconv.FormatBool(noAds),
			"no_cookie":     strconv.FormatBool(noCookie),
			"delay":         strconv.FormatInt(delay, 10),
			"block_tracker": strconv.FormatBool(blockTracker),
			"timeout":       strconv.FormatInt(timeout, 10),
			"element":       element,
			"quality":       strconv.Itoa(imageQuality),
			"format":        imageFormat,
		})
		if err != nil {
			log.Println("err", err)
		}
		jsonData = jsonDataCustom
	}
	body = jsonData
	//request to api
	fullURL := apiLink

	fullUrlApi := c.Request().URL.String()
	mongo.InsertOne(context.Background(), map[string]interface{}{
		"user_id":    userData.UserId,
		"access_key": access_key,
		"url":        url,
		"fullUrl":    fullUrlApi,
		"created":    time.Now(),
	})

	//async
	asyncChromeStr := c.QueryParam("async")
	asyncChrome, err := strconv.ParseBool(asyncChromeStr)
	if err != nil {
		asyncChrome = false
	}
	if asyncChrome {
		go func() {
			// Make the HTTP request asynchronously
			resp, errRes := http.Post(fullURL, "application/json", bytes.NewBuffer(body))
			if errRes != nil {
				// Handle the connection error gracefully
				log.Println("errRes", errRes)
				return
			}
			//update screenshot_usage
			_, errUpdateScreenshotUsage := db.NewQuery(`
	UPDATE screenshot_usage
	SET screenshots_taken = screenshots_taken + 1
	WHERE user_id = {:user_id}
`).Bind(dbx.Params{
				"user_id": userData.UserId,
			}).Execute()
			if errUpdateScreenshotUsage != nil {
				log.Println("errUpdateScreenshotUsage", errUpdateScreenshotUsage)
			}
			defer resp.Body.Close()
			bodyAsync, errBody := ioutil.ReadAll(resp.Body)
			if errBody != nil {
				log.Println("errBody", errBody)
				return
			}

			// * SAVE TO S3 * //
			imageType := http.DetectContentType(bodyAsync)
			//imageType to dot
			dotTypeImage := "." + strings.Split(imageType, "/")[1]
			log.Println("dotTypeImage async", dotTypeImage)
			if saveToS3 && asyncChrome && customData.BucketDefault != "" && customData.BucketAccessKey != "" && customData.BucketSecretKey != "" && customData.BucketEndpoint != "" {
				log.Println("save to s3 async")
				err := lib.UploadToS3(bodyAsync, pathFileName+dotTypeImage, customData.BucketDefault, customData.BucketAccessKey, customData.BucketSecretKey, customData.BucketEndpoint)
				if err != nil {
					log.Println("err", err)
				}
			}
		}()

		return c.JSON(200, map[string]interface{}{
			"status":  "success",
			"message": "Screenshot is being processed",
		})
	} else {
		resp, errRes := http.Post(fullURL, "application/json", bytes.NewBuffer(body))
		if errRes != nil {
			// Handle the connection error gracefully
			return c.JSON(200, map[string]interface{}{
				"status":  "error",
				"message": errRes.Error(),
			})
		}
		defer resp.Body.Close()

		//update screenshot_usage
		_, errUpdateScreenshotUsage := db.NewQuery(`
	UPDATE screenshot_usage
	SET screenshots_taken = screenshots_taken + 1
	WHERE user_id = {:user_id}
`).Bind(dbx.Params{
			"user_id": userData.UserId,
		}).Execute()
		if errUpdateScreenshotUsage != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status":  "error",
				"message": errUpdateScreenshotUsage.Error(),
			})
		}

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
			// * SAVE TO S3 * //
			imageType := http.DetectContentType(body)
			//imageType to dot
			dotTypeImage := "." + strings.Split(imageType, "/")[1]
			// log.Println("dotTypeImage", dotTypeImage)
			if saveToS3 && !asyncChrome && customData.BucketDefault != "" && customData.BucketAccessKey != "" && customData.BucketSecretKey != "" && customData.BucketEndpoint != "" {
				log.Println("save to s3")
				err := lib.UploadToS3(body, pathFileName+dotTypeImage, customData.BucketDefault, customData.BucketAccessKey, customData.BucketSecretKey, customData.BucketEndpoint)
				if err != nil {
					log.Println("err", err)
				}
				return c.JSON(200, map[string]interface{}{
					"status":          "success",
					"file":            pathFileName + dotTypeImage,
					"bucket":          customData.BucketDefault,
					"bucket_endpoint": customData.BucketEndpoint,
					"message":         "Screenshot is saved to S3",
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

			//convert to json
			return c.JSONBlob(200, body)
		}
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
