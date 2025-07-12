package model

import "time"

type User struct {
	ID          string
	Name        string
	LastName    string
	Email       string
	CommonName  string
	FingerPrint string
	IsAdmin     bool
	Revoked     bool
	CreatedAt   time.Time
}
