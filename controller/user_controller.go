package controller

import (
	"net/http"
	"tokokurma/db"
	"tokokurma/db/model"
	"tokokurma/helper"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (o UserController) Get(c *gin.Context) {
	var res *[]model.User
	err := db.ConnDB.Debug().Scopes(db.Paginate(c.Request)).Find(&res).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (o UserController) First(c *gin.Context) {
	var res *model.User
	err := db.ConnDB.Debug().Where("id", c.Param("id")).First(&res).Error
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			statusCode = http.StatusNoContent
		}
		c.JSON(statusCode, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (o UserController) Add(c *gin.Context) {
	pwd, err := helper.HashPassword(c.PostForm("password"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	req := &model.User{
		Username: c.PostForm("username"),
		Name:     c.PostForm("name"),
		Hp:       c.PostForm("hp"),
		Password: pwd,
	}
	err = db.ConnDB.Debug().Create(&req).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, req)
}

func (o UserController) Delete(c *gin.Context) {
	var res *model.User
	err := db.ConnDB.Debug().Where("id", c.Param("id")).Delete(&res).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "success")
}

func (o UserController) Update(c *gin.Context) {
	var res *model.User
	q := db.ConnDB.Debug().Where("id", c.Param("id")).First(&res)
	err := q.Error
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			statusCode = http.StatusNoContent
		}
		c.JSON(statusCode, err.Error())
		return
	}

	res.Username = c.PostForm("username")
	res.Name = c.PostForm("name")
	res.Hp = c.PostForm("hp")

	pwd, err := helper.HashPassword(c.PostForm("password"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	res.Password = pwd
	err = q.Updates(&res).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}
