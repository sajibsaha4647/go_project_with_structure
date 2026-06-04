package model

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token	string `json:"token,omitempty"`
}


func ValidateLogin(userLogin UserLogin) bool {
	user := FindUserByEmail(userLogin.Email)
	if user.Id != 0 && user.Password == userLogin.Password {
		return true
	}
	return false
}