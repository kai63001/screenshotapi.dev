package api

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func TakeScreenshotByAPI(c echo.Context, db dbx.Builder, mongo *mongo.Collection, rdb *redis.Client) error {
	//get all query params
	url := c.QueryString()

	//list of links api to take screenshot
	listApi := []string{
		"http://screenshot:1323/v1/screenshot",
	}

	indexOfApi, err := getNextApiIndex(rdb, context.Background(), listApi)
	if err != nil {
		return err
	}

	apiLink := listApi[indexOfApi]

	//request to api
	fullURL := apiLink + "?" + url
	resp, errRes := http.Get(fullURL)
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
	if contentType != "application/json" {

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
		//get body response to json
		body, errBody := ioutil.ReadAll(resp.Body)
		if errBody != nil {
			return c.JSON(200, map[string]interface{}{
				"status":  "error",
				"message": errBody.Error(),
			})
		}
		jsonBody := string(body)
		//convert to json
		return c.JSONBlob(200, []byte(jsonBody))
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
