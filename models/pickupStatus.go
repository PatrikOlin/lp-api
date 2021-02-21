package models

import "errors"

type PickupStatus string

const (
	Resolved   PickupStatus = "Resolved"
	Unresolved              = "Unresolved"
	Claimed                 = "Claimed"
	Canceled                = "Canceled"
)

// metod för att kolla så att status är en av dom godkända
// stringsen, används på föjande vis:
// if err := p.Status.IsValid(); err != nil {
// return nil, err }
func (s PickupStatus) IsValid() error {
	switch s {
	case Resolved, Unresolved, Claimed, Canceled:
		return nil
	}
	return errors.New("Invalid status")
}
