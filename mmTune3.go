package main

// Follows here: http://blog.golang.org/profiling-go-programs
// Now concern the Memory
import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/archcra/pep/minimax"
)

var memprofile = flag.String("memprofile", "", "write memory profile to this file")

func main() {
	flag.Parse()

	fen := "cC7/5g1h1/3Re1R2/9/2e6/9/9/9/1hr1s4/3G1c1r1" //微信天天象棋第38关[楚]
	roundColor := 1
	depthLimit := 5

	got := minimax.MinimaxCc(fen, depthLimit, roundColor)

	fmt.Printf("\nFirst move of board fen (%s) got %s.\n", fen, got.Move)

	//profile memory
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}
	//end profile

}
