package model

import (
	"time"
)

type Session struct {
	ID string
	UserID int
	MaxAge int
	Deleted bool
	CreateUser string
	CreateTimestamp time.Time
	UpdateUser string
	UpdateTimestamp time.Time
}