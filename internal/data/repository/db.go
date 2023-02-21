package repository

import (
	"context"
	"database/sql"
	"github.com/mustafaerbay/cleango/internal/biz/user"
	"fmt"
)

type DBRepository struct {
	db *sql.DB
}

func NewDBRepository(db *sql.DB) *DBRepository {
	return &DBRepository{
		db: db,
	}
}

func (r *DBRepository) ListUsers() ([]*user.User, error) {
	users := make([]*user.User, 0)

	rows, err := r.db.Query("SELECT id, first_name, last_name, email FROM users")
	if err != nil {
		return nil, fmt.Errorf("error listing users: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var u user.User
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email); err != nil {
			return nil, fmt.Errorf("error scanning user row: %v", err)
		}
		users = append(users, &u)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating user rows: %v", err)
	}

	return users, nil
}

func (r *DBRepository) GetUser(id string) (*user.User, error) {
	var u user.User

	row := r.db.QueryRow("SELECT id, first_name, last_name, email FROM users WHERE id=?", id)
	if err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, user.ErrUserNotFound
		}
		return nil, fmt.Errorf("error getting user: %v", err)
	}

	return &u, nil
}

func (r *DBRepository) CreateUser(u *user.User) error {
	result, err := r.db.Exec("INSERT INTO users (id, first_name, last_name, email) VALUES (?, ?, ?, ?)",
		u.ID, u.FirstName, u.LastName, u.Email)
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}
	if rowsAffected != 1 {
		return fmt.Errorf("unexpected rows affected: %d", rowsAffected)
	}

	return nil
}

func (r *DBRepository) UpdateUser(u *user.User) error {
	result, err := r.db.Exec("UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?",
		u.FirstName, u.LastName, u.Email, u.ID)
	if err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}
	if rowsAffected != 1 {
		return fmt.Errorf("unexpected rows affected: %d", rowsAffected)
	}

	return nil
}

func (r *DBRepository) DeleteUser(id string) error {
	result, err := r.db.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}
	if rowsAffected != 1 {
		return fmt.Errorf("unexpected rows affected: %d", rowsAffected)
	}

	return nil
}

