package repository

import (
	"echosample/pkg/db/model"
	"echosample/pkg/management"
	"fmt"
)

func FindUser() model.User {
	user := model.User{}
	db := management.GetDBManager().GetDB()
	db.First(&user)
	fmt.Println(user.ID)
	fmt.Println(user.Name)
	fmt.Println(user.Age)

	return user
}