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

var queue [variables.N_FLOORS][variables.N_BUTTON_TYPES]variables.QueueOrderType

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

func QueueSetLocal(floor int, buttonType int) {
	queue[floor][buttonType] = variables.LOCAL
}

func QueueSetRemote(floor int, button int) {
	queue[floor][button] = variables.REMOTE
}

func QueuePop(floor int, buttonType int) {
	queue[floor][buttonType] = variables.NONE
}

func LocalQueueInit() {
	for floor := 0; floor < variables.N_FLOORS; floor++ {
		for button := 0; button < variables.N_BUTTON_TYPES; button++ {
			QueuePop(floor, button)
		}
	}
}

//Place new order in queue as local.
func QueueRecieveOrderLocal(order elevio.ButtonEvent) {
	orderT := int(order.Button)
	QueueSetLocal(order.Floor, orderT)
	fmt.Println("Order added to local queue")
	QueuePrintLocal()
}

func QueueRecieveOrderRemote(order elevio.ButtonEvent) {
	orderT := int(order.Button)
	QueueSetRemote(order.Floor, orderT)
	fmt.Println("Order added to remote queue")
	QueuePrintRemote()
}

func QueueRemoveOrder(floor int, currentDirection ElevDir, ID int) {
	if ID == Elev.ElevID {
		QueuePop(floor, int(Cab))
	}
	if !(queueCheckBelow(floor) || queueCheckAbove(floor)) {
		QueuePop(floor, int(HallUp))
		QueuePop(floor, int(HallDown))
		return
	}
	switch currentDirection {
	case Up:
		QueuePop(floor, int(HallUp))
		if queueCheckAbove(floor) == false {
			QueuePop(floor, int(HallDown))
		}
		break
	case Down:
		QueuePop(floor, int(HallDown))
		if queueCheckBelow(floor) == false {
			QueuePop(floor, int(HallUp))
		}
		break
	case Stop:
		QueuePop(floor, int(HallUp))
		QueuePop(floor, int(HallDown))
		break
	}
}

//Return elevator direciton after searching through queue for orders.
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

func queueCheckCurrentFloorSameDir(currentFloor int, currentDirection ElevDir) bool {
	if queue[currentFloor][Cab] == variables.LOCAL {
		return true
	} else if (currentDirection == Up || currentDirection == Stop || queueCheckBelow(currentFloor) == false) && queue[currentFloor][HallUp] == variables.LOCAL {
		return true
	} else if (currentDirection == Down || currentDirection == Stop || queueCheckAbove(currentFloor) == false) && queue[currentFloor][HallDown] == variables.LOCAL {
		return true
	}

	return false
}

func queueCheckBelow(currentFloor int) bool {
	if currentFloor == 0 {
		return false
	}
	for floor := currentFloor - 1; floor > -1; floor-- {
		if queueCheckLocalCallOnFloor(floor) {
			return true
		}
	}
	return false
}

func queueCheckAbove(currentFloor int) bool {
	if currentFloor == variables.N_FLOORS-1 {
		return false
	}
	for floor := currentFloor + 1; floor < variables.N_FLOORS; floor++ {
		if queueCheckLocalCallOnFloor(floor) {
			return true
		}
	}
	return false
}

func QueueCheckEmpty(queueType variables.QueueOrderType) bool {
	empty := true
	for floor := 0; floor < variables.N_FLOORS; floor++ {
		for button := 0; button < variables.N_BUTTON_TYPES-1; button++ {
			if queue[floor][button] == queueType {
				return false
			}
		}
	}
	return empty
}

func queueMakeRemoteLocal() {
	for floor := 0; floor < variables.N_FLOORS; floor++ {
		for button := 0; button < variables.N_BUTTON_TYPES-1; button++ {
			if queue[floor][button] == variables.REMOTE {
				queue[floor][button] = variables.LOCAL
			}
		}
	}
}

func queueCheckLocalCallOnFloor(floor int) bool {
	for button := 0; button < variables.N_BUTTON_TYPES; button++ {
		if queue[floor][button] == variables.LOCAL {
			return true
		}
	}
	return false
}

//----------------Debugging tools--------------------
func QueuePrintRemote() {
	fmt.Println("Remote queue")
	fmt.Println("\n   HallUp   HallDn    Cab  ")
	fmt.Println("-" + strings.Repeat("|-------|", variables.N_BUTTON_TYPES))
	for floor := variables.N_FLOORS - 1; floor > -1; floor-- {
		fmt.Print(floor)
		for button := 0; button < variables.N_BUTTON_TYPES; button++ {
			queuePos := queue[floor][button]
			if queuePos == variables.REMOTE {
				fmt.Print("| ", "true ", " |")
			} else {
				fmt.Print("| ", "_____", " |")
			}
		}
		fmt.Println()
	}
	fmt.Print("-"+strings.Repeat("---------", variables.N_BUTTON_TYPES), "\n\n")
}

func QueuePrintLocal() {
	fmt.Println("Local queue")
	fmt.Println("\n   HallUp   HallDn    Cab  ")
	fmt.Println("-" + strings.Repeat("|-------|", variables.N_BUTTON_TYPES))
	for floor := variables.N_FLOORS - 1; floor > -1; floor-- {
		fmt.Print(floor)
		for button := 0; button < variables.N_BUTTON_TYPES; button++ {
			queuePos := queue[floor][button]
			if queuePos == variables.LOCAL {
				fmt.Print("| ", "true ", " |")
			} else {
				fmt.Print("| ", "_____", " |")
			}
		}
		fmt.Println()
	}
	fmt.Print("-"+strings.Repeat("---------", variables.N_BUTTON_TYPES), "\n\n")
}

//-----------------------------------------------------
