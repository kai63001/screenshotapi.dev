package module

type ScreenshotUsage struct {
	//id is string
	UserId                string `json:"user_id" db:"user_id"`
	ScreenshotTaken       int64  `json:"screenshots_taken" db:"screenshots_taken"`
	NextResetQuota        string `json:"next_reset_quota" db:"next_reset_quota"`
	DisableExtra          bool   `json:"activate_extra" db:"disable_extra"`
	ExtraScreenshotsTaken int64  `json:"extra_screenshots_taken" db:"extra_screenshots_taken"`
}

type GetQuotaScreenshot struct {
	ScreenshotUsage
	SubscriptionPlans
}
