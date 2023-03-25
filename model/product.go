package model

// Product table
type Product struct {
	ProductID   uint    `gorm:"primaryKey" json:"productid"`
	Name        string  `gorm:"size:200" json:"name"`
	Description string  `gorm:"size:500" json:"description"`
	Price       float32 `json:"price"`
	Stock       float32 `json:"stock"`
	Depreced    bool    `json:"depreced"`
}

// GetProducts method for get all products
func (p *Product) GetProducts() ([]Product, error) {
	var products []Product
	err := DB.Find(&products).Error
	return products, err
}

// SaveProducts method for saving products
func (p *Product) SaveProducts() error {
	return DB.Create(p).Error
}
