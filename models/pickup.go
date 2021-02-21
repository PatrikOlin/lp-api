package models

import (
	"time"

	"github.com/PatrikOlin/lp-api/db"
	"gorm.io/gorm"
)

type Pickup struct {
	gorm.Model
	UserID          uint          `json:"userID"`
	Latitude        float64       `json:"latitude"`
	Longitude       float64       `json:"longitude"`
	Status          PickupStatus  `json:"status"`
	PickupTimeStart time.Time     `json:"pickupTimeStart"`
	PickupTimeEnd   time.Time     `json:"pickupTimeEnd"`
	Haul            Haul          `json:"haul" gorm:"foreignKey:PickupID"`
	Propositions    []Proposition `json:"propositions"`
}

func (p *Pickup) CreatePickupRecord() error {
	res := db.GlobalDB.Create(&p)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
