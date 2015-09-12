// Run this to test: go test github.com/archcra/pep/boardHelper
// after . ~/.bashrc (Set env variable: GOPATH

package boardHelper

import (
	"testing"
)

func TestGeneralPos(t *testing.T) {
	cases := []struct {
		fen   string
		color int
		want  Pos
	}{
		{"rheagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR", 1, Pos{10, 5}},
		{"rheagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C3C3/9/RHEAGAEHR", -1, Pos{1, 5}},
		{"rheagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C2C4/9/3G5", 1, Pos{10, 4}},
		{"5g3/9/1c2c4/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR", -1, Pos{1, 6}},
		{"9/9/1c2c4/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR", -1, Pos{0, 0}},
	}

	for _, c := range cases {
		got := GetGeneralPos(Fen2Board(c.fen), c.color)

		if got.row != c.want.row || got.col != c.want.col {
			t.Errorf("General Position of board fen (%s) want {%d:%d}, but got {%d:%d}.",
				c.fen, c.want.row, c.want.col, got.row, got.col)
		}
	}
}
