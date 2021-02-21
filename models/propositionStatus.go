package models

import (
	"errors"
)

type PropositionStatus string

const (
	Accepted PropositionStatus = "Accepted"
	Pending                    = "Pending"
	Rejected                   = "Rejected"
)

func (ps PropositionStatus) IsValid() error {
	switch ps {
	case Accepted, Pending, Rejected:
		return nil
	}

	return errors.New("Invalid status")
}
