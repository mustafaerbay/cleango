package repository

import (
	"github.com/mustafaerbay/cleango/internal/biz/user"
	"github.com/google/uuid"
)

/*
This example shows a simple implementation of the user.Repository interface, which defines the persistence layer for handling CRUD operations on users. The UserRepository struct contains a slice of users, which is used to store the user data.

The NewUserRepository function creates a new instance of the UserRepository struct, and the various methods on the UserRepository struct implement the functionality specified in the Repository interface.

Note that this example assumes that the user package has already been defined, and that the necessary dependencies have been imported. In a real-world application, the UserRepository may interact with a database or other external data source to persist and retrieve user data.
*/
type UserRepository struct {
	users []*user.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make([]*user.User, 0),
	}
}

func (r *UserRepository) ListUsers() ([]*user.User, error) {
	return r.users, nil
}

func (r *UserRepository) GetUser(id string) (*user.User, error) {
	for _, u := range r.users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, user.ErrUserNotFound
}

func (r *UserRepository) CreateUser(u *user.User) error {
	u.ID = uuid.New().String()
	r.users = append(r.users, u)
	return nil
}

func (r *UserRepository) UpdateUser(u *user.User) error {
	for i, existingUser := range r.users {
		if existingUser.ID == u.ID {
			r.users[i] = u
			return nil
		}
	}
	return user.ErrUserNotFound
}

func (r *UserRepository) DeleteUser(id string) error {
	for i, existingUser := range r.users {
		if existingUser.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return user.ErrUserNotFound
}
