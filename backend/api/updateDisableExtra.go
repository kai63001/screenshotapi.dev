package api

import (
	"encoding/json"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/apis"
)

func UpdateDisableExtra(c echo.Context, db dbx.Builder) error {
	info := apis.RequestInfo(c)

	record := info.AuthRecord

	if record == nil {
		return c.JSON(200, map[string]interface{}{
			"status":  "error",
			"message": "not authenticated",
		})
	}

	userId := record.Get("id").(string)

	body := c.Request().Body
	if body != nil {
		var jsonMap map[string]interface{}
		err := json.NewDecoder(body).Decode(&jsonMap)
		if err != nil {
			return c.JSON(400, "invalid json")
		}

		status, ok := jsonMap["status"].(bool)
		if !ok {
			return c.JSON(400, "status not provided or not a boolean")
		}

		// update disable_extra on screenshot_usage by user_id
		_, errorUpdateScreenshotUsage := db.NewQuery(`
			UPDATE screenshot_usage
			SET disable_extra = {:disable_extra}
			WHERE user_id = {:user_id}
		`).Bind(dbx.Params{
			"user_id":       userId,
			"disable_extra": status,
		}).Execute()

		if errorUpdateScreenshotUsage != nil {
			return c.JSON(400, errorUpdateScreenshotUsage.Error())
		}

		return c.JSON(200, map[string]interface{}{
			"status":  "success",
			"message": "disable extra updated",
		})
	}

	return c.JSON(400, map[string]interface{}{
		"status":  "error",
		"message": "update disable extra failed",
	})

}
