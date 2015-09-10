// Run this to test: go test github.com/archcra/pep/boardHelper
// after . ~/.bashrc (Set env variable: GOPATH

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
        {"4g4/4a4/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR", 2,5,Pos{1,4}},
        {"4g4/4a4/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/4A4/4G4", 9,5,Pos{10,4}},
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


func TestElephantMoves(t *testing.T) {
	cases := []struct {
		fen string
        row int
        col int
        want Pos
	}{
		{"rheagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR", 1,3, Pos{3,5}},
        {"1heagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR", 1,7,Pos{3,5}},
        {"4g4/3a4/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C2E2C1/9/RH1AGAEHR", 8,5,Pos{6,3}},
        {"4g4/4a4/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C2E2C1/4A4/4G4", 8,5,Pos{10,7}},
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

func TestGeneralMoves(t *testing.T) {
	cases := []struct {
		fen string
        row int
        col int
        want Pos
	}{
		{"rheagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR", 1,5, Pos{2,5}},
        {"1he1gaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR", 1,5,Pos{1,4}},
        {"4g4/3a4/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C2E2C1/9/RH1AGAEHR", 10,5,Pos{9,5}},
        {"4g4/9/1c5c1/s1s3s1s/9/9/S1S3S1S/9/9/4G4", 10,5,Pos{1,5}},
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

func TestSoldierMoves(t *testing.T) {
	cases := []struct {
		fen string
        row int
        col int
        want Pos
	}{
		{"rheagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR", 4,5, Pos{5,5}},
        {"1he1gaehr/9/9/9/9/9/S1S1s1S1S/1C5C1/9/RHEAGAEHR", 7,5,Pos{7,6}},
        {"4g4/3a4/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C2E2C1/9/RH1AGAEHR", 7,5,Pos{6,5}},
        {"4g4/9/1c5c1/s1s1S1s1s/9/9/9/9/9/4G4", 4,5,Pos{4,6}},
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


func TestHorseMoves(t *testing.T) {
	cases := []struct {
		fen string
        row int
        col int
        want Pos
	}{
		{"rheagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR", 1,2, Pos{3,3}},
        {"1he1gaehr/9/9/9/9/9/S1S1s1S1S/1C5C1/9/RHEAGAEHR", 1,8,Pos{3,9}},
        {"4g4/3a4/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C2E2C1/9/RH1AGAEHR", 10,2,Pos{8,1}},
        {"4g4/9/1c2H2c1/s1s1S1s1s/9/9/9/9/9/4G4", 3,5,Pos{2,3}},
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

func TestRookMoves(t *testing.T) {
	cases := []struct {
		fen string
        row int
        col int
        want Pos
	}{
		{"rheagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR", 1,1, Pos{3,1}},
        {"1he1gaehr/9/9/9/9/9/S1S1s1S1S/1C5C1/9/RHEAGAEHR", 1,9,Pos{3,9}},
        {"4g4/3a4/1c5c1/s1s1R1s1s/9/9/S1S1S1S1S/1C2E2C1/9/RH1AGAEH1", 4,5,Pos{4,7}},
        {"4g4/9/1c2R2c1/s1s1S1s1s/9/9/9/9/9/4G4", 3,5,Pos{1,5}},
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


func TestCannonMoves(t *testing.T) {
	cases := []struct {
		fen string
        row int
        col int
        want Pos
	}{
		{"rheagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR", 3,2, Pos{10,2}},
        {"1he1gaehr/9/1c5c1/9/9/9/S1S1s1S1S/1C5C1/9/RHEAGAEHR", 3,8,Pos{3,5}},
        {"4g4/3a4/1c5c1/s1s1R1s1s/9/9/S1S1S1S1S/1C5C1/9/RH1AGAEH1", 8,8,Pos{8,5}},
        {"4g4/9/ec2C2c1/s1s1S1s1s/9/9/9/9/9/4G4", 3,5,Pos{3,1}},
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
