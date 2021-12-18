package db

import (
	"encoding/json"
	"fmt"

	"github.com/taskdb/bolt"
	"github.com/taskpoc/model"
)

func GetTaskDetails(id string) ([]model.Task, error) {
	var tasks []model.Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(dbname)).Bucket([]byte("Task"))
		matched := false
		b.ForEach(func(k, v []byte) error {
			if !matched {
				var task model.Task
				if id == string(k) {
					matched = true
					err := json.Unmarshal(v, &task)
					if err != nil {
						return err
					}
					tasks = append(tasks, task)
				} else if id == "" {
					err := json.Unmarshal(v, &task)
					if err != nil {
						return err
					}
					tasks = append(tasks, task)
				}
				return nil
			} else {
				return nil
			}
		})
		if len(tasks) == 0 {
			return fmt.Errorf("Data not found!")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func InsertTaskDetails(task model.Task) error {
	data, err := json.Marshal(task)
	if err != nil {
		return err
	}
	err = db.Update(func(tx *bolt.Tx) error {

		err := tx.Bucket([]byte(dbname)).Bucket([]byte("TASK")).Put([]byte(task.Id), []byte(data))
		if err != nil {
			return fmt.Errorf("could not insert Task: %v", err)
		}
		return nil
	})
	fmt.Println("Added Task")
	return err
}
