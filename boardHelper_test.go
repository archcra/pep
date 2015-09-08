// First part, test initBoard func

package pep

import (
	"testing"
    "strings"
    "strconv"
)

func TestInitBoard(t *testing.T) {

	cases := []struct {
		in string 
        want int
	}{
		{"rows", 13},
		{"columns", 12},
		{"1,5", 254},
        {"10,5", 126},
        {"11,1", -1},
	}

	got := InitFullBoard()

	for _, c := range cases {
		switch c.in {
		case "rows":
			if len(got) != c.want {
				t.Errorf("Full board rows (%q) , want %d", got, c.want)
			}
		case "columns":
			if len(got[0]) != c.want {
				t.Errorf("Full board columns %q, want %d",  got, c.want)
			}
		default:
            pos := strings.Split(c.in, ",")
            // t.Errorf("pos are: %s, %s", pos[0], pos[1])
            row, _ := strconv.Atoi(pos[0])
            col, _ := strconv.Atoi(pos[1])
			if got[row][col] != c.want {
				t.Errorf("Position (%s) of full board == %q, want %d", c.in, got, c.want)
			}
		}
    }

	// source code: http://golang.org/pkg/testing/#T.Logf

	/*
	Board is [[-1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1] [-1 208 158 140 139 254 139 140 158 208 -1 -1] [-1 0 0 0 0 0 0 0 0 0 -1 -1] [-1 0 163 0 0 0 0 0 163 0 -1 -1] [-1 143 0 143 0 143 0 143 0 143 -1 -1] [-1 0 0 0 0 0 0 0 0 0 -1 -1] [-1 0 0 0 0 0 0 0 0 0 -1 -1] [-1 15 0 15 0 15 0 15 0 15 -1 -1] [-1 0 35 0 0 0 0 0 35 0 -1 -1] [-1 0 0 0 0 0 0 0 0 0 -1 -1] [-1 80 30 12 11 126 11 12 30 80 -1 -1] [-1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1] [-1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1]]
	*/

	// 检查总行

	// Check total columns
	// Check piece at 1,5 is black general
	// Check piece at 10,5 is red general
	// Check there is no piece at 11,1

}

func BenchmarkInitBoard(b *testing.B) {
	board := InitFullBoard()
	b.Logf("Board is %v", board)
}
