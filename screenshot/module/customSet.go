package module

type CustomSet struct {
	// Model
	Id              string `json:"id" db:"id"`
	Name            string `json:"name" db:"name"`
	UserId          string `json:"user_id" db:"user_id"`
	CSS             string `json:"css" db:"css"`
	JavaScript      string `json:"javascript" db:"javascript"`
	Cookies         string `json:"cookies" db:"cookies"`
	LocalStorage    string `json:"localStorage" db:"localStorage"`
	UserAgent       string `json:"user_agent" db:"user_agent"`
	Headers         string `json:"headers" db:"headers"`
	BucketEndpoint  string `json:"bucket_endpoint" db:"bucket_endpoint"`
	BucketDefault   string `json:"bucket_default" db:"bucket_default"`
	BucketAccessKey string `json:"bucket_access_key" db:"bucket_access_key"`
	BucketSecretKey string `json:"bucket_secret_key" db:"bucket_secret_key"`
}
