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

var queueLocal [variables.N_FLOORS][variables.N_BUTTON_TYPES]variables.QueueOrderType

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


func localQueueSetLocal(floor int, buttonType int) {
	queueLocal[floor][buttonType] = variables.LOCAL
}

func localQueueGet(floor int, buttonType int) variables.QueueOrderType {
	return queueLocal[floor][buttonType]
}

func localQueuePop(floor int, buttonType int) {
	queueLocal[floor][buttonType] = variables.NONE
}

func LocalQueueInit() {
	fmt.Println("Queue initializing")
	for floor := 0; floor < variables.N_FLOORS; floor++ {
		for button := 0; button < variables.N_BUTTON_TYPES; button++ {
			localQueuePop(floor, button)
		}
	}
	fmt.Println("Local Queue initialized!")
}

func localQueueRecieveOrder(order elevio.ButtonEvent) {
	orderT := int(order.Button)
	localQueueSetLocal(order.Floor, orderT)
	fmt.Println("Order added to queue")
	localQueuePrint()
}

func localQueueRemoveOrder(floor int, currentDirection ElevDir) {
	localQueuePop(floor, int(Cab))
	if !(localQueueCheckBelow(floor) || localQueueCheckAbove(floor)) {
		localQueuePop(floor, int(HallUp))
		localQueuePop(floor, int(HallDown))
		return
	}
	switch currentDirection {
	case Up:
		localQueuePop(floor, int(HallUp))
		if localQueueCheckAbove(floor) == false {
			localQueuePop(floor, int(HallDown))
		}
		break
	case Down:
		localQueuePop(floor, int(HallDown))
		if localQueueCheckBelow(floor) == false {
			localQueuePop(floor, int(HallUp))
		}
		break
	case Stop:
		localQueuePop(floor, int(HallUp))
		localQueuePop(floor, int(HallDown))
		break
	}
}

func localQueueReturnElevDir(currentFloor int, currentDirection ElevDir) ElevDir {
	switch currentDirection {
	case Up:
		if localQueueCheckAbove(currentFloor) == true {
			return currentDirection
		} else if localQueueCheckBelow(currentFloor) == true && localQueueCheckAbove(currentFloor) == false {
			return Down
		}
	case Down:
		if localQueueCheckBelow(currentFloor) == true {
			return currentDirection
		} else if localQueueCheckAbove(currentFloor) == true && localQueueCheckBelow(currentFloor) == false {
			return Up
		}
	case Stop:
		if localQueueCheckAbove(currentFloor) == true {
			return Up
		} else if localQueueCheckBelow(currentFloor) == true {
			return Down
		}
	}
	return Stop
}



// Returns true if the there exist an order on current floor with same direction or no
//direction beyond current floor
func localQueueCheckCurrentFloorSameDir(currentFloor int, currentDirection ElevDir) bool {
	//Check current floor same direction
	if queueLocal[currentFloor][Cab] == variables.LOCAL {
		return true
	} else if (currentDirection == Up || currentDirection == Stop) && queueLocal[currentFloor][HallUp] == variables.LOCAL {
		return true
	} else if (currentDirection == Down || currentDirection == Stop) && queueLocal[currentFloor][HallDown] == variables.LOCAL {
		return true
	}

	//Check current floor noe orders beyond
	if currentDirection == Up && localQueueCheckAbove(currentFloor) == false {
		return true
	}
	if currentDirection == Down && localQueueCheckBelow(currentFloor) == false {
		return true
	}

	return false
}

func localQueuePrint() {
	fmt.Println("\n   HallUp   HallDn    Cab  ")
	fmt.Println("-" + strings.Repeat("|-------|", variables.N_BUTTON_TYPES))
	for floor := variables.N_FLOORS - 1; floor > -1; floor-- {
		fmt.Print(floor)
		for button := 0; button < variables.N_BUTTON_TYPES; button++ {
			queuePos := queueLocal[floor][button]
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

func localQueueCheckBelow(currentFloor int) bool {
	if currentFloor == 0 {
		return false
	}
	for floor := currentFloor - 1; floor > -1; floor-- {
		for button := 0; button < variables.N_BUTTON_TYPES; button++ {
			if queueLocal[floor][button] == variables.LOCAL {
				return true
			}
		}

	}
	return false
}

func localQueueCheckAbove(currentFloor int) bool {
	if currentFloor == variables.N_FLOORS-1 {
		return false
	}
	for floor := currentFloor + 1; floor < variables.N_FLOORS; floor++ {
		for button := 0; button < variables.N_BUTTON_TYPES; button++ {
			if queueLocal[floor][button] == variables.LOCAL {
				return true
			}
		}
	}
	return false
}

func remoteQueueSetOrder(floor int, button int) {
	queueLocal[floor][button] = variables.REMOTE
}

func remoteQueueRecieveOrder(order elevio.ButtonEvent) {
	orderT := int(order.Button)
	remoteQueueSetOrder(order.Floor, orderT)
	fmt.Println("Order added to queue")
	remoteQueuePrint()
}

func remoteQueuePrint() {
	fmt.Println("\n   HallUp   HallDn    Cab  ")
	fmt.Println("-" + strings.Repeat("|-------|", variables.N_BUTTON_TYPES))
	for floor := variables.N_FLOORS - 1; floor > -1; floor-- {
		fmt.Print(floor)
		for button := 0; button < variables.N_BUTTON_TYPES; button++ {
			queuePos := queueLocal[floor][button]
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
