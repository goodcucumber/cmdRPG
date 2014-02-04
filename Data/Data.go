package Data

import "../Map"

var Player [2]byte
var Wall [2]byte

var DispMap string
var DispStatus string
var DispDialog string

var X int
var Y int

var Log string

//
var CurrentMap *Map.Map
var Maps []*Map.Map

//
func init() {
	Player[0] = 'M'
	Player[1] = 'E'
	Wall[0] = '*'
	Wall[1] = '*'
	DispMap = ""
	DispDialog = ""
	DispStatus = ""
	X = 1
	Y = 1
	Log = ""
	CurrentMap = nil
}
