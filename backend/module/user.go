package module

type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserForKey struct {
	UserId string `json:"user_id"`
}
