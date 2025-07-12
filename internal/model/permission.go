package model

import "time"

type Permission struct {
	ID        int
	UserID    int
	Scope     string
	CreatedAt time.Time
}
