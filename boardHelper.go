package pep

const (
	ROW_WIDTH  = 9
	ROW_HEIGHT = 10
	
	ADVISOR = 11
    ELEPHANT = 12
    SOLDIER = 15
    HORSE = 30
    CANNON = 35
    ROOK = 80
    GENERAL = 126
    RIVER = 5
    RP_MARGIN = 2
    COLOR = 128
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
    "80": "R",
    "30": "H",
    "12": "E",
    "11": "A",
    "126": "G",
    "35": "C",
    "15": "S",
}

func InitFullBoard() [13][12]int { // https://groups.google.com/forum/#!topic/golang-nuts/sPYRl4RHFdU
     board := [13][12] int {{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
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
  {-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},}
    return board
}

