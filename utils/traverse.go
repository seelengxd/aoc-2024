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
