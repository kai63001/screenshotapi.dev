package api

import (
	"encoding/json"
	"log"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/apis"
	stripe "github.com/stripe/stripe-go/v76"
	session "github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/customer"
)

func Subscription(c echo.Context, db dbx.Builder) error {
	//check auth
	info := apis.RequestInfo(c)
	record := info.AuthRecord

	if record == nil {
		return c.JSON(200, map[string]interface{}{
			"status":  "error",
			"message": "not authenticated",
		})
	}

	stripe.Key = os.Getenv("STRIPE_KEY")

	stripe_customer_id := record.Get("stripe_customer_id").(string)
	if stripe_customer_id == "" {
		//update stripe_customer_id
		params := &stripe.CustomerParams{
			Email: stripe.String(record.Get("email").(string)),
		}
		customerData, err := customer.New(params)
		if err != nil {
			return c.JSON(200, map[string]interface{}{
				"status":  "error",
				"message": err.Error(),
			})
		}

		// seti_id = si.PaymentMethod.ID

		stripe_customer_id = customerData.ID
		_, err = db.NewQuery(`
			UPDATE users
			SET stripe_customer_id = {:stripe_customer_id}
			WHERE id = {:id}
		`).Bind(dbx.Params{
			"stripe_customer_id": stripe_customer_id,
			"id":                 record.GetId(),
		}).Execute()
		if err != nil {
			return c.JSON(200, map[string]interface{}{
				"status":  "error",
				"message": err.Error(),
			})
		}
	}

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		log.Println("err", err)
		return err
	}

	pricing_id := json_map["pricing_id"].(string)

	//create subscription
	params := &stripe.CheckoutSessionParams{
		Customer: stripe.String(stripe_customer_id),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(pricing_id),
				Quantity: stripe.Int64(1),
			},
		},
		Mode: stripe.String("subscription"),
		PaymentMethodTypes: []*string{
			stripe.String("card"),
		},
		AllowPromotionCodes: stripe.Bool(true),
		SuccessURL:          stripe.String(os.Getenv("FRONTEND_URL") + "/subscription/success"),
		CancelURL:           stripe.String(os.Getenv("FRONTEND_URL") + "/subscription/cancel"),
	}

	subscription, err := session.New(params)
	if err != nil {
		return c.JSON(200, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status": "success",
		"data":   subscription,
		"url":    subscription.URL,
	})
}
