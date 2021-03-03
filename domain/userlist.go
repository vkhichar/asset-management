package domain

type UserList struct {
	ID    int    `db:"int"`
	Name  string `db:"name"`
	Email string `db:"email"`
	// Created_at time.Date   `db:"created_at"`
	// Updated_at Date   `db:"updated_at"`
}
