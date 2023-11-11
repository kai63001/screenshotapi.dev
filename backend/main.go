package main

import (
	"log"
	"os"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"

	api "backend/api"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/api/screenshot", api.TakeScreenshot)
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})

	app.OnModelAfterCreate("users").Add(func(e *core.ModelEvent) error {
		log.Println("user created", e.Model.GetId())
		// collection := &models.Collection{
		// 	Name: "screenshot_usage",

		_, err := app.Dao().DB().NewQuery(`
			INSERT INTO screenshot_usage (user_id)
			VALUES ({:user_id})
		`).Bind(dbx.Params{
			"user_id": e.Model.GetId(),
		}).Execute()
		if err != nil {
			return err
		}
		//after create user insert into screenshot_usage table
		// e.Model.GetId()

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
