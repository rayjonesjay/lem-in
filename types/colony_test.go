package types

import (
	"errors"
	"lem-in/xerrors"
	"math"
	"testing"
)

func TestCheckNumAnts(t *testing.T) {
	type testData struct {
		name  string
		input uint32
		want  error
	}
	tests := []testData{
		{"testA", 0, xerrors.ErrZeroAnts},
		{"testB", 1, nil},
		{"testC", 1_000, xerrors.ErrMaxAntNumExceeded},
		{"testD", 1_000_000, xerrors.ErrMaxAntNumExceeded},
		{"testE", math.MaxUint32, xerrors.ErrMaxAntNumExceeded},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			C := Colony{NumberOfAnts: test.input}
			got := CheckNumAnts(C)
			if errors.Is(got, test.want) && got != nil {
				t.Errorf("\nGOT %v\nWANT %v", got, test.want)
			}
		})
	}
}
