package model

import (
	"time"
)

// Facture table
type Facture struct {
	FactureID uint            `gorm:"primaryKey" json:"factureID"`
	UserID    uint            `json:"userid"`
	User      User            `gorm:"foreginKey:UserID" json:"user"`
	Date      time.Time       `json:"date"`
	Country   string          `gorm:"size:200" json:"country"`
	City      string          `gorm:"size:200" json:"city"`
	Detail    []FactureDetail `gorm:"foreginKey:FactureDetailID" json:"detail"`
}
