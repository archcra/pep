package main

import (
	"flag"
	"fmt"

	"log"
	"net/http"

	"github.com/archcra/pep/minimax"
)
import _ "net/http/pprof"

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {

	fen := "cC7/5g1h1/3Re1R2/9/2e6/9/9/9/1hr1s4/3G1c1r1" //微信天天象棋第38关[楚]
	roundColor := 1
	depthLimit := 5

	got := minimax.MinimaxCc(fen, depthLimit, roundColor)

	fmt.Printf("\nFirst move of board fen (%s) got %s.\n", fen, got.Move)

	//profile
	log.Println(http.ListenAndServe("localhost:6060", nil))
	//end profile
}
