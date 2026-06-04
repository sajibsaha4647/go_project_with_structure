package model

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"userType"`
}

var userList []User

func StoreUser(u User) {
	if len(userList) == 0 {
		u.Id = 1
	} else {
		u.Id = userList[len(userList)-1].Id + 1
	}
	userList = append(userList, u)
}

func GetAllUsers() []User {
	return userList
}

func GetUserById(id int) User {
	for _, user := range userList {
		if id == user.Id {
			return user
		}
	}
	return User{}
}

func UpdateUser(id int, updatedUser User) User {
	for i, user := range userList {
		if id == user.Id {
			userList[i] = updatedUser
			return updatedUser
		}
	}
	return User{}
}

func DeleteUser(id int) bool {
	for i, user := range userList {
		if id == user.Id {
			userList = append(userList[:i], userList[i+1:]...)
			return true
		}
	}
	return false
}

func FindUserByEmail(email string) User {
   for _, user := range userList {
	if user.Email == email {
		return user
	}
   }
   return User{}
}
