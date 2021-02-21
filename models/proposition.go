package models

import (
	"time"

	"gorm.io/gorm"
)

type Proposition struct {
	gorm.Model
	PickupID   uint              `json:"pickupID"`
	PickupTime time.Time         `json:"pickupTime"`
	Status     PropositionStatus `json:"status"`
}
