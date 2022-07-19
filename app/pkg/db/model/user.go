package model

import (
	"time"
)

type User struct {
	ID int
	Name string
	Age int
	Password string
	LoginTimestamp time.Time
	LogoutTimestamp time.Time
	Deleted bool
	CreateUser string
	CreateTimestamp time.Time
	UpdateUser string
	UpdateTimestamp time.Time
}

type UserPK struct {
	ID int
	Deleted bool
}