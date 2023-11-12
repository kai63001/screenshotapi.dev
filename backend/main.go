package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"

	api "backend/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/api/screenshot", api.TakeScreenshot)
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})

	app.OnModelAfterCreate("users").Add(func(e *core.ModelEvent) error {
		_, err := app.Dao().DB().NewQuery(`
			INSERT INTO screenshot_usage (user_id,subscription_plan)
			VALUES ({:user_id}, {:plan})
		`).Bind(dbx.Params{
			"user_id": e.Model.GetId(),
			"plan":    os.Getenv("FREE_PLAN_ID"),
		}).Execute()
		if err != nil {
			return err
		}
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
