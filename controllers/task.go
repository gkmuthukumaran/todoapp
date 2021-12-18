package controllers

import (
	"net/http"

	"github.com/taskpoc/interfaces/db"
	"github.com/taskpoc/model"
	"github.com/taskpoc/utils"
	"github.com/fatih/structs"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type TaskList struct{
	db *sql.DB
	}
	
	func NewTaskListController(db *sql.DB) *TaskList {
	
		controller := new(TaskList)
		controller.db = db
		return controller
		
	}
	
	func (h *TaskList) GetTaskById(ctx echo.Context) error{
	
		staffid := ctx.Param("id")
		log.Println("---", staffid)
	
		respArray,err := models.GetTaskById(staffid)
	
		if err != nil{
			return ctx.JSON(http.StatusInternalServerError, utils.BldGnrRsp(500, err.Error(), nil))
		}
	
		if  len(respArray) == 0{
			return ctx.JSON(http.StatusOK, utils.BldGnrRsp(200, " No Data Found ", nil))
		}
	
		respInterface := make([]interface{}, len(respArray))
		for i, evc := range respArray {
			mapForm := (structs.New(evc)).Map()
			respInterface[i] = mapForm
		}
		return ctx.JSON(http.StatusOK, utils.BldGnrRsp(200, "Read Successfully",&respInterface))
	}
	
	func (h *TaskList) GetTasksByCategory(ctx echo.Context) error{
	
		staffid := ctx.Param("id")
		log.Println("---", staffid)
	
		respArray,err := models.GetTasksByCategory(category)
	
		if err != nil{
			return ctx.JSON(http.StatusInternalServerError, utils.BldGnrRsp(500, err.Error(), nil))
		}
	
		if  len(respArray) == 0{
			return ctx.JSON(http.StatusOK, utils.BldGnrRsp(200, " No Data Found ", nil))
		}
	
		respInterface := make([]interface{}, len(respArray))
		for i, evc := range respArray {
			mapForm := (structs.New(evc)).Map()
			respInterface[i] = mapForm
		}
		return ctx.JSON(http.StatusOK, utils.BldGnrRsp(200, "Read Successfully",&respInterface))
	}


func (a *Task) InsertTaskDetails() (err error){
	db := sqldb.CreateCon()

	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO task(task_code, task, content, category ) VALUES(?,?,?,?)")
	if err != nil {
		return
	}
	_, err = stmt.Exec(a.taskCode,a.task, a.content, a.Category)
	if err != nil {
		return
	}

	err = tx.Commit()
	if err != nil {
		return
	}
	return nil
}