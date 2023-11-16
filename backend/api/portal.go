package api

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/stripe/stripe-go/v72"
	session "github.com/stripe/stripe-go/v72/billingportal/session"
)

func StripePortal(c echo.Context, db dbx.Builder) error {
	//check auth
	info := apis.RequestInfo(c)

	record := info.AuthRecord

	if record == nil {
		return c.JSON(200, map[string]interface{}{
			"status":  "error",
			"message": "not authenticated",
		})
	}

	//get stripe_customer_id
	stripe_customer_id := record.Get("stripe_customer_id").(string)
	if stripe_customer_id == "" {
		return c.JSON(200, map[string]interface{}{
			"status":  "error",
			"message": "no stripe_customer_id",
		})
	}

	//create portal stripe
	stripe.Key = "sk_test_51OC3rdH2Tv3zxv6JmOoTODjamxamEZxuCZWVKc0tftmPC08vzGsjn2lUBI0u8vDNZN3ffwIXybl5ZcLO48OB8GQc00JsZ15Wkf"

	result := &stripe.BillingPortalSessionParams{
		Customer:  stripe.String(stripe_customer_id),
		ReturnURL: stripe.String("https://capture.pocketbase.app/portal"),
	}

	portal, err := session.New(result)
	if err != nil {
		return c.JSON(200, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status":  "success",
		"message": "portal created",
		"portal":  portal,
	})

}
