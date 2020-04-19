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

func QueueSetOrderRemote(floor int, button int) {
	queue[floor][button] = variables.REMOTE
}

//Pop order in queue. Only pop cabcall if it is a local order.
func QueuePop(floor int, buttonType int) {

	queue[floor][buttonType] = variables.NONE
}

func LocalQueueInit() {
	fmt.Println("Queue initializing")
	for floor := 0; floor < variables.N_FLOORS; floor++ {
		for button := 0; button < variables.N_BUTTON_TYPES; button++ {
			QueuePop(floor, button)
		}
	}
	fmt.Println("Local Queue initialized!")
}

//Place new order in queue as local.
func QueueRecieveOrderLocal(order elevio.ButtonEvent) {
	orderT := int(order.Button)
	QueueSetLocal(order.Floor, orderT)
	fmt.Println("Order added to queue")
	QueuePrintLocal()
}

func QueueRecieveOrderRemote(order elevio.ButtonEvent) {
	orderT := int(order.Button)
	QueueSetOrderRemote(order.Floor, orderT)
	fmt.Println("Order added to queue")
	QueuePrintRemote()
}

func QueueRemoveOrder(floor int, currentDirection ElevDir, ID int) {
	if ID == Elev.ElevID {
		QueuePop(floor, int(Cab))
	}
	if !(QueueCheckBelow(floor) || QueueCheckAbove(floor)) {
		QueuePop(floor, int(HallUp))
		QueuePop(floor, int(HallDown))
		return
	}
	switch currentDirection {
	case Up:
		QueuePop(floor, int(HallUp))
		if QueueCheckAbove(floor) == false {
			QueuePop(floor, int(HallDown))
		}
		break
	case Down:
		QueuePop(floor, int(HallDown))
		if QueueCheckBelow(floor) == false {
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
func QueueReturnElevDir(currentFloor int, currentDirection ElevDir) ElevDir {
	switch currentDirection {
	case Up:
		if QueueCheckAbove(currentFloor) == true {
			return currentDirection
		} else if QueueCheckBelow(currentFloor) == true && QueueCheckAbove(currentFloor) == false {
			return Down
		}
	case Down:
		if QueueCheckBelow(currentFloor) == true {
			return currentDirection
		} else if QueueCheckAbove(currentFloor) == true && QueueCheckBelow(currentFloor) == false {
			return Up
		}
	case Stop:
		if QueueCheckAbove(currentFloor) == true {
			return Up
		} else if QueueCheckBelow(currentFloor) == true {
			return Down
		}
	}
	return Stop
}

func QueueCheckCurrentFloorSameDir(currentFloor int, currentDirection ElevDir) bool {
	if queue[currentFloor][Cab] == variables.LOCAL {
		return true
	} else if (currentDirection == Up || currentDirection == Stop || QueueCheckBelow(currentFloor) == false) && queue[currentFloor][HallUp] == variables.LOCAL {
		return true
	} else if (currentDirection == Down || currentDirection == Stop || QueueCheckAbove(currentFloor) == false) && queue[currentFloor][HallDown] == variables.LOCAL {
		return true
	}

	return false
}

func QueueCheckBelow(currentFloor int) bool {
	if currentFloor == 0 {
		return false
	}
	for floor := currentFloor - 1; floor > -1; floor-- {
		if QueueCheckLocalCallOnFloor(floor) {
			return true
		}
	}
	return false
}

func QueueCheckAbove(currentFloor int) bool {
	if currentFloor == variables.N_FLOORS-1 {
		return false
	}
	for floor := currentFloor + 1; floor < variables.N_FLOORS; floor++ {
		if QueueCheckLocalCallOnFloor(floor) {
			return true
		}
	}
	return false
}

func CheckQueueEmpty(queueType variables.QueueOrderType) bool {
	empty := true
	for floor := 0; floor < variables.N_FLOORS-1; floor++ {
		for button := 0; button < variables.N_BUTTON_TYPES; button++ {
			queuePos := queue[floor][button]
			if queuePos == queueType {
				return false
			}
		}
	}
	return empty
}

func QueueMakeRemoteLocal() {
	for floor := 0; floor < variables.N_FLOORS; floor++ {
		for button := 0; button < variables.N_BUTTON_TYPES-1; button++ {
			queuePos := queue[floor][button]
			if queuePos == variables.REMOTE {
				queue[floor][button] = variables.LOCAL
			}
		}
	}
}

func QueueCheckLocalCallOnFloor(floor int) bool {
	for button := 0; button < variables.N_BUTTON_TYPES; button++ {
		if queue[floor][button] == variables.LOCAL {
			return true
		}
	}
	return false
}

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
