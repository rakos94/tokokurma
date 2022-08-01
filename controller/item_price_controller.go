package controller

import (
	"net/http"
	"tokokurma/db"
	"tokokurma/db/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ItemPriceController struct{}

func NewItemPriceController() *ItemPriceController {
	return &ItemPriceController{}
}

func (o ItemPriceController) Get(c *gin.Context) {
	var res *[]model.ItemPrice
	err := db.ConnDB.Debug().Find(&res).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (o ItemPriceController) First(c *gin.Context) {
	var res *model.ItemPrice
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

func (o ItemPriceController) Add(c *gin.Context) {
	// req := &model.ItemPrice{
	// 	Name: c.PostForm("name"),
	// }
	// err := db.ConnDB.Debug().Create(&req).Error
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// c.JSON(http.StatusOK, req)

	var res *model.Item
	err := db.ConnDB.Debug().Where("id", c.Param("id")).First(&res).Error
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			statusCode = http.StatusNoContent
		}
		c.JSON(statusCode, err.Error())
		return
	}
}

func (o ItemPriceController) Delete(c *gin.Context) {
	var res *model.ItemPrice
	err := db.ConnDB.Debug().Where("id", c.Param("id")).Delete(&res).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "success")
}

func (o ItemPriceController) Update(c *gin.Context) {
	// var res *model.ItemPrice
	// q := db.ConnDB.Debug().Where("id", c.Param("id")).First(&res)
	// err := q.Error
	// if err != nil {
	// 	statusCode := http.StatusInternalServerError
	// 	if err == gorm.ErrRecordNotFound {
	// 		statusCode = http.StatusNoContent
	// 	}
	// 	c.JSON(statusCode, err.Error())
	// 	return
	// }

	// res.Name = c.PostForm("name")
	// err = q.Updates(&res).Error
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// c.JSON(http.StatusOK, res)
}
