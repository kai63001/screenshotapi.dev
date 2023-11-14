package api

import (
	"io/ioutil"

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

	endpointSecret := "whsec_pOncUtqQuU2ZGCuhNUzQxjdaeYGBSjHM"
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
		// Then define and call a function to handle the event customer.subscription.created
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
