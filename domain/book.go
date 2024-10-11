package domain

import "time"

type Document struct {
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
