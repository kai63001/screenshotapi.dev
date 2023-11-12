package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"

	api "backend/api"
	lib "backend/lib"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/api/screenshot", func(c echo.Context) error {
			database := app.Dao().DB()
			return api.TakeScreenshot(c, database)
		})
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})

	app.OnModelAfterCreate("users").Add(func(e *core.ModelEvent) error {
		_, errCreateScreenshotUsage := app.Dao().DB().NewQuery(`
			INSERT INTO screenshot_usage (user_id,subscription_plan)
			VALUES ({:user_id}, {:plan})
		`).Bind(dbx.Params{
			"user_id": e.Model.GetId(),
			"plan":    os.Getenv("FREE_PLAN_ID"),
		}).Execute()
		if errCreateScreenshotUsage != nil {
			return errCreateScreenshotUsage
		}

		type AccessKey struct {
		}

		listAccessKey := []AccessKey{}

		accessKeyExists := true
		accessKey := ""
		//create pointer for store result

		for accessKeyExists {
			accessKey = lib.GenerateRandomString(15)
			errAccessKeyExists := app.Dao().DB().NewQuery(`
            SELECT access_key
            FROM access_keys
            WHERE access_key = {:access_key}
    `).Bind(dbx.Params{
				"access_key": accessKey,
			}).All(&listAccessKey)

			log.Println("accessKeyExists", accessKeyExists)

			if errAccessKeyExists != nil {
				return errAccessKeyExists
			}

			if (len(listAccessKey)) == 0 {
				accessKeyExists = false
			}
			listAccessKey = []AccessKey{}
		}

		// create access key
		_, errCreateAccessKey := app.Dao().DB().NewQuery(`
			INSERT INTO access_keys (user_id,access_key)
			VALUES ({:user_id}, {:access_key})
		`).Bind(dbx.Params{
			"user_id":    e.Model.GetId(),
			"access_key": accessKey,
		}).Execute()
		if errCreateAccessKey != nil {
			return errCreateAccessKey
		}
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
