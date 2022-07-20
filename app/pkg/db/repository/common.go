package repository

import (
	"echosample/pkg/management"
	"echosample/pkg/log"
)

var (
	db = management.GetDBManager().GetDB()
	logger = log.GetInstance("repository")
)

