package domain

import "time"

type Book struct {
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
