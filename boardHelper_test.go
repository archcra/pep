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


func TestBoard2Fen(t *testing.T) {
	got := InitFullBoard()
    wantFen := "rheagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR"
    fen := board2Fen(got);
    if fen != wantFen {
        t.Errorf("Fen of full board should be: %s but is: %s.", wantFen, fen)
    }
}

func TestValidFen(t *testing.T) {
	cases := []struct {
		in string 
        want bool
	}{
		{"rheagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR", true},
        {"rheagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR/", false},
		{"columns", false},
        {"r8/9/9/9/9/9/9/9/9/H8", true},
		{"r8/9/9/9/9/9/9/9/9/H8/", false},
        {"r8/9/9/9/9/9/9/9/z/H8", false},
	}

	for _, c := range cases {
        got := validFen(c.in)
				if got != c.want {
				t.Errorf("Fen string (%s) , should be %b", c.in, c.want)
			}
		}
    
    
}


// With composite type literal
// Ref: http://stackoverflow.com/questions/17912893/missing-type-in-composite-literal
func TestFen2Board(t *testing.T) {
      type (  b struct {
                pos string
                value  int
        }
        )
    

	cases := []struct {
		in string 
        want map[string]int
	}{
		{"rheagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR", map[string]int{"1,5": 254}},
		{"columns", map[string]int{"1,5": -1}},
        {"r8/9/9/9/9/9/9/9/9/H8", map[string]int{"1,5": 0}},
	}

	for _, c := range cases {
        got := Fen2Board(c.in)
				
        for key, value := range c.want {
        pos := strings.Split(key, ",")
            // t.Errorf("pos are: %s, %s", pos[0], pos[1])
            row, _ := strconv.Atoi(pos[0])
            col, _ := strconv.Atoi(pos[1])
			if got[row][col] != value {
				t.Errorf("Fen string (%s) got board position %s should be %d but is %d!", c.in, key, value, got[row][col])
            }
		}
    }

// Net cases set
cases2 := []struct {
		in string 
        want b
	}{
		{"rheagaehr/9/1c5c1/s1s1s1s1s/9/9/S1S1S1S1S/1C5C1/9/RHEAGAEHR", b{"10,5", 126}},
		{"columns", b{"10,5",  -1}},
        {"r8/8s/9/9/9/9/9/9/9/H8", b{"2,9",  143}},
	}

	for _, c := range cases2 {
        got := Fen2Board(c.in)
				pos := strings.Split(c.want.pos, ",")
            // t.Errorf("pos are: %s, %s", pos[0], pos[1])
            row, _ := strconv.Atoi(pos[0])
            col, _ := strconv.Atoi(pos[1])
			if got[row][col] != c.want.value {
				t.Errorf("Fen string (%s) got board position %s should be %d but is %d!", c.in, c.want.pos, c.want.value, got[row][col])
			
		}
    }
}
