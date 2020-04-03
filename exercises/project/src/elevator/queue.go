package elevator

import (
	"fmt"
	"strings"

	"../elevio"
	"../variables"
)

//queue
type OrderType int

type Direction int

var queue [variables.N_FLOORS][variables.N_BUTTON_TYPES]bool

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

//message

type ElevatorMessage struct {
	MessageType string
	Button      int //1 = hallup 2 = halldown 3 = cab
	Floor       int
}

func queueSet(floor int, buttonType int) {
	queue[floor][buttonType] = true
}

func queueGet(floor int, buttonType int) bool {
	return queue[floor][buttonType]
}

func queuePop(floor int, buttonType int) {
	queue[floor][buttonType] = false
}

func QueueInit() {
	for floor := 0; floor < variables.N_FLOORS; floor++ {
		for button := 0; button < variables.N_BUTTON_TYPES; button++ {
			queuePop(floor, button)
		}
	}
	fmt.Println("Init queue good")
}

func queueRecieveOrder(order elevio.ButtonEvent) {
	orderT := int(order.Button)
	queueSet(order.Floor, orderT)
	fmt.Println("Order added to queue")
	queuePrint()
}

func queueRemoveOrder(floor int, currentDirection ElevDir) {
	queuePop(floor, int(Cab))
	if !(queueCheckBelow(floor) || queueCheckAbove(floor)) {
		queuePop(floor, int(HallUp))
		queuePop(floor, int(HallDown))
		return
	}
	switch currentDirection {
	case Up:
		queuePop(floor, int(HallUp))
		break
	case Down:
		queuePop(floor, int(HallDown))
		break
	case Stop:
		queuePop(floor, int(HallUp))
		queuePop(floor, int(HallDown))
	}
}

func queueReturnElevDir(currentFloor int, currentDirection ElevDir) ElevDir {
	switch currentDirection {
	case Up:
		if queueCheckAbove(currentFloor) == true {
			return currentDirection
		} else if queueCheckBelow(currentFloor) == true {
			return Down
		}
	case Down:
		if queueCheckBelow(currentFloor) == true {
			return currentDirection
		} else if queueCheckAbove(currentFloor) == true {
			return Up
		}
	case Stop:
		if queueCheckAbove(currentFloor) == true {
			return Up
		} else if queueCheckBelow(currentFloor) == true {
			return Down
		}
	}
	return Stop
}

// Returns true if the there exist an order on current floor with same direction
func queueCheckCurrentFloorSameDir(currentFloor int, currentDirection ElevDir) bool {
	if queue[currentFloor][Cab] {
		return true
	} else if (currentDirection == Up || currentDirection == Stop) && queue[currentFloor][HallUp] {
		return true
	} else if (currentDirection == Down || currentDirection == Stop) && queue[currentFloor][HallDown] {
		return true
	}
	return false
}

func queuePrint() {
	fmt.Println("\n   HallUp   HallDn    Cab  ")
	fmt.Println("-" + strings.Repeat("|-------|", variables.N_BUTTON_TYPES))
	for floor := variables.N_FLOORS - 1; floor > -1; floor-- {
		fmt.Print(floor)
		for button := 0; button < variables.N_BUTTON_TYPES; button++ {
			i := queue[floor][button]
			if i {
				fmt.Print("| ", "true ", " |")
			} else {
				fmt.Print("| ", "_____", " |")
			}
		}
		fmt.Println()
	}
	fmt.Print("-"+strings.Repeat("---------", variables.N_BUTTON_TYPES), "\n\n")
}

func queueCheckBelow(currentFloor int) bool {
	if currentFloor == 0 {
		return false
	}
	for floor := 0; floor < currentFloor; floor++ {
		for button := 0; button < variables.N_BUTTON_TYPES; button++ {
			if queue[floor][button] == true {
				return true
			}
		}

	}
	return false
}

func queueCheckAbove(currentFloor int) bool {
	if currentFloor == variables.N_FLOORS-1 {
		return false
	}
	for floor := currentFloor; floor < variables.N_FLOORS; floor++ {
		for button := 0; button < variables.N_BUTTON_TYPES; button++ {
			if queue[floor][button] == true {
				return true
			}
		}
	}
	return false
}
