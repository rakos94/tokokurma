package controller

import (
	"net/http"
	"tokokurma/db"
	"tokokurma/db/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CustomerController struct{}

func NewCustomerController() *CustomerController {
	return &CustomerController{}
}

func (o CustomerController) Get(c *gin.Context) {
	var res *[]model.Customer
	err := db.ConnDB.Debug().Scopes(db.Paginate(c.Request)).Find(&res).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (o CustomerController) First(c *gin.Context) {
	var res *model.Customer
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

func (o CustomerController) Add(c *gin.Context) {
	req := &model.Customer{
		Name: c.PostForm("name"),
		Hp:   c.PostForm("hp"),
	}
	err := db.ConnDB.Debug().Create(&req).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, req)
}

func (o CustomerController) Delete(c *gin.Context) {
	var res *model.Customer
	err := db.ConnDB.Debug().Where("id", c.Param("id")).Delete(&res).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "success")
}

func (o CustomerController) Update(c *gin.Context) {
	var res *model.Customer
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

	res.Name = c.PostForm("name")
	res.Hp = c.PostForm("hp")
	err = q.Updates(&res).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}
