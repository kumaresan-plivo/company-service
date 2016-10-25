package models

import (
	"errors"
	"log"
	"time"

	"../db"
	"github.com/jinzhu/gorm"
)

type Company struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      bool   `json:"active,omitempty"`
}


func (h Company) GetByID(id string) (*Company, error) {
	db := db.GetDB()
	company := &Company{}
	return db.First(&company, id)
}
