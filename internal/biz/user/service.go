package user

import "github.com/google/uuid"
/*
This example shows a simple implementation of the user.Service interface,
 which defines the business logic layer for handling CRUD operations on users. 
 The userService struct contains a store object of type Store, which is used to interact with 
 the persistence layer.

The NewService function creates a new instance of the userService struct, 
and the various methods on the userService struct implement the functionality specified 
in the Service interface.

Note that this example assumes that the user.Store interface has already been defined 
in the internal/biz/user/store.go file, and that the user.User model object has already 
been defined in the internal/biz/user/model.go file. The userService simply acts as an adapter 
between the business logic layer and the persistence layer, 
which is implemented in the internal/biz/user package.
*/
type Service interface {
	ListUsers() ([]*User, error)
	GetUser(id string) (*User, error)
	CreateUser(u *User) error
	UpdateUser(id string, u *User) error
	DeleteUser(id string) error
}

type userService struct {
	store Store
}

func NewService(store Store) Service {
	return &userService{store: store}
}

func (s *userService) ListUsers() ([]*User, error) {
	return s.store.ListUsers()
}

func (s *userService) GetUser(id string) (*User, error) {
	return s.store.GetUser(id)
}

func (s *userService) CreateUser(u *User) error {
	u.ID = uuid.New().String()
	return s.store.CreateUser(u)
}

func (s *userService) UpdateUser(id string, u *User) error {
	existingUser, err := s.store.GetUser(id)
	if err != nil {
		return err
	}

	u.ID = existingUser.ID

	return s.store.UpdateUser(u)
}

func (s *userService) DeleteUser(id string) error {
	return s.store.DeleteUser(id)
}
