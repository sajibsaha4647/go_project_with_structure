package repo

import (
	"ecommerce/model"

	"github.com/jmoiron/sqlx"
)

// type UserRepo interface {
// 	Store(model.User) (*model.User, error)
// 	GetAllUsers()  ([]model.User, error)
// 	GetUserById(id int) *model.User
// 	UpdateUserById(id int, updatedUser model.User) *model.User
// 	DeleteUserById(id int) bool
// 	FindUserByEmail(email string) *model.User
// }

type UserRepo interface {
	Store(model.User) (*model.User, error)
	GetAllUsers() ([]model.User, error) // add error
	GetUserById(id int) *model.User
	UpdateUserById(id int, updatedUser model.User) *model.User // keep pointer
	DeleteUserById(id int) bool
	FindUserByEmail(email string) *model.User // add missing method
}

type userRepo struct {
	// userList []model.User
	dbConnector *sqlx.DB
}

func NewUserRepo(dbCon *sqlx.DB) UserRepo {
	return &userRepo{
		dbConnector: dbCon,
	}
}

func (u *userRepo) Store(user model.User) (*model.User, error) {
	query := `
        INSERT INTO users (name, email, password, user_type)
        VALUES (:name, :email, :password, :user_type)
        RETURNING id
    `
	var userId int
	row, err := u.dbConnector.NamedQuery(query, user)
	if err != nil {
		return nil, err
	}
	if row.Next() {
		err = row.Scan(&userId)
		if err != nil {
			return nil, err
		}
	}
	user.Id = userId
	return &user, nil
}

func (u *userRepo) GetAllUsers() ([]model.User, error) {
	var users []model.User
	query := `SELECT id, name, email, password, user_type from users`
	err := u.dbConnector.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (u *userRepo) GetUserById(id int) *model.User {
	// for _, user := range u.userList {
	// 	if id == user.Id {
	// 		return user
	// 	}
	// }
	// return model.User{}

	query := `SELECT id, name, email, password, user_type from users WHERE id = $1`
	var user model.User
	err := u.dbConnector.Get(&user, query, id)
	if err != nil {
		return nil
	}
	return &user
}
func (u *userRepo) UpdateUserById(id int, updatedUser model.User) *model.User {
	query := `UPDATE users SET name = $1, email = $2, password = $3, user_type = $4 WHERE id = $5`
	_, err := u.dbConnector.Exec(query, updatedUser.Name, updatedUser.Email, updatedUser.Password, updatedUser.UserType, id)
	if err != nil {
		return nil
	}
	updatedUser.Id = id
	return &updatedUser
}

func (u *userRepo) DeleteUserById(id int) bool {
	// for i, user := range u.userList {
	// 	if id == user.Id {
	// 		u.userList = append(u.userList[:i], u.userList[i+1:]...)
	// 		return true
	// 	}
	// }
	// return false

	query := `DELETE FROM users WHERE id = $1`
	result, error := u.dbConnector.Exec(query, id)
	if error != nil {
		return false
	}
	deleteid, err := result.RowsAffected()
	if err != nil {
		return false
	}
	return deleteid > 0
}

func (u *userRepo) FindUserByEmail(email string) *model.User {
	// for _, user := range u.userList {
	// 	if email == user.Email {
	// 		return user
	// 	}
	// }
	// return model.User{}
	query := `SELECT id, name, email,password, user_type from users WHERE EMAIL = $1`
	var user model.User
	err := u.dbConnector.Get(&user, query, email)
	if err != nil {
		return nil
	}
	return &user
}
