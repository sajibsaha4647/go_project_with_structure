package user

import (
	"ecommerce/domain"

	"github.com/jmoiron/sqlx"
)

type postgresRepository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Store(user domain.User) (*domain.User, error) {
	query := `INSERT INTO users (name, email, password, user_type) VALUES (:name, :email, :password, :user_type) RETURNING id`
	row, err := r.db.NamedQuery(query, user)
	if err != nil {
		return nil, err
	}
	if row.Next() {
		if err = row.Scan(&user.Id); err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func (r *postgresRepository) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Select(&users, `SELECT id, name, email, password, user_type FROM users`)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *postgresRepository) GetUserById(id int) *domain.User {
	var user domain.User
	err := r.db.Get(&user, `SELECT id, name, email, password, user_type FROM users WHERE id = $1`, id)
	if err != nil {
		return nil
	}
	return &user
}

func (r *postgresRepository) UpdateUserById(id int, u domain.User) *domain.User {
	query := `UPDATE users SET name = $1, email = $2, password = $3, user_type = $4 WHERE id = $5`
	_, err := r.db.Exec(query, u.Name, u.Email, u.Password, u.UserType, id)
	if err != nil {
		return nil
	}
	u.Id = id
	return &u
}

func (r *postgresRepository) DeleteUserById(id int) bool {
	result, err := r.db.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return false
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return false
	}
	return rows > 0
}

func (r *postgresRepository) FindUserByEmail(email string) *domain.User {
	var user domain.User
	err := r.db.Get(&user, `SELECT id, name, email, password, user_type FROM users WHERE email = $1`, email)
	if err != nil {
		return nil
	}
	return &user
}
