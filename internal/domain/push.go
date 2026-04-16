package domain

import "time"

type Push struct {
	ID        string
	UserID    string
	Message   string
	Status    string
	Attempts  int
	LastError string
	CreatedAt time.Time
	SentAt    time.Time
}
