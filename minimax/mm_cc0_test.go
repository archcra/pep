// Run this to test: go test github.com/archcra/pep/boardHelper
// after . ~/.bashrc (Set env variable: GOPATH

package minimax

import "testing"

func TestMm0Cc(t *testing.T) {
	cases := []struct {
		fen        string
		roundColor int
		depthLimit int
		want       string
	}{
		{"3g1S3/4S4/9/9/9/9/9/9/4s4/3s1G3", 1, 3, "1:6-1:5"},       //passed
		{"3g1S3/4S4/9/9/9/9/9/9/4s4/3s1G3", -1, 3, "10:4-10:5"},    // passed
		{"3a1a3/C1H1g4/4e4/9/8C/9/9/9/4s4/3s1G3", 1, 1, "2:1-2:5"}, //马后炮 passed
		{"3aga3/C8/4e4/3H5/8C/9/9/9/4s4/3s1G3", 1, 3, "4:4-2:3"},   //马后炮 passed
		//{"cC7/5g1h1/3Re1R2/9/2e6/9/9/9/1hr1s4/3G1c1r1", 1, 5, "3:4-2:4"}, //微信天天象棋第38关[楚]

	}

	for _, c := range cases {
		got := MinimaxCc(c.fen, c.depthLimit, c.roundColor)

		if got.Move != c.want {
			t.Errorf("First move of board fen (%s) want %s, but got %s.", c.fen, c.want, got.Move)
		}
	}
}
