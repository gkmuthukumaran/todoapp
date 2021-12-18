package controllers

import (
	"net/http"
	"time"

	"github.com/taskpoc/operations"
	"github.com/taskpoc/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func Login(c echo.Context) error {

	username := c.FormValue("username")
	password := c.FormValue("password")

	if operations.IsValidUser(username, password) {
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = username
		claims["exp"] = time.Now().Add(time.Minute * 300).Unix()

		t, err := token.SignedString([]byte(viper.GetString("jwt.key")))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.BldGnrRsp(500, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, utils.BldGnrRsp(200, "logged in Successfully", utils.ToInterfaceArrayFromString(t)))
	}
	return c.JSON(http.StatusInternalServerError, utils.BldGnrRsp(500, "Invalid Username/Password", nil))
}
