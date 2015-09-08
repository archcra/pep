// First part, test initBoard func


package pep


import (
    "testing"
)


func TestInitBoard(t *testing.T) {
    /*
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}*/
    
    /*
	for _, c := range cases {
		got := InitFullBoard()
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
*/
    
    board :=  InitFullBoard();
    t.Logf("Board is %v", board)
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
     board :=  InitFullBoard();
    b.Logf("Board is %v", board)
}
