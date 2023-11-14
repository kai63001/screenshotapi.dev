package api

import (
	"io/ioutil"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/stripe/stripe-go/webhook"
)

func Hook(c echo.Context, db dbx.Builder) error {

	//get body
	payload, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(200, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	endpointSecret := os.Getenv("STRIPE_HOOK_SECRET")
	// Pass the request body and Stripe-Signature header to ConstructEvent, along
	// with the webhook signing key.
	event, err := webhook.ConstructEvent(payload, c.Request().Header.Get("Stripe-Signature"),
		endpointSecret)
	if err != nil {
		return c.JSON(200, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	switch event.Type {
	case "customer.subscription.created":
		customerId := event.Data.Object["customer"].(string)
		subscriptionId := event.Data.Object["id"].(string)
		product := event.Data.Object["plan"].(map[string]interface{})["product"].(string)

		type SubscriptionPlan struct {
			Id string `db:"id"`
		}

		//get id subscription plans from product id = striep_product_id
		subscriptionPlanId := SubscriptionPlan{}
		err := db.NewQuery(`
			SELECT id
			FROM subscription_plans
			WHERE stripe_product_id = {:stripe_product_id}
		`).Bind(dbx.Params{
			"stripe_product_id": product,
		}).One(&subscriptionPlanId)
		if err != nil {
			return c.JSON(200, map[string]interface{}{
				"status":  "error",
				"message": err.Error(),
			})
		}

		//update user subscription
		_, errUpdate := db.NewQuery(`
			UPDATE users
			SET subscription_plan = {:subscription_plan_id}, stripe_subscription_id = {:stripe_subscription_id}, subscription_status = 'active'
			WHERE stripe_customer_id = {:stripe_customer_id}
		`).Bind(dbx.Params{
			"subscription_plan_id":   subscriptionPlanId.Id,
			"stripe_subscription_id": subscriptionId,
			"stripe_customer_id":     customerId,
		}).Execute()
		if errUpdate != nil {
			return c.JSON(200, map[string]interface{}{
				"status":  "error",
				"message": errUpdate.Error(),
			})
		}
	case "customer.subscription.deleted":
		// Then define and call a function to handle the event customer.subscription.deleted
	case "invoice.payment_failed":
		// Then define and call a function to handle the event invoice.payment_failed
	// ... handle other event types
	default:
		return c.JSON(200, map[string]interface{}{
			"status":  "error",
			"message": "Unhandled event type " + event.Type,
		})
	}

	return nil
}
