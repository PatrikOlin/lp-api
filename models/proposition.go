package models

import (
	"time"

	"github.com/PatrikOlin/lp-api/db"
	"gorm.io/gorm"
)

type Proposition struct {
	gorm.Model
	PickupID   uint              `json:"pickupID"`
	UserID     uint              `json:"userID"`
	PickupTime time.Time         `json:"pickupTime"`
	Status     PropositionStatus `json:"status"`
}

func (p *Proposition) CreatePropositionRecord() error {
	res := db.GlobalDB.Create(&p)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
