package boardHelper

import (
	"regexp"
    "strings"
    "strconv"
    //"fmt"
)

const (
	ROW_WIDTH  = 9
	ROW_HEIGHT = 10

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
)

var pieceValue = map[string]int{
	"r": COLOR + ROOK,
	"h": COLOR + HORSE,
	"e": COLOR + ELEPHANT,
	"a": COLOR + ADVISOR,
	"g": COLOR + GENERAL,
	"c": COLOR + CANNON,
	"s": COLOR + SOLDIER,
	"R": ROOK,
	"H": HORSE,
	"E": ELEPHANT,
	"A": ADVISOR,
	"G": GENERAL,
	"C": CANNON,
	"S": SOLDIER,
}

var pieceStr = map[string]string{
	"208": "r",
	"158": "h",
	"140": "e",
	"139": "a",
	"254": "g",
	"163": "c",
	"143": "s",
	"80":  "R",
	"30":  "H",
	"12":  "E",
	"11":  "A",
	"126": "G",
	"35":  "C",
	"15":  "S",
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
		{-1, 80, 30, 12, 11, 126, 11, 12, 30, 80, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}}
	return board
}

// board object to FEN string
// returns false if the obj is not a valid position object
func board2Fen(board [13][12]int) string {
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
func validFen(fen string ) bool {

    // cut off any move, castling, etc info from the end
    // we're only interested in position information
    r := regexp.MustCompile(" .+$")
	fen = r.ReplaceAllString(fen, "")
    // FEN should be 8 sections separated by slashes
    chunks := strings.Split(fen, "/")
    // fmt.Println("Now chunks is:",len(chunks))

    
    if len(chunks) != ROW_HEIGHT{
        return false
    }

    // check the piece sections
    for i := 0; i < ROW_HEIGHT; i++ {
        if chunks[i] == "" ||
            len(chunks[i]) > ROW_WIDTH {
            return false
        }
        match, _ := regexp.MatchString("[^rheagcsRHEAGCS1-9]", chunks[i])
        if match{
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
    for i:=0; i< 13; i++ {
        for j:=0; j< 12; j++ {
            board[i][j]=-1
        }
    }
    
    if !validFen(fen) {
        return board
    }

    rows := strings.Split(fen, "/")

   // var position = {};

    //var currentRow = ROW_HEIGHT;
    for  i := 0; i < ROW_HEIGHT; i++ {
        row := strings.Split(rows[i],"")

        // loop through each character in the FEN section
        //emptySquares := 0
            columnIndex := 0
        for j := 0; j < len(row); j++ {
            // number / empty squares
            match, _ := regexp.MatchString("[1-9]", row[j])
        if match{
            
            emptySquares, _  := strconv.Atoi(row[j])
            for k := 0; k < emptySquares; k++ {
                    board[i + 1][columnIndex + 1] = 0
                    columnIndex = columnIndex + 1
                }
        }else {// piece
                board[i + 1][columnIndex + 1] = pieceValue[row[j]]
                columnIndex = columnIndex + 1
            }
        }
    }

    return board
}

func BoardPiecesScore(board [13][12]int) int {
    score := 0;
    // Function same as piecesScore, this is for efficiency consideration
    for row := 1; row <= ROW_HEIGHT; row++ {
        for column := 1; column <= ROW_WIDTH; column++{
            piece := board[row][column];
            if piece > 0 {
                score = score + pieceScore[pieceStr[strconv.Itoa(piece)]];
            }
        }
    }
    return score;
}

