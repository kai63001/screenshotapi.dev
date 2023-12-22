package api

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/apis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetHistoryScreenshotAPI(c echo.Context, db dbx.Builder, mongo *mongo.Collection) error {
	//this is list of history screenshot api support pagination

	//check auth
	info := apis.RequestInfo(c)

	record := info.AuthRecord
	if record == nil {
		return c.JSON(200, map[string]interface{}{
			"status":  "error",
			"message": "not authenticated",
		})
	}

	type HistoryScreenshot struct {
		Id        string    `json:"_id" bson:"_id"`
		UserId    string    `json:"user_id" bson:"user_id"`
		AccessKey string    `json:"access_key" bson:"access_key"`
		Url       string    `json:"url" bson:"url"`
		FullUrl   string    `json:"fullUrl" bson:"fullUrl"`
		Created   time.Time `json:"created" bson:"created"`
	}

	listHistoryScreenshot := []HistoryScreenshot{}

	// get list history screenshot from mongo with pagination and limit
	pageStr := c.QueryParam("page")
	limitStr := c.QueryParam("limit")

	// set default values if not found
	page := 1
	limit := 10

	// convert page and limit to integers
	if pageStr != "" {
		pageInt, err := strconv.Atoi(pageStr)
		if err == nil {
			page = pageInt
		}
	}

	if limitStr != "" {
		limitInt, err := strconv.Atoi(limitStr)
		if err == nil {
			limit = limitInt
		}
	}
	skip := (page - 1) * limit

	result, err := mongo.Find(c.Request().Context(), map[string]interface{}{
		"user_id": record.GetId(),
	}, options.Find().SetSort(bson.D{{Key: "created", Value: -1}}).SetSkip(int64(skip)).SetLimit(int64(limit)))
	if err != nil {
		return err
	}

	for result.Next(c.Request().Context()) {
		var historyScreenshot HistoryScreenshot
		err := result.Decode(&historyScreenshot)
		if err != nil {
			return err
		}

		listHistoryScreenshot = append(listHistoryScreenshot, historyScreenshot)
	}

	// get total count
	totalCount, err := mongo.CountDocuments(c.Request().Context(), map[string]interface{}{
		"user_id": record.GetId(),
	})
	if err != nil {
		return err
	}

	hasNextPage := false
	hasPrevPage := false

	if len(listHistoryScreenshot) == limit {
		hasNextPage = true
	}

	if page > 1 {
		hasPrevPage = true
	}

	return c.JSON(200, map[string]interface{}{
		"status":      "success",
		"message":     "list history screenshot",
		"data":        listHistoryScreenshot,
		"totalCount":  totalCount,
		"hasNextPage": hasNextPage,
		"hasPrevPage": hasPrevPage,
	})

}
