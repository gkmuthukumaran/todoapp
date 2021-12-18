package controllers

import (
	"net/http"

	"github.com/blogpoc/interfaces/db"
	"github.com/fatih/structs"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/taskapp/model"
	"github.com/taskapp/utils"
)

func GetTasks(c echo.Context) error {
	id := c.Param("id")
	data, err := db.GetTaskDetails(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.BldGnrRsp(http.StatusNotFound, err.Error(), nil))
	}

	respInterface := make([]interface{}, len(data))
	for i, evc := range data {
		mapForm := (structs.New(evc)).Map()
		respInterface[i] = mapForm
	}

	return c.JSON(http.StatusOK, utils.BldGnrRsp(200, "Success", &respInterface))
}

func PostTask(c echo.Context) error {

	var task model.Task

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, utils.BldGnrRsp(http.StatusBadRequest, err.Error(), nil))
	}

	id := uuid.New().String()

	task.Id = id
	err := db.InsertTaskDetails(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BldGnrRsp(500, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, utils.BldGnrRsp(200, "Success", utils.ToInterfaceArrayFromString(id)))

}
