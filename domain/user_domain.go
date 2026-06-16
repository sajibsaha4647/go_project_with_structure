package domain


type User struct {
	Id       int    `json:"id"       db:"id"`
	Name     string `json:"name"     db:"name"`
	Email    string `json:"email"    db:"email"`
	Password string `json:"password" db:"password"`
	UserType string `json:"userType" db:"user_type"`
}


type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token,omitempty"`
}

