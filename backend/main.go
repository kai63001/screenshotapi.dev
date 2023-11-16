package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	api "backend/api"
	lib "backend/lib"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := pocketbase.New()

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_DB_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("capture").Collection("logs")

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/api/screenshot", func(c echo.Context) error {
			database := app.Dao().DB()
			return api.TakeScreenshot(c, database, collection)
		})
		e.Router.GET("/api/history", func(c echo.Context) error {
			database := app.Dao().DB()
			return api.GetHistoryScreenshotAPI(c, database, collection)
		})
		e.Router.POST("/api/subscription", func(c echo.Context) error {
			database := app.Dao().DB()
			return api.Subscription(c, database)
		})
		e.Router.POST("/api/hook", func(c echo.Context) error {
			database := app.Dao().DB()
			return api.Hook(c, database)
		})
		e.Router.POST("/api/portal", func(c echo.Context) error {
			database := app.Dao().DB()
			return api.StripePortal(c, database)
		})
		e.Router.PATCH("/api/access_key", func(c echo.Context) error {
			database := app.Dao().DB()
			return api.ResetAccessKey(c, database)
		})
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})

	app.OnModelAfterCreate("users").Add(func(e *core.ModelEvent) error {
		_, errCreateScreenshotUsage := app.Dao().DB().NewQuery(`
			INSERT INTO screenshot_usage (user_id)
			VALUES ({:user_id})
		`).Bind(dbx.Params{
			"user_id": e.Model.GetId(),
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

		// update subscription_plan to free user table
		_, errUpdateSubscriptionPlan := app.Dao().DB().NewQuery(`
			UPDATE users
			SET subscription_plan = {:subscription_plan}
			WHERE id = {:id}
		`).Bind(dbx.Params{
			"subscription_plan": os.Getenv("FREE_PLAN_ID"),
			"id":                e.Model.GetId(),
		}).Execute()
		if errUpdateSubscriptionPlan != nil {
			return errUpdateSubscriptionPlan
		}

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
