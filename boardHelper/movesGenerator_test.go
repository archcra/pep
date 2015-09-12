// Run this to test: go test github.com/archcra/pep/boardHelper
// after . ~/.bashrc (Set env variable: GOPATH

package boardHelper

import (
	"testing"
)

func TestGenerateMoves(t *testing.T) {
	cases := []struct {
		fen        string
		roundColor int // 1 RED; -1 BLACK
		want       string
	}{
		{"rheagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR", 1, "8:2-8:5"},
		{"1heagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR", -1, "3:8-3:5"},
		{"4g4/4a4/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/4C4/9/RHEAGAEHR", 1, "8:5-4:5"},
		{"rheagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR", -1, "3:2-10:2"},
	}

	for _, c := range cases {
		got := Generate(Fen2Board(c.fen), c.roundColor)

		found := false
		for _, v := range got {
			if v.Move == c.want {
				found = true
			}
		}
		if !found {
			t.Errorf("Moves of board fen (%s) on round color %d should have are %s, but got %s", c.fen, c.roundColor, c.want, extractMoveStr(got))
		}
	}
}

func extractMoveStr(moveResult []MoveResult) string {
	var result string
	for _, v := range moveResult {
		result = result + v.Move + ";"
	}
	return result
}
