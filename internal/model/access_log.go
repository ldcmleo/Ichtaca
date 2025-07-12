package model

import "time"

type AccessLog struct {
	ID            int
	UserID        int
	RepoRequested string
	Package       string
	Result        string
	ClientIP      string
	TimeStamp     time.Time
}
