package module

type ScreenshotUsage struct {
	//id is string
	UserId          string `json:"user_id" db:"user_id"`
	ScreenshotTaken int64  `json:"screenshots_taken" db:"screenshots_taken"`
}

type GetQuotaScreenshot struct {
	ScreenshotUsage
	SubscriptionPlans
}
