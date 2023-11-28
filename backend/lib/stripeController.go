package lib

import (
	"log"
	"os"
	"time"

	"github.com/pocketbase/dbx"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/usagerecord"
)

func ReportUsage(db dbx.Builder) {

	stripe.Key = os.Getenv("STRIPE_KEY")

	usageQuantity := int64(100)

	log.Println("usageQuantity", usageQuantity)

	params := &stripe.UsageRecordParams{
		SubscriptionItem: stripe.String("si_P5JaH9qiRpP9Nv"),
		Quantity:         stripe.Int64(usageQuantity),
		Timestamp:        stripe.Int64(time.Now().Unix()),
		Action:           stripe.String(string(stripe.UsageRecordActionSet)),
	}

	usagerecord.New(params)
}
