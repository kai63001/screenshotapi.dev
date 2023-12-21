package lib

import (
	"backend/module"
	"log"
	"os"
	"time"

	"github.com/pocketbase/dbx"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/sub"
	"github.com/stripe/stripe-go/usagerecord"
)

func ReportUsage(db dbx.Builder) {

	stripe.Key = os.Getenv("STRIPE_KEY")

	//loop screenshotUsage
	listScreenshotUsage := []module.ScreenshotUsage{}
	errorListScreenshotUsage := db.Select("user_id", "screenshots_taken", "extra_screenshots_taken", "next_reset_quota", "disable_extra").From("screenshot_usage").Where(
		dbx.NewExp("extra_screenshots_taken > 0"),
	).All(&listScreenshotUsage)
	if errorListScreenshotUsage != nil {
		log.Println("errorListScreenshotUsage", errorListScreenshotUsage)
	}

	for _, screenshotUsage := range listScreenshotUsage {
		//get subscription id
		subscriptionId := getSubscriptionIdFromUserId(db, screenshotUsage.UserId)
		subscriptionItem := getSubscriptionItem(db, subscriptionId)
		if subscriptionItem == nil {
			log.Println("subscriptionItem is nil")
			return
		}
		//loop through the subscriptionItem.Items.Data find .Plan.Amount == null
		var planId string
		for _, item := range subscriptionItem.Items.Data {
			if item.Plan.Amount == 0 {
				planId = item.ID
			}
		}
		log.Println("Found item : ", planId)
		// * send report usage to stripe
		sendReportUsage(db, screenshotUsage, planId)
		// * reset extra screenshot usage
		resetExtraScreenshotUsage(db, screenshotUsage)
		// * update screenshot extra usage history
		updateScreenshotUsageExtraHistory(db, screenshotUsage, planId)
	}
}

func resetExtraScreenshotUsage(db dbx.Builder, screenshotUsage module.ScreenshotUsage) {
	_, errorUpdateScreenshotUsage := db.NewQuery(`
		UPDATE screenshot_usage
		SET extra_screenshots_taken = 0
		WHERE user_id = {:user_id}
	`).Bind(dbx.Params{
		"user_id": screenshotUsage.UserId,
	}).Execute()

	if errorUpdateScreenshotUsage != nil {
		log.Println("errorUpdateScreenshotUsage", errorUpdateScreenshotUsage)
	}
}

func updateScreenshotUsageExtraHistory(db dbx.Builder, screenshotUsage module.ScreenshotUsage, planId string) {
	_, errorUpdateScreenshotUsage := db.NewQuery(`
		INSERT INTO extra_usage_histories (extra_screenshots_takend,user_id, subscription_item_id)
		VALUES ({:extra_screenshots_taken}, {:user_id})
	`).Bind(dbx.Params{
		"extra_screenshots_taken": screenshotUsage.ExtraScreenshotsTaken,
		"user_id":                 screenshotUsage.UserId,
		"subscription_item_id":    planId,
	}).Execute()

	if errorUpdateScreenshotUsage != nil {
		log.Println("errorUpdateScreenshotUsage", errorUpdateScreenshotUsage)
	}
}

func sendReportUsage(db dbx.Builder, screenshotUsage module.ScreenshotUsage, planId string) {
	params := &stripe.UsageRecordParams{
		SubscriptionItem: stripe.String("si_P5JaH9qiRpP9Nv"),
		Quantity:         stripe.Int64(screenshotUsage.ExtraScreenshotsTaken),
		Timestamp:        stripe.Int64(time.Now().Unix()),
		Action:           stripe.String(string(stripe.UsageRecordActionSet)),
	}

	usagerecord.New(params)
}

func getSubscriptionIdFromUserId(db dbx.Builder, userId string) string {
	userData := module.User{}

	errorUserData := db.Select("stripe_subscription_id").From("users").Where(
		dbx.NewExp("id = {:id}", dbx.Params{"id": userId}),
	).One(&userData)

	if errorUserData != nil {
		log.Println("errorUserData", errorUserData)
	}

	subscriptionId := userData.StripeSubscriptionId

	return subscriptionId
}

func getSubscriptionItem(db dbx.Builder, subscriptionId string) *stripe.Subscription {

	params := &stripe.SubscriptionParams{}
	result, err := sub.Get(subscriptionId, params)
	if err != nil {
		log.Println("err", err)
	}

	return result
}
