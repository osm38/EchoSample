package repository

import (
	"echosample/pkg/db/model"
	"time"
)

func RegisterSession(sessID string, userID int) {
	u := createModel()
	u.ID = sessID
	u.UserID = userID
	result := db.Create(&u)
	if result.Error != nil {
		logger.Error(result.Error, "RegisterSession is Fail!!", nil)
		panic("RegisterSession is Fail!!")
	}
}

func createModel() model.Session {
	t := time.Now()
	m := model.Session{
		StartTime: t,
		Deleted: false,
		CreateUser: "system",
		CreateTimestamp: t,
		UpdateUser: "system",
		UpdateTimestamp: t,
	}

	return m
}