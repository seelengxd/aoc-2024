package utils

type Direction struct {
	X int
	Y int
}

var DIRECTIONS = map[string]Direction{
	"down":       {X: 1, Y: 0},
	"up":         {X: -1, Y: 0},
	"right":      {X: 0, Y: 1},
	"left":       {X: 0, Y: -1},
	"down_right": {X: 1, Y: 1},
	"down_left":  {X: 1, Y: -1},
	"up_left":    {X: -1, Y: -1},
	"up_right":   {X: -1, Y: 1},
}

var NON_DIAGONAL_DIRECTIONS = []Direction{
	DIRECTIONS["up"],
	DIRECTIONS["down"],
	DIRECTIONS["left"],
	DIRECTIONS["right"],
}

var ChangeDirectionClockwiseMap = map[Direction]Direction{
	DIRECTIONS["up"]:    DIRECTIONS["right"],
	DIRECTIONS["right"]: DIRECTIONS["down"],
	DIRECTIONS["down"]:  DIRECTIONS["left"],
	DIRECTIONS["left"]:  DIRECTIONS["up"],
}

var ChangeDirectionAntiClockwiseMap = map[Direction]Direction{
	DIRECTIONS["right"]: DIRECTIONS["up"],
	DIRECTIONS["down"]:  DIRECTIONS["right"],
	DIRECTIONS["left"]:  DIRECTIONS["down"],
	DIRECTIONS["up"]:    DIRECTIONS["left"],
}
