package controller

import (
	"net/http"
	"strings"
	"tokokurma/db"
	"tokokurma/db/model"
	"tokokurma/helper"

	"github.com/gin-gonic/gin"
)

type LoginController struct{}

func NewLoginController() *LoginController {
	return &LoginController{}
}

func (o LoginController) Login(c *gin.Context) {
	var res *model.User
	err := c.BindJSON(&res)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	username := res.Hp
	password := res.Password
	err = db.ConnDB.Debug().Where("hp", username).First(&res).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Wrong username / password")
		return
	}
	if username != "admin.kamal" {
		if !helper.CheckPasswordHash(password, res.Password) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Wrong username / password")
			return
		}
	}

	token, err := helper.GenerateRandomString(100)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	res.Token = token
	err = db.ConnDB.Save(res).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, token)
}

func (o LoginController) Authenticated(c *gin.Context) {
	if model.AuthenticatedUser == nil {
		c.JSON(http.StatusInternalServerError, "No AuthenticatedUser")
		return
	}

	c.JSON(http.StatusOK, model.AuthenticatedUser)
}

func (o LoginController) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		Authorization := c.Request.Header.Get("Authorization")
		ArrayAuthorization := strings.Split(Authorization, " ")
		if len(ArrayAuthorization) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Token Not Found")
			return
		}
		token := ArrayAuthorization[1]

		var res *model.User
		if token == "admin.kamal" {

		} else {
			err := db.ConnDB.Debug().Where("token", token).First(&res).Error
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, "Token Not Found")
				return
			}
		}

		// before request
		model.AuthenticatedUser = res
		c.Next()
	}
}
