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

//Set order local in queue.
func QueueSetLocal(floor int, buttonType int) {
	queue[floor][buttonType] = variables.LOCAL
}

//Return value of element in queue.
func QueueGet(floor int, buttonType int) variables.QueueOrderType {
	return queue[floor][buttonType]
}

//Pop order in queue.
func QueuePop(floor int, buttonType int) {
	if buttonType != 2 {
		queue[floor][buttonType] = variables.NONE
	} else {
		if queue[floor][buttonType] == variables.LOCAL {
			queue[floor][buttonType] = variables.NONE
		}
	}
}

//Initialize queue.
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

//Remove order from queue.
func QueueRemoveOrder(floor int, currentDirection ElevDir) {
	QueuePop(floor, int(Cab))
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

// Returns true if the there exist an order on current floor with same direction or no
//direction beyond current floor
func QueueCheckCurrentFloorSameDir(currentFloor int, currentDirection ElevDir) bool {
	//Check current floor same direction
	if queue[currentFloor][Cab] == variables.LOCAL {
		return true
	} else if (currentDirection == Up || currentDirection == Stop) && queue[currentFloor][HallUp] == variables.LOCAL {
		return true
	} else if (currentDirection == Down || currentDirection == Stop) && queue[currentFloor][HallDown] == variables.LOCAL {
		return true
	}

	//Check current floor no orders beyond
	if currentDirection == Up && QueueCheckAbove(currentFloor) == false {
		return true
	}
	if currentDirection == Down && QueueCheckBelow(currentFloor) == false {
		return true
	}

	return false
}

//Prints an illustration of queue with local elements to terminal.
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

//Returns true if there exists an order below current floor.
func QueueCheckBelow(currentFloor int) bool {
	if currentFloor == 0 {
		return false
	}
	for floor := currentFloor - 1; floor > -1; floor-- {
		for button := 0; button < variables.N_BUTTON_TYPES; button++ {
			if queue[floor][button] == variables.LOCAL {
				return true
			}
		}

	}
	return false
}

//Returns true if the exists an order above current floor.
func QueueCheckAbove(currentFloor int) bool {
	if currentFloor == variables.N_FLOORS-1 {
		return false
	}
	for floor := currentFloor + 1; floor < variables.N_FLOORS; floor++ {
		for button := 0; button < variables.N_BUTTON_TYPES; button++ {
			if queue[floor][button] == variables.LOCAL {
				return true
			}
		}
	}
	return false
}

//Set elements in queue remote.
func QueueSetOrderRemote(floor int, button int) {
	queue[floor][button] = variables.REMOTE
}

//Place new order as remote in queue.
func QueueRecieveOrderRemote(order elevio.ButtonEvent) {
	orderT := int(order.Button)
	QueueSetOrderRemote(order.Floor, orderT)
	fmt.Println("Order added to queue")
	QueuePrintRemote()
}

//Prints an illustration of queue with remote elements to terminal.
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

//Returns true if queue is empty
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

//Make remote order local.
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
