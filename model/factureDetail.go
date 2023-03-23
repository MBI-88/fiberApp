package model

// FactureDetail table
type FactureDetail struct {
	FactureDetailID uint    `gorm:"primaryKey" json:"fatureDetailID"`
	FactureID       uint    `json:"factureID"`
	Facture         Facture `gorm:"foreginKey:FactureID" json:"facture"`
	Size            float32 `json:"size"`
	Price           float32 `json:"price"`
	Discount        float32 `json:"discount"`
}
