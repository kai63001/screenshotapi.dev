package api

import (
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

	//print record
	log.Println(record.Get("stripe_customer_id"))
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

	//create subscription
	params := &stripe.CheckoutSessionParams{
		Customer: stripe.String(stripe_customer_id),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String("price_1OC40uH2Tv3zxv6J2x9TstWP"),
				Quantity: stripe.Int64(1),
			},
		},
		Mode: stripe.String("subscription"),
		PaymentMethodTypes: []*string{
			stripe.String("card"),
		},
		AllowPromotionCodes: stripe.Bool(true),
		SuccessURL:          stripe.String("http://localhost:5173/subscription"),
	}

	subscription, err := session.New(params)
	if err != nil {
		return c.JSON(200, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	log.Println("subscription", subscription.URL)

	return c.JSON(200, map[string]interface{}{
		"status": "success",
		"url":    subscription.URL,
	})
}
