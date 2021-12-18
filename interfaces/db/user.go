package db

import (
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/taskpoc/model"
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
	err = db.Update(func(tx *bolt.Tx) error {

		err := tx.Bucket([]byte(dbname)).Bucket([]byte("USER")).Put([]byte(user.Username), []byte(data))
		fmt.Println("insertTest", err)
		if err != nil {
			return fmt.Errorf("could not insert User: %v", err)
		}
		return nil
	})
	fmt.Println("Added User")
	return err
}

func GetUserDetails(username string) (*model.User, error) {
	var user model.User
	fmt.Println(username)
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(dbname)).Bucket([]byte("USER"))
		fmt.Println(b)
		b.ForEach(func(k, v []byte) error {
			fmt.Println("testLog", string(k), string(v))
			if username == string(k) {
				err := json.Unmarshal(v, &user)
				if err != nil {
					return err
				}
			}
			return nil
		})
		if len(username) == 0 {
			return fmt.Errorf("Data not found!")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	return &user, nil
}
