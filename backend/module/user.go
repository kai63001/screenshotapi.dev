package module

type User struct {
	Id       string `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type UserForKey struct {
	UserId string `json:"user_id"`
}
