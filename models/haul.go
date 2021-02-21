package models

import "gorm.io/gorm"

type Haul struct {
	gorm.Model
	PickupID  uint        `json:"pickupID"`
	ImageURLs []HaulImage `json:"imageURLs"`
	Value     float64     `json:"value"`
	BagCount  uint8       `json:"bagCount"`
}

type HaulImage struct {
	gorm.Model
	HaulID uint   `json:"haulID"`
	URL    string `json:"URL"`
}
