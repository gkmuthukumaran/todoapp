package db

import (
	"encoding/json"
	"fmt"

	"github.com/todoapp/model"
)

func initialiseUserDetails() {

	user, err := GetUserDetails("todoTokenAuth")

	fmt.Println("testSetupInit", err)
	if err == nil && user != nil && user.Username != "" {
		return
	}

	newuser := model.User{

		Username: "todoTokenAuth",

		Password: "p@ss1234",
	}
	fmt.Println("beforeInsert")
	InsertUserDetails(newuser)

}

func InsertUserDetails(user model.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	fmt.Println(data)
	fmt.Println("Added User")
	return err
}

func GetUserDetails(username string) (*model.User, error) {
	var user model.User
	fmt.Println(username)
	fmt.Println(user)
	return &user, nil
}
