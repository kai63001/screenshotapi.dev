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
	"github.com/pocketbase/pocketbase/tools/cron"
	"github.com/redis/go-redis/v9"
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
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,                           // use default DB
	})

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_DB_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("capture").Collection("logs")

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		//scheduler
		database := app.Dao().DB()

		scheduler := cron.New()

		// prints "Hello!" every 2 minutes 0 */12 * * *
		scheduler.MustAdd("ResetQuotaPerMonth", "0 */12 * * *", func() {
			lib.ResetQuotaPerMonth(database)
		})

		scheduler.MustAdd("ResetQuotaPerMonth", "0 */2 * * *", func() {
			lib.ReportUsage(database)
		})

		scheduler.Start()

		//* ------------------ API SCREENSHOT ------------------ *//
		e.Router.GET("/v1/screenshot", func(c echo.Context) error {
			return api.TakeScreenshotByAPI(c, database, collection, rdb)
		})
		e.Router.POST("/v1/screenshot", func(c echo.Context) error {
			return api.TakeScreenshotByAPI(c, database, collection, rdb)
		})
		//* ------------------ API SCREENSHOT ------------------ *//

		e.Router.GET("/v1/history", func(c echo.Context) error {
			return api.GetHistoryScreenshotAPI(c, database, collection)
		})
		e.Router.POST("/v1/subscription", func(c echo.Context) error {
			return api.Subscription(c, database)
		})
		e.Router.POST("/v1/hook", func(c echo.Context) error {
			return api.Hook(c, database)
		})
		e.Router.POST("/v1/portal", func(c echo.Context) error {
			return api.StripePortal(c, database)
		})
		e.Router.PATCH("/v1/access_key", func(c echo.Context) error {
			return api.ResetAccessKey(c, database)
		})
		e.Router.POST("/v1/update_disable_extra", func(c echo.Context) error {
			return api.UpdateDisableExtra(c, database)
		})
		e.Router.DELETE("/v1/delete_account", func(c echo.Context) error {
			return api.DeleteAccount(c, database)
		})
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))

		return nil
	})

	app.OnModelAfterCreate("users").Add(func(e *core.ModelEvent) error {
		_, errCreateScreenshotUsage := app.Dao().DB().NewQuery(`
			INSERT INTO screenshot_usage (user_id,next_reset_quota)
			VALUES ({:user_id}, {:next_reset_quota})
		`).Bind(dbx.Params{
			"user_id": e.Model.GetId(),
			// set next reset quota to 30 days
			"next_reset_quota": lib.GetNextResetQuota(),
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
