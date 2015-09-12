// Run this to test: go test github.com/archcra/pep/boardHelper
// after . ~/.bashrc (Set env variable: GOPATH

package minimax

import (
	"testing"

	"github.com/archcra/pep/boardHelper"
)

func TestMm0(t *testing.T) {
	cases := []struct {
		fen        string
		roundColor int
		depthLimit int
		want       string
	}{
		{"rheagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR", 1, 2, "10:1-9:1"},
		{"rheagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR", -1, 2, "1:1-2:1"},
		{"4g4/4S4/9/9/9/9/9/9/9/4G4", 1, 2, "2:5-1:5"},
		{"4g4/4S4/9/9/9/9/9/9/9/4G4", -1, 2, "1:5-1:6"},
		{"4g4/4S4/9/9/9/9/9/9/4A4/4G4", -1, 2, "1:5-2:5"},
		{"cC7/5g1h1/3Re1R2/9/2e6/9/9/9/1hr1s4/3G1c1r1", 1, 4, "1:5-2:5"}, //微信天天象棋第38关[楚]

	}

	for _, c := range cases {
		got := Minimax(boardHelper.Fen2Board(c.fen), c.depthLimit, c.roundColor)

		if got.Move != c.want {
			t.Errorf("First move of board fen (%s) want %s, but got %s.", c.fen, c.want, got.Move)
		}
	}
}
