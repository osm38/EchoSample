package repository

import (
	"echosample/pkg/db/model"
	"fmt"
)

func FindUser() model.User {
	user := model.User{}
	db.First(&user)
	fmt.Println(user.ID)
	fmt.Println(user.Name)
	fmt.Println(user.Age)

	return user
}

func FindUserByID(id int) model.User {
	condition := model.User{
		ID: 1,
		Deleted: false,
	}
	var users []model.User

	db.Where(&condition).Find(&users)
	fmt.Println("result: ", len(users))

	for i := 0; i < len(users); i++ {
		fmt.Println("-- Start ", i, " --")
		user := users[i]
		fmt.Println("ID :", user.ID)
		fmt.Println("Name :", user.Name)
		fmt.Println("Age :", user.Age)
		fmt.Println("Password :", user.Password)
		fmt.Println("LoginTimeStamp :", user.LoginTimestamp)
		fmt.Println("LogoutTimeStamp :", user.LogoutTimestamp)
		fmt.Println("Deleted :", user.Deleted)
		fmt.Println("CreateUser :", user.CreateUser)
		fmt.Println("CreateTimeStamp :", user.CreateTimestamp)
		fmt.Println("UpdateUser :", user.UpdateUser)
		fmt.Println("UpdateTimeStamp :", user.UpdateTimestamp)
		fmt.Println("-- End ", i, " --")
	}

	return users[0]
}

func FindUserByName(name string) model.User {
	condition := model.User{
		Name: name,
		Deleted: false,
	}

	var user model.User

	result := db.Where(&condition).First(&user)

	if result.Error != nil {
		logger.Error(result.Error, "FindUserByName is Fail!!", nil)
		panic("FindUserByName is Fail!!")
	}

	return user
}