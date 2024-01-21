package api

import (
	"backend/lib"
	"backend/module"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/apis"
)

func DeleteAccount(c echo.Context, db dbx.Builder) error {
	info := apis.RequestInfo(c)

	record := info.AuthRecord

	if record == nil {
		return c.JSON(200, map[string]interface{}{
			"status":  "error",
			"message": "not authenticated",
		})
	}

	userId := record.Get("id").(string)

	//get customer id with user id
	user := module.User{}
	errorGetUser := db.Select("stripe_customer_id").From("users").Where(
		dbx.NewExp("id = {:user_id}", dbx.Params{
			"user_id": userId,
		}),
	).One(&user)
	if errorGetUser != nil {
		return c.JSON(400, errorGetUser.Error())
	}

	//* ----------------- delete stripe customer ----------------- *//
	// check if not found stripe customer id
	if user.StripeCustomerId != "" {
		resultDelte := lib.DeleteStripeUserWithCustomerId(db, user.StripeCustomerId)
		if !resultDelte {
			return c.JSON(400, "delete stripe customer failed")
		}
	}

	//* ----------------- delete user ----------------- *//
	//Delete user by id
	_, errorDeleteUser := db.NewQuery(`
		DELETE FROM users
		WHERE id = {:user_id}
	`).Bind(dbx.Params{
		"user_id": userId,
	}).Execute()
	if errorDeleteUser != nil {
		return c.JSON(400, errorDeleteUser.Error())
	}

	return c.JSON(200, map[string]interface{}{
		"status":  "success",
		"message": "user deleted",
	})

}
