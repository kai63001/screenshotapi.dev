package module

type SubscriptionPlans struct {
	Name                string `json:"name" db:"name"`
	IncludedScreenshots int64  `json:"included_screenshots" db:"included_screenshots"`
	RateLimitPerMinute  int64  `json:"rate_limit_per_minute" db:"rate_limit_per_minute"`
}
