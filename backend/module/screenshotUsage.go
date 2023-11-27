package module

type ScreenshotUsage struct {
	//id is string
	UserId          string `json:"user_id" db:"user_id"`
	ScreenshotTaken int64  `json:"screenshots_taken" db:"screenshots_taken"`
	NextResetQuota  string `json:"next_reset_quota" db:"next_reset_quota"`
}

type GetQuotaScreenshot struct {
	ScreenshotUsage
	SubscriptionPlans
}
