// Run this to test: go test github.com/archcra/pep/boardHelper
// after . ~/.bashrc (Set env variable: GOPATH

package boardHelper

import (
	"testing"
)

func TestEvaluate(t *testing.T) {
	cases := []struct {
		fen  string
		want int
	}{
		{"rheagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR", 0},
		{"rheagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C3C3/9/RHEAGAEHR", -10},
		{"rheagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C2C4/9/RHEAGAEHR", 28},
		{"rheagaehr/9/1c2c4/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR", -28},
		{"rheaSaehr/9/1c2c4/s1s1s1s1s/9/9/S1S3S1S/1C5C1/9/RHEAGAEHR", 125951},
		{"rheagaehr/9/1c2c4/s1s3s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAsAEHR", -125997},
	}

	for _, c := range cases {
		got := Evaluate(Fen2Board(c.fen))

		if got != c.want {
			t.Errorf("Mobility of board fen (%s) want %d, but got %d.", c.fen, c.want, got)
		}
	}
}
