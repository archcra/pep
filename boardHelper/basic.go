package boardHelper

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	BOARD_COLUMNS = 9
	BOARD_ROWS    = 10

	ADVISOR   = 11
	ELEPHANT  = 12
	SOLDIER   = 15
	HORSE     = 30
	CANNON    = 35
	ROOK      = 80
	GENERAL   = 126
	RIVER     = 5
	RP_MARGIN = 2 //RP_MARGIN: River to Palace's margin
	COLOR     = 128
	LIMBO     = 128

	RED   = 1
	BLACK = -1
)

var pieceValue = map[string]int{
	"r": LIMBO + ROOK,
	"h": LIMBO + HORSE,
	"e": LIMBO + ELEPHANT,
	"a": LIMBO + ADVISOR,
	"g": LIMBO + GENERAL,
	"c": LIMBO + CANNON,
	"s": LIMBO + SOLDIER,
	"R": ROOK,
	"H": HORSE,
	"E": ELEPHANT,
	"A": ADVISOR,
	"G": GENERAL,
	"C": CANNON,
	"S": SOLDIER,
}

var pieceStr = map[int]string{
	ROOK + LIMBO:     "r",
	HORSE + LIMBO:    "h",
	ELEPHANT + LIMBO: "e",
	ADVISOR + LIMBO:  "a",
	GENERAL + LIMBO:  "g",
	CANNON + LIMBO:   "c",
	SOLDIER + LIMBO:  "s",
	ROOK:             "R",
	HORSE:            "H",
	ELEPHANT:         "E",
	ADVISOR:          "A",
	GENERAL:          "G",
	CANNON:           "C",
	SOLDIER:          "S",
}

var pieceScore = map[string]int{
	"r": -ROOK,
	"h": -HORSE,
	"e": -ELEPHANT,
	"a": -ADVISOR,
	"g": -GENERAL * 1000,
	"c": -CANNON,
	"s": -SOLDIER,
	"R": ROOK,
	"H": HORSE,
	"E": ELEPHANT,
	"A": ADVISOR,
	"G": GENERAL * 1000,
	"C": CANNON,
	"S": SOLDIER,
}

func InitFullBoard() [13][12]int { // https://groups.google.com/forum/#!topic/golang-nuts/sPYRl4RHFdU
	board := [13][12]int{{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, 208, 158, 140, 139, 254, 139, 140, 158, 208, -1, -1},
		{-1, 0, 0, 0, 0, 0, 0, 0, 0, 0, -1, -1},
		{-1, 0, 163, 0, 0, 0, 0, 0, 163, 0, -1, -1},
		{-1, 143, 0, 143, 0, 143, 0, 143, 0, 143, -1, -1},
		{-1, 0, 0, 0, 0, 0, 0, 0, 0, 0, -1, -1},
		{-1, 0, 0, 0, 0, 0, 0, 0, 0, 0, -1, -1},
		{-1, 15, 0, 15, 0, 15, 0, 15, 0, 15, -1, -1},
		{-1, 0, 35, 0, 0, 0, 0, 0, 35, 0, -1, -1},
		{-1, 0, 0, 0, 0, 0, 0, 0, 0, 0, -1, -1},
		{-1, ROOK, HORSE, 12, 11, 126, 11, 12, HORSE, ROOK, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}}
	return board
}

// board object to FEN string
// returns false if the obj is not a valid position object
func Board2Fen(board [13][12]int) string {
	var fen string
	// currentRow := 8

	for row := 1; row <= 10; row++ {
		for column := 1; column <= 9; column++ {
			piece := board[row][column]

			// piece exists
			if piece != 0 {
				fen += pieceCodeToFen(piece)
			} else { // empty space
				fen += "1"
			}
		}
		fen += "/" // Next row
	}

	// squeeze the numbers together
	// haha, I love this solution...
	r := regexp.MustCompile("111111111")
	fen = r.ReplaceAllString(fen, "9")
	r = regexp.MustCompile("11111111")
	fen = r.ReplaceAllString(fen, "8")
	r = regexp.MustCompile("1111111")
	fen = r.ReplaceAllString(fen, "7")
	r = regexp.MustCompile("111111")
	fen = r.ReplaceAllString(fen, "6")
	r = regexp.MustCompile("11111")
	fen = r.ReplaceAllString(fen, "5")
	r = regexp.MustCompile("1111")
	fen = r.ReplaceAllString(fen, "4")
	r = regexp.MustCompile("111")
	fen = r.ReplaceAllString(fen, "3")
	r = regexp.MustCompile("11")
	fen = r.ReplaceAllString(fen, "2")

	return fen[:len(fen)-1]
}

// convert code to fen string
func pieceCodeToFen(piece int) string {
	for key, value := range pieceValue {
		if value == piece {
			return key
		}
	}
	return "?" // This should not happen.
}

// TODO: this whole function could probably be replaced with a single regex
func validFen(fen string) bool {

	// cut off any move, castling, etc info from the end
	// we're only interested in position information
	r := regexp.MustCompile(" .+$")
	fen = r.ReplaceAllString(fen, "")
	// FEN should be 8 sections separated by slashes
	chunks := strings.Split(fen, "/")
	// fmt.Println("Now chunks is:",len(chunks))

	if len(chunks) != BOARD_ROWS {
		return false
	}

	// check the piece sections
	for i := 0; i < BOARD_ROWS; i++ {
		if chunks[i] == "" ||
			len(chunks[i]) > BOARD_COLUMNS {
			return false
		}
		match, _ := regexp.MatchString("[^rheagcsRHEAGCS1-9]", chunks[i])
		if match {
			return false
		}
	}
	return true
}

// convert FEN string to Board object
// returns nil if the FEN string is invalid
func Fen2Board(fen string) [13][12]int {
	var board [13][12]int

	// Init to all -1, magic number TODO
	for i := 0; i < 13; i++ {
		for j := 0; j < 12; j++ {
			board[i][j] = -1
		}
	}

	if !validFen(fen) {
		return board
	}

	rows := strings.Split(fen, "/")

	for i := 0; i < BOARD_ROWS; i++ {
		row := strings.Split(rows[i], "")

		// loop through each character in the FEN section
		//emptySquares := 0
		columnIndex := 0
		for j := 0; j < len(row); j++ {
			// number / empty squares
			match, _ := regexp.MatchString("[1-9]", row[j])
			if match {

				emptySquares, _ := strconv.Atoi(row[j])
				for k := 0; k < emptySquares; k++ {
					board[i+1][columnIndex+1] = 0
					columnIndex = columnIndex + 1
				}
			} else { // piece
				board[i+1][columnIndex+1] = pieceValue[row[j]]
				columnIndex = columnIndex + 1
			}
		}
	}

	return board
}

func BoardPiecesScore(board [13][12]int) int {
	score := 0
	// Function same as piecesScore, this is for efficiency consideration
	for row := 1; row <= BOARD_ROWS; row++ {
		for column := 1; column <= BOARD_COLUMNS; column++ {
			piece := board[row][column]
			if piece > 0 {
				score = score + pieceScore[pieceStr[piece]]
			}
		}
	}
	return score
}
