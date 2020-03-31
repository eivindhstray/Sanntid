package main

import(
	"./elevio"
)

//global variables
const N_FLOORS = 4
const N_BUTTON_TYPES = 3

type OrderType int

type Direction int


var queue [N_FLOORS][N_BUTTON_TYPES]bool

const (
	HallUp   OrderType = 0
	HallDown           = 1
	Cab                = 2
)

type Order struct {
	orderT OrderType
	floor  int
}

var OrderToButtonTypesMap = map[OrderType]elevio.ButtonType{
	HallUp:   elevio.BT_HallUp,
	HallDown: elevio.BT_HallDown,
	Cab:      elevio.BT_Cab,
}



var elevator Elevator

const (
	Up   ElevDir = 1
	Down         = -1
	Stop         = 0
)

type Elevator struct {
	currentFloor int
	dir          ElevDir
}
