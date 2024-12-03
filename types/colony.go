package types

import (
	"fmt"
	"lem-in/xerrors"
)

const (
	MaxAntsPerColony = 1_000
)

// Colony is a model of the ant farm also known as colony
type Colony struct {
	StartRoom    string
	EndRoom      string
	NumberOfAnts uint32 // number of ants cannot be negative
}

// CheckNumAnts checks the number of ants per colony if they have exceeded MaxAntsPerColony
// and exits with status code 1
func CheckNumAnts(c Colony) error {
	if c.NumberOfAnts >= MaxAntsPerColony {
		return fmt.Errorf(xerrors.ErrMaxAntNumExceeded.Error(), MaxAntsPerColony)
	}
	return nil
}
