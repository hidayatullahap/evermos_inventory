package entity

type Inventory struct {
	ID        int64 `json:"id" gorm:"id"`
	ProductID int64 `json:"product_id" gorm:"product_id"`
	Quantity  int64 `json:"quantity" gorm:"quantity"`
}

func (t Inventory) TableName() string {
	return "inventories"
}
