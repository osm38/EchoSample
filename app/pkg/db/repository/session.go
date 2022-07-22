package repository

import (
	"echosample/pkg/db/model"
	"time"
)

func RegisterSession(sessID string, userID int) model.Session {
	s := createRegisterModel()
	s.ID = sessID
	s.UserID = userID
	result := db.Create(&s)
	if result.Error != nil {
		logger.Error(result.Error, "RegisterSession is Fail!!", nil)
		panic("RegisterSession is Fail!!")
	}

	return s
}

func FindSessionByID(sessID string) model.Session {
	condition := model.Session{
		ID: sessID,
		Deleted: false,
	}

	s := model.Session{}

	result := db.Where(&condition).First(&s)
	if result.Error != nil {
		logger.Error(result.Error, "SearchSession is Fail!!", nil)
		panic("SearchSession is Fail!!")
	}

	return s;
}

func createRegisterModel() model.Session {
	t := time.Now()
	m := model.Session{
		MaxAge: 86400 * 7,
		Deleted: false,
		CreateUser: "system",
		CreateTimestamp: t,
		UpdateUser: "system",
		UpdateTimestamp: t,
	}

	return m
}
