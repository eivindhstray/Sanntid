// New order and remove order on the following form:
// "NewOrder (1), floor, direction"
// "RomeveOrder (0), floor, direction"

//Assume queue has handled the newOrder and removeOrder such that queue is up to date.

package elevator

import (
	"fmt"

	"../elevio"
	"../variables"
)

//decisionAlgorithm making a choice every time a new order is stashed into queue remote.
//The algorithm tuns through the queue to find the remote order.
//The elevator with the lowest cost makes the order local.

func decisionAlgorithm(buttonPush elevio.ButtonEvent) {
	var CostArray [variables.N_ELEVATORS + 1]int
	var correctFloor int
	var correctButton int
	//Init cost array
	for elev := 1; elev < variables.N_ELEVATORS+1; elev++ {
		CostArray[elev] = 0
	}

	//Find the remote order and determine cost for every elevator
	for elevator := 1; elevator < variables.N_ELEVATORS+1; elevator++ {
		cost := 0
		for floors := 0; floors < variables.N_FLOORS; floors++ {
			for buttons := 0; buttons < variables.N_BUTTON_TYPES-1; buttons++ {
				if queueLocal[floors][buttons] == variables.REMOTE {
					cost = floors - Elev.ElevState[elevator][0]
					if cost < 0 {
						cost = -cost
					}
					correctFloor = floors
					correctButton = buttons

					if correctFloor >= Elev.ElevState[elevator][0] && Elev.ElevState[elevator][1] == -1 {
						cost = cost + 10
					}
					if correctFloor <= Elev.ElevState[elevator][0] && Elev.ElevState[elevator][1] == 1 {
						cost = cost + 10
					}

					CostArray[elevator] = cost
					fmt.Println("Elevator #: ", elevator, "%n Cost: ", cost)
				}
			}
		}
	}
	/*
		for i := 0; i < variables.N_ELEVATORS+1; i++ {
			fmt.Println("Elev ", i, "costvalue ", CostArray[i])
		}*/

	//Find best elevator
	var bestElev int
	bestElev = 1
	for elevator := 2; elevator < variables.N_ELEVATORS+1; elevator++ {
		if CostArray[elevator] < CostArray[bestElev] {
			bestElev = elevator
			//fmt.Println("Elevator:", elevator)
			//fmt.Println("Bestelev:", bestElev)
		}

	}
	fmt.Println("Best elevator : ", bestElev)
	fmt.Println("Elevator ID : ", Elev.ElevID)

	//Set local in queue of best elevator
	if bestElev == Elev.ElevID {
		//localQueueRecieveOrder(buttonPush)
		queueLocal[correctFloor][correctButton] = variables.LOCAL
	}
}
