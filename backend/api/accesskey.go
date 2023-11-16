package api

import (
	lib "backend/lib"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/apis"
)

func ResetAccessKey(c echo.Context, db dbx.Builder) error {

	//check auth
	info := apis.RequestInfo(c)

	record := info.AuthRecord

	if record == nil {
		return c.JSON(200, map[string]interface{}{
			"status":  "error",
			"message": "not authenticated",
		})
	}

	type AccessKey struct {
	}

	listAccessKey := []AccessKey{}

	accessKeyExists := true
	accessKey := ""
	//create pointer for store result

	for accessKeyExists {
		accessKey = lib.GenerateRandomString(15)
		errAccessKeyExists := db.NewQuery(`
            SELECT access_key
            FROM access_keys
            WHERE access_key = {:access_key}
    `).Bind(dbx.Params{
			"access_key": accessKey,
		}).All(&listAccessKey)

		if errAccessKeyExists != nil {
			return errAccessKeyExists
		}

		if (len(listAccessKey)) == 0 {
			accessKeyExists = false
		}
		listAccessKey = []AccessKey{}
	}

	//update access_key
	_, err := db.NewQuery(`
		UPDATE access_keys
		SET access_key = {:access_key}
		WHERE user_id = {:user_id}
	`).Bind(dbx.Params{
		"access_key": accessKey,
		"user_id":    record.GetId(),
	}).Execute()
	if err != nil {
		return c.JSON(200, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status":     "success",
		"message":    "access key reset",
		"access_key": accessKey,
	})
}
