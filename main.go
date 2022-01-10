package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
	"github.com/todoapp/controllers"
)

func main() {

	fmt.Println("TodoPOC")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	port := viper.GetString("app.port")

	app := echo.New()

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	//Add cors headers
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	// Route => handler
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Task API!\n")
	})
	taskGroup := app.Group("/v1/api")
	taskGroup.POST("/task", controllers.InsertTaskDetails)
	taskGroup.GET("/task", controllers.GetAllTask)
	taskGroup.GET("/category", controllers.GetAllCategory)
	taskGroup.GET("/taskbyid/:id", controllers.GetTaskById)
	taskGroup.PUT("/taskupdate/:id", controllers.UpdateTaskById)
	taskGroup.DELETE("/taskdelete/:id", controllers.DeleteTask)
	taskGroup.GET("/taskbycategory/:category", controllers.GetTasksByCategory)

	// Start server
	app.Logger.Fatal((app.Start((":" + port))))
}
