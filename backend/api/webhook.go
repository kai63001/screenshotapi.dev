package api

import (
	"io/ioutil"
	"log"
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
	case "invoice.paid":
		customerId := event.Data.Object["customer"].(string)
		invoice_pdf := event.Data.Object["invoice_pdf"].(string)
		paymentId := event.Data.Object["payment_intent"].(string)
		status := event.Data.Object["status"].(string)
		totalAmount := event.Data.Object["total"].(float64)
		// productId := event.Data.Object["lines"]

		productId := event.Data.Object["lines"].(map[string]interface{})["data"].([]interface{})[0].(map[string]interface{})["plan"].(map[string]interface{})["product"].(string)
		productDescription := event.Data.Object["lines"].(map[string]interface{})["data"].([]interface{})[0].(map[string]interface{})["description"].(string)

		// log.Println("invoice.paid", customerId, invoice_pdf, paymentId, status, totalAmount)

		// //insert payments
		_, errInsert := db.NewQuery(`
			INSERT INTO payments (customer_id, stripe_payment_id, invoice_pdf, status, total_amount, stripe_product_id, product_description)
			VALUES ({:customer_id}, {:stripe_payment_id}, {:invoice_pdf}, {:status}, {:total_amount}, {:product_id}, {:product_description})
		`).Bind(dbx.Params{
			"customer_id":         customerId,
			"stripe_payment_id":   paymentId,
			"invoice_pdf":         invoice_pdf,
			"status":              status,
			"total_amount":        totalAmount,
			"product_id":          productId,
			"product_description": productDescription,
		}).Execute()
		if errInsert != nil {
			return c.JSON(200, map[string]interface{}{
				"status":  "error",
				"message": errInsert.Error(),
			})
		}

	case "customer.subscription.created":
		customerId := event.Data.Object["customer"].(string)
		subscriptionId := event.Data.Object["id"].(string)
		items := event.Data.Object["items"].(map[string]interface{})
		data := items["data"].([]interface{})

		product := ""
		if len(data) > 0 {
			firstItem := data[0].(map[string]interface{})
			plan := firstItem["plan"].(map[string]interface{})
			product = plan["product"].(string)
		}
		log.Println("product:", product)

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
			log.Println("error get subscription plan id", err)
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

	case "checkout.session.expired":
		customerId := event.Data.Object["customer"].(string)

		//update user subscription
		_, errUpdate := db.NewQuery(`
			UPDATE users
			SET subscription_plan = '', stripe_subscription_id = '', subscription_status = 'cancelled'
			WHERE stripe_customer_id = {:stripe_customer_id}
		`).Bind(dbx.Params{
			"stripe_customer_id": customerId,
		}).Execute()
		if errUpdate != nil {
			return c.JSON(200, map[string]interface{}{
				"status":  "error",
				"message": errUpdate.Error(),
			})
		}
	case "customer.subscription.updated":
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
		customerId := event.Data.Object["customer"].(string)

		//update user subscription
		_, errUpdate := db.NewQuery(`
			UPDATE users
			SET subscription_plan = {:subscription_plan}, stripe_subscription_id = '', subscription_status = 'cancelled'
			WHERE stripe_customer_id = {:stripe_customer_id}
		`).Bind(dbx.Params{
			"stripe_customer_id": customerId,
			"subscription_plan":  os.Getenv("FREE_PLAN_ID"),
		}).Execute()
		if errUpdate != nil {
			return c.JSON(200, map[string]interface{}{
				"status":  "error",
				"message": errUpdate.Error(),
			})
		}

		status := event.Data.Object["status"].(string)

		// //insert payments
		_, errInsert := db.NewQuery(`
			INSERT INTO payments (customer_id, stripe_payment_id, invoice_pdf, status, total_amount, stripe_product_id, product_description)
			VALUES ({:customer_id}, {:stripe_payment_id}, {:invoice_pdf}, {:status}, {:total_amount}, {:product_id}, {:product_description})
		`).Bind(dbx.Params{
			"customer_id":         customerId,
			"stripe_payment_id":   "",
			"invoice_pdf":         "",
			"status":              status,
			"total_amount":        0,
			"product_id":          "",
			"product_description": "",
		}).Execute()
		if errInsert != nil {
			return c.JSON(200, map[string]interface{}{
				"status":  "error",
				"message": errInsert.Error(),
			})
		}
	case "invoice.payment_failed":
		customerId := event.Data.Object["customer"].(string)

		//update user subscription
		_, errUpdate := db.NewQuery(`
			UPDATE users
			SET subscription_plan = {:subscription_plan}, stripe_subscription_id = ', subscription_status = 'cancelled'
			WHERE stripe_customer_id = {:stripe_customer_id}
		`).Bind(dbx.Params{
			"stripe_customer_id": customerId,
			"subscription_plan":  os.Getenv("FREE_PLAN_ID"),
		}).Execute()
		if errUpdate != nil {
			return c.JSON(200, map[string]interface{}{
				"status":  "error",
				"message": errUpdate.Error(),
			})
		}

		//insert payments
		_, errInsert := db.NewQuery(`
			INSERT INTO payments (customer_id, stripe_payment_id, invoice_pdf, status, total_amount)
			VALUES ({:customer_id}, {:stripe_payment_id}, {:invoice_pdf}, {:status}, {:total_amount})
		`).Bind(dbx.Params{
			"customer_id":       customerId,
			"stripe_payment_id": "",
			"invoice_pdf":       "",
			"status":            "Failed",
			"total_amount":      "",
		}).Execute()
		if errInsert != nil {
			return c.JSON(200, map[string]interface{}{
				"status":  "error",
				"message": errInsert.Error(),
			})
		}
	default:
		return c.JSON(200, map[string]interface{}{
			"status":  "error",
			"message": "Unhandled event type " + event.Type,
		})
	}

	return nil
}
