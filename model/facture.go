package model

import (
	"gorm.io/gorm/clause"
	"time"
)

// Facture table
type Facture struct {
	FactureID uint            `gorm:"primaryKey" json:"factureID"`
	UserID    uint            `json:"userid"`
	User      User            `gorm:"foreginKey:UserID" json:"user"`
	Date      time.Time       `gorm:"default:current_timestamp" json:"date"`
	Country   string          `gorm:"size:200" json:"country"`
	City      string          `gorm:"size:200" json:"city"`
	Detail    []FactureDetail `gorm:"foreginKey:FactureID" json:"detail"`
}


// SaveFacture method for saving facture
func (f *Facture) SaveFacture() error {
	return DB.Create(f).Error
}

// GetFactures method for getting all factures
func (Facture) GetFactures() ([]Facture,error) {
	var facture []Facture 
	err := DB.Model(&Facture{}).Preload(clause.Associations).Preload("Detail.Product").Find(&facture).Error
	return facture, err
}

// GetFacture method return facture by id
func (f *Facture) GetFacture(id uint) error {
	f.FactureID = id
	return DB.Preload(clause.Associations).Preload("Detail.Product").First(&f).Error
}

// GetFactureUser method return facture by user
func (f Facture) GetFactureUser(id uint) ([]Facture,error) {
	var facture []Facture
	err := DB.Where("user_id = ?",id).Preload(clause.Associations).Preload("Detail.Product").First(&facture).Error 
	return facture,err
}