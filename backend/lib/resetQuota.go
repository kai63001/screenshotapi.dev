package lib

import (
	"backend/module"
	"log"
	"time"

	"github.com/pocketbase/dbx"
)

// this function will get all user and reset quota per month
func ResetQuotaPerMonth(db dbx.Builder) {
	screenshotUsage := []module.ScreenshotUsage{}
	errUser := db.NewQuery(`
		SELECT user_id, next_reset_quota
		FROM screenshot_usage
	`).All(&screenshotUsage)

	if errUser != nil {
		log.Println("errUser", errUser)
	}

	log.Println("screenshotUsage", screenshotUsage)

	for _, sc := range screenshotUsage {
		// get created date is less than or equal to 30 days
		if CheckDate(sc.NextResetQuota) {
			// reset quota
			_, errResetQuota := db.NewQuery(`
				UPDATE screenshot_usage
				SET screenshots_taken = 0, next_reset_quota = {:next_reset_quota}
				WHERE user_id = {:user_id}
			`).Bind(dbx.Params{
				"user_id": sc.UserId,
				// set next reset quota to 30 days
				"next_reset_quota": GetNextResetQuota(),
			}).Execute()

			if errResetQuota != nil {
				log.Println("errResetQuota", errResetQuota)
			}

			log.Println("reset quota for user", sc.UserId)
		}
	}
}

func CheckDate(nextResetQuota string) bool {

	//convert string to time
	nowNextResetQuota, err := time.Parse("2006-01-02 15:04:05.999999999-07:00", nextResetQuota)
	if err != nil {
		log.Println("err", err)
	}

	now := time.Now()
	return nowNextResetQuota.Year() == now.Year() && nowNextResetQuota.Month() == now.Month() && nowNextResetQuota.Day() <= now.Day()
}

func GetNextResetQuota() time.Time {
	now := time.Now()
	return now.AddDate(0, 1, 0)
}
