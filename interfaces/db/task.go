package db

import (
	"fmt"
	"log"

	// "github.com/taskdb/bolt"
	"github.com/todoapp/model"
)

func GetAllTask() ([]model.Task, error) {
	db := CreateCon()
	queryResult, err := db.Query("SELECT id, task, content, category, status from taskList")

	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	res := []model.Task{}

	for queryResult.Next() {
		task := model.Task{}
		err := queryResult.Scan(&task.Id, &task.Task, &task.Content, &task.Category, &task.Status)
		log.Println("beforIff-----", err)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		res = append(res, task) //return datalist
	}
	return res, nil
}

func GetAllCategory() ([]model.Category, error) {
	db := CreateCon()
	queryResult, err := db.Query("SELECT id, category from category")

	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	res := []model.Category{}

	for queryResult.Next() {
		task := model.Category{}
		err := queryResult.Scan(&task.Id, &task.Category)
		log.Println("beforIff-----", err)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		res = append(res, task) //return datalist
	}
	return res, nil
}

func GetTasksByCategory(category string) ([]model.Task, error) {
	db := CreateCon()
	queryResult, err := db.Query("SELECT id, task, content, category, status from taskList where category = '" + category + "'")

	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	res := []model.Task{}

	for queryResult.Next() {
		task := model.Task{}
		err := queryResult.Scan(&task.Id, &task.Task, &task.Content, &task.Category, &task.Status)
		log.Println("beforIff-----", err)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		res = append(res, task) //return datalist
	}
	return res, nil
}

func GetTaskDetails(id string) ([]model.Task, error) {
	db := CreateCon()
	task := model.Task{}
	err := db.QueryRow("SELECT id, task, content, category, status from taskList where id='"+id+"'").Scan(&task.Id, &task.Task, &task.Content, &task.Category, &task.Status)

	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	} else {
		return []model.Task{task}, nil
	}
}

func DeletTaskDetails(id string) error {
	db := CreateCon()
	_, err := db.Exec("DELETE from taskList where id='" + id + "'")

	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		return err
	} else {
		return nil
	}
}

func InsertTaskDetails(t model.Task) (err error) {

	db := CreateCon()

	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
	}

	stmt, err := tx.Prepare("INSERT INTO taskList(id, task, content, category, status ) VALUES(?,?,?,?,?)")
	if err != nil {
		return
	}
	_, err = stmt.Exec(t.Id, t.Task, t.Content, t.Category, t.Status)
	if err != nil {
		return
	}

	err = tx.Commit()
	if err != nil {
		return
	}
	return nil
}

func UpdateTaskDetails(id string, status string) (err error) {

	db := CreateCon()
	fmt.Println(status)
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
	}

	stmt, err := tx.Prepare("Update taskList set status = ? where id = ?")
	if err != nil {
		return
	}
	_, err = stmt.Exec(status, id)
	if err != nil {
		return
	}

	err = tx.Commit()
	if err != nil {
		return
	}
	return nil
}
