package user

import "ecommerce/domain"

type Service interface {
	Login(email, password string) (*domain.User, bool, error)
	CreateUser(u domain.User) (*domain.User, error)
	GetAllUsers() ([]domain.User, error)
	GetUserById(id int) *domain.User
	UpdateUserById(id int, u domain.User) *domain.User
	DeleteUserById(id int) bool
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Login(email, password string) (*domain.User, bool, error) {
	u := s.repo.FindUserByEmail(email)
	if u == nil {
		return nil, false, nil
	}
	return u, u.Password == password, nil
}

func (s *service) CreateUser(u domain.User) (*domain.User, error) {
	return s.repo.Store(u)
}

func (s *service) GetAllUsers() ([]domain.User, error) {
	return s.repo.GetAllUsers()
}

func (s *service) GetUserById(id int) *domain.User {
	return s.repo.GetUserById(id)
}

func (s *service) UpdateUserById(id int, u domain.User) *domain.User {
	return s.repo.UpdateUserById(id, u)
}

func (s *service) DeleteUserById(id int) bool {
	return s.repo.DeleteUserById(id)
}
