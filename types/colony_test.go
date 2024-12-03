package types

import (
	"lem-in/xerrors"
	"testing"
)

func TestCheckNumAnts(t *testing.T) {
	type testData struct {
		input int
		want  error
	}
	tests := []testData{
		{0, xerrors.ErrZeroAnts},
		{1, nil},
		{1_000, nil},
		{1_000_000, xerrors.ErrMaxAntNumExceeded},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := CheckNumAnts(c.)
		})
	}
}
