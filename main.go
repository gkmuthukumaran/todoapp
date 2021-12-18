package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
	"github.com/taskpoc/controllers"
	"github.com/taskpoc/interfaces/db"
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

	db.SetupDB("DB")
	defer db.Closedb()
	// Echo instance
	app := echo.New()

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	//Add cors headers
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))
	// Route => handler
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Task API!\n")
	})
	taskGroup := app.Group("/v1/api")
	taskGroup.Use(middleware.JWT([]byte(viper.GetString("jwt.key"))))
	app.POST("/login", controllers.Login)
	taskGroup.GET("/task", controllers.GetAllTask)
	taskGroup.GET("/task/:id", controllers.GetTaskById)
	taskGroup.GET("/task/:category", controllers.GetTasksByCategory)
	taskGroup.DELETE("/task/:id", controllers.DeleteTask)
	taskGroup.POST("/task", controllers.PostTask)
	// Start server
	app.Logger.Fatal((app.Start((":" + port))))
}
