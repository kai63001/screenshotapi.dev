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

	plan_id := json_map["plan_id"].(string)
	is_yearly := json_map["is_yearly"].(bool)
	//get subscription plan with id
	type SubscriptionPlan struct {
		Id string `db:"id"`
		// StripePricingPlan struct {
		// 	Monthly struct {
		// 		Flat string `db:"flat"`
		// 		Unit string `db:"unit"`
		// 	} `db:"monthly"`
		// 	Yearly struct {
		// 		Flat string `db:"flat"`
		// 		Unit string `db:"unit"`
		// 	} `db:"yearly"`
		// } `db:"stripe_pricing_id"`
		StripePricingPlan string `db:"stripe_pricing_id"`
	}

	subscriptionPlan := SubscriptionPlan{}

	err = db.NewQuery(`
		SELECT id, stripe_pricing_id
		FROM subscription_plans
		WHERE id = {:id}
	`).Bind(dbx.Params{
		"id": plan_id,
	}).One(&subscriptionPlan)
	if err != nil {
		return c.JSON(200, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	type StripePricingPlanType struct {
		Monthly struct {
			Flat string `json:"flat"`
			Unit string `json:"unit"`
		} `json:"monthly"`
		Yearly struct {
			Flat string `json:"flat"`
			Unit string `json:"unit"`
		} `json:"yearly"`
	}

	var StripePricingPlan StripePricingPlanType
	//json
	err = json.Unmarshal([]byte(subscriptionPlan.StripePricingPlan), &StripePricingPlan)

	// log.Println("subscriptionPlan", subscriptionPlan)

	var flatId string
	var unitId string
	if is_yearly {
		flatId = StripePricingPlan.Yearly.Flat
		unitId = StripePricingPlan.Yearly.Unit
	} else {
		flatId = StripePricingPlan.Monthly.Flat
		unitId = StripePricingPlan.Monthly.Unit
	}

	//create subscription
	params := &stripe.CheckoutSessionParams{
		Customer: stripe.String(stripe_customer_id),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(flatId),
				Quantity: stripe.Int64(1),
			},
			{
				Price: stripe.String(unitId),
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
