package controller

import (
	"errors"
	"net/http"
	"strconv"
	"tokokurma/db"
	"tokokurma/db/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderController struct{}

func NewOrderController() *OrderController {
	return &OrderController{}
}

func (o OrderController) Get(c *gin.Context) {
	var res []model.Order
	err := db.ConnDB.Debug().Scopes(db.Paginate(c.Request)).Preload("UserOrder").Preload("Customer").Find(&res).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	for i, v := range res {
		res[i].CustomerName = v.Customer.Name
		res[i].Customer = nil
	}
	c.JSON(http.StatusOK, res)
}

func (o OrderController) First(c *gin.Context) {
	var res *model.Order
	err := db.ConnDB.Debug().Where("id", c.Param("id")).Preload("CustomerOrder").Preload("UserOrder").First(&res).Error
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

func (o OrderController) Add(c *gin.Context) {
	CustomerID, _ := strconv.Atoi(c.PostForm("customer_id"))
	ItemID, _ := strconv.Atoi(c.PostForm("item_id"))
	Quantity, _ := strconv.Atoi(c.PostForm("quantity"))
	res := &model.Order{
		CustomerID: CustomerID,
		Quantity:   Quantity,
	}

	res, err := getItemData(res, ItemID, Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = db.ConnDB.Debug().
		Preload("CustomerOrder").
		Preload("UserOrder").Create(&res).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = updateCustomerOrder(res, CustomerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = updateUserOrder(res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (o OrderController) Delete(c *gin.Context) {
	var res *model.Order
	err := db.ConnDB.Debug().Where("id", c.Param("id")).Delete(&res).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "success")
}

func (o OrderController) Update(c *gin.Context) {
	var res *model.Order
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

	CustomerID, _ := strconv.Atoi(c.PostForm("customer_id"))
	ItemID, _ := strconv.Atoi(c.PostForm("item_id"))
	Quantity, _ := strconv.Atoi(c.PostForm("quantity"))
	res.CustomerID = CustomerID

	res, err = getItemData(res, ItemID, Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = q.Updates(&res).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = updateCustomerOrder(res, CustomerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = updateUserOrder(res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func getItemData(res *model.Order, ItemID int, Quantity int) (*model.Order, error) {
	var item *model.Item
	err := db.ConnDB.Debug().Where("id", ItemID).Preload("ItemPrice").First(&item).Error
	if err != nil {
		return nil, err
	}

	if Quantity > item.ItemPrice.Stock {
		return nil, errors.New("order melebihi stock")
	}

	res.ItemID = ItemID
	res.ItemName = item.Name
	price := 0
	if item.ItemPrice != nil {
		price = item.ItemPrice.Price
		item.ItemPrice.Stock -= Quantity
		err = db.ConnDB.Save(item.ItemPrice).Error
		if err != nil {
			return nil, err
		}
	}
	res.TotalPrice = price * Quantity

	return res, nil
}

func updateCustomerOrder(res *model.Order, CustomerID int) error {
	var customer *model.Customer
	err := db.ConnDB.Debug().Where("id", CustomerID).First(&customer).Error
	if err != nil {
		return err
	}

	res.CustomerOrder = &model.CustomerOrder{
		OrderID:      res.ID,
		CustomerName: customer.Name,
		CustomerHp:   customer.Hp,
	}

	err = db.ConnDB.Save(res.CustomerOrder).Error
	if err != nil {
		return err
	}
	return nil
}

func updateUserOrder(res *model.Order) error {
	res.UserOrder = &model.UserOrder{
		OrderID: res.ID,
		UserID:  model.AuthenticatedUser.ID,
	}
	err := db.ConnDB.Save(res.UserOrder).Error
	if err != nil {
		return err
	}
	return nil
}
