// First part, test initBoard func

package boardHelper

import (
	"testing"
)


func TestAdvisorMoves(t *testing.T) {
	cases := []struct {
		fen string
        row int
        col int
        want Pos
	}{
		{"rheagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR", 1,4, Pos{2,5}},
        {"1heagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR", 1,6,Pos{2,5}},
	}

	for _, c := range cases {
        got := getPossibleMovesInfo(Fen2Board(c.fen), c.row,c.col)
        
        found := false
        for _, v := range got.moves {
            if v.row == c.want.row && v.col == c.want.col {
            found = true
            }
        }
        if !found {
             t.Errorf("Moves on pos %d:%d of board fen (%s) are %q, but got %q.", c.row, c.col, c.fen, c.want, got.moves)   
        }
    }
}