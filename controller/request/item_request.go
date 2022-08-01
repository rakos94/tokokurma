package request

type ItemAddRequest struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}
