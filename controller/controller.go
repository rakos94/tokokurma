package controller

var (
	Login     *LoginController
	Customer  *CustomerController
	Item      *ItemController
	ItemPrice *ItemPriceController
	Order     *OrderController
	User      *UserController
)

func init() {
	Login = NewLoginController()
	Customer = NewCustomerController()
	Item = NewItemController()
	ItemPrice = NewItemPriceController()
	Order = NewOrderController()
	User = NewUserController()
}
