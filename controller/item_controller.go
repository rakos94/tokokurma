package controller

import (
	"net/http"
	"strconv"
	"tokokurma/controller/request"
	"tokokurma/db"
	"tokokurma/db/model"
	"tokokurma/helper"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ItemController struct{}

func NewItemController() *ItemController {
	return &ItemController{}
}

func (o ItemController) Get(c *gin.Context) {
	var res *[]model.Item
	err := db.ConnDB.Debug().Scopes(db.Paginate(c.Request)).Preload("ItemPrice").Find(&res).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (o ItemController) First(c *gin.Context) {
	var res *model.Item
	err := db.ConnDB.Debug().Preload("ItemPrice").Where("id", c.Param("id")).First(&res).Error
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

func (o ItemController) Add(c *gin.Context) {
	reqData := request.ItemAddRequest{}
	if err := c.ShouldBindJSON(&reqData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req := &model.Item{
		Name: reqData.Name,
		ItemPrice: &model.ItemPrice{
			Price: reqData.Price,
			Stock: reqData.Stock,
		},
	}
	if req.Name == "" {
		c.JSON(http.StatusBadRequest, "Empty Name")
		return
	}
	err := db.ConnDB.Debug().Create(&req).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, req)
}

func (o ItemController) Delete(c *gin.Context) {
	var res *model.Item
	err := db.ConnDB.Debug().Where("id", c.Param("id")).Delete(&res).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "success")
}

func (o ItemController) Update(c *gin.Context) {
	Price, _ := strconv.Atoi(c.PostForm("price"))
	Stock, _ := strconv.Atoi(c.PostForm("stock"))
	var res *model.Item
	q := db.ConnDB.Debug().Where("id", c.Param("id")).Preload("ItemPrice").First(&res)
	err := q.Error
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			statusCode = http.StatusNoContent
		}
		c.JSON(statusCode, err.Error())
		return
	}

	var oldData *model.Item
	helper.DeepCopy(res, &oldData)

	res.Name = c.PostForm("name")
	res.ItemPrice.Price = Price
	res.ItemPrice.Stock = Stock

	err = db.ConnDB.Save(res.ItemPrice).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	res.CreateLog(oldData)

	err = q.Updates(&res).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (o ItemController) History(c *gin.Context) {
	var res *model.Item
	err := db.ConnDB.Debug().Where("id", c.Param("id")).Preload("Logs", func(_db *gorm.DB) *gorm.DB {
		return _db.Scopes(db.Paginate(c.Request))
	}).First(&res).Error
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
