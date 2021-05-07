package entity

type Inventory struct {
	ID        string `json:"id" gorm:"id"`
	ProductID string `json:"product_id" gorm:"product_id"`
	Quantity  string `json:"quantity" gorm:"quantity"`
}
