package controllers

import (
	"fmt"
	"net/http"

	"github.com/fatih/structs"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	sqldb "github.com/todoapp/interfaces/db"
	"github.com/todoapp/model"
	"github.com/todoapp/utils"
)

func GetTaskById(ctx echo.Context) error {

	taskid := ctx.Param("id")
	respArray := []model.Task{}
	var err error

	respArray, err = sqldb.GetTaskDetails(taskid)
	log.Println("---", taskid)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.BldGnrRsp(500, err.Error(), nil))
	}

	if len(respArray) == 0 {
		return ctx.JSON(http.StatusOK, utils.BldGnrRsp(200, " No Data Found ", nil))
	}

	respInterface := make([]interface{}, len(respArray))
	for i, evc := range respArray {
		mapForm := (structs.New(evc)).Map()
		respInterface[i] = mapForm
	}
	return ctx.JSON(http.StatusOK, utils.BldGnrRsp(200, "Read Successfully", &respInterface))
}

func UpdateTaskById(ctx echo.Context) error {

	taskid := ctx.Param("id")
	taskstatus := ctx.QueryParam("status")
	var err error
	// fmt.Println("teste" + taskid)
	err = sqldb.UpdateTaskDetails(taskid, taskstatus)

	if err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, "Updation Failed : "+err.Error())
	}
	return ctx.JSON(http.StatusOK, "Updation Success")
}

func DeleteTask(ctx echo.Context) error {

	taskid := ctx.Param("id")
	var err error

	err = sqldb.DeletTaskDetails(taskid)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.BldGnrRsp(500, err.Error(), nil))
	}

	return ctx.JSON(http.StatusOK, utils.BldGnrRsp(200, "Deleted Successfully", nil))
}

func GetAllTask(ctx echo.Context) error {

	respArray := []model.Task{}
	var err error

	respArray, err = sqldb.GetAllTask()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.BldGnrRsp(500, err.Error(), nil))
	}

	if len(respArray) == 0 {
		return ctx.JSON(http.StatusOK, utils.BldGnrRsp(200, " No Data Found ", nil))
	}

	respInterface := make([]interface{}, len(respArray))
	for i, evc := range respArray {
		mapForm := (structs.New(evc)).Map()
		respInterface[i] = mapForm
	}
	return ctx.JSON(http.StatusOK, utils.BldGnrRsp(200, "Read Successfully", &respInterface))
}

func GetAllCategory(ctx echo.Context) error {

	respArray := []model.Category{}
	var err error

	respArray, err = sqldb.GetAllCategory()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.BldGnrRsp(500, err.Error(), nil))
	}

	if len(respArray) == 0 {
		return ctx.JSON(http.StatusOK, utils.BldGnrRsp(200, " No Data Found ", nil))
	}

	respInterface := make([]interface{}, len(respArray))
	for i, evc := range respArray {
		mapForm := (structs.New(evc)).Map()
		respInterface[i] = mapForm
	}
	return ctx.JSON(http.StatusOK, utils.BldGnrRsp(200, "Read Successfully", &respInterface))
}

func GetTasksByCategory(ctx echo.Context) error {

	category := ctx.Param("category")
	respArray := []model.Task{}
	var err error

	respArray, err = sqldb.GetTasksByCategory(category)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.BldGnrRsp(500, err.Error(), nil))
	}

	if len(respArray) == 0 {
		return ctx.JSON(http.StatusOK, utils.BldGnrRsp(200, " No Data Found ", nil))
	}

	respInterface := make([]interface{}, len(respArray))
	for i, evc := range respArray {
		mapForm := (structs.New(evc)).Map()
		respInterface[i] = mapForm
	}
	return ctx.JSON(http.StatusOK, utils.BldGnrRsp(200, "Read Successfully", &respInterface))
}

func InsertTaskDetails(ctx echo.Context) error {

	taskDetail := new(model.Task)

	if err := ctx.Bind(taskDetail); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, "Unable to bind request body")
	}
	fmt.Println(taskDetail)
	taskDetail.Id = uuid.New().String()
	err := sqldb.InsertTaskDetails(*taskDetail)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, "Insertion Failed : "+err.Error())
	}
	return ctx.JSON(http.StatusOK, "Insertion Success")
}
