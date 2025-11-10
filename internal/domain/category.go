package domain

import "time"

type Category struct {
	ID        int
	NameTK    string
	NameEN    string
	NameRU    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
