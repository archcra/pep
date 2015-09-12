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
		{"rheagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR", 1, 4, "a"},
	}

	for _, c := range cases {
		got := Minimax(boardHelper.Fen2Board(c.fen), c.depthLimit, c.roundColor)

		if got.Move != c.want {
			t.Errorf("First move of board fen (%s) want %s, but got %s.", c.fen, c.want, got.Move)
		}
	}
}
