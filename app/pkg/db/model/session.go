package model

import (
	"time"
)

type Session struct {
	ID string
	UserID int
	StartTime time.Time
	Deleted bool
	CreateUser string
	CreateTimestamp time.Time
	UpdateUser string
	UpdateTimestamp time.Time
}