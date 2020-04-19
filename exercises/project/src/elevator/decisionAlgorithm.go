package elevator

import (
	"fmt"

	"../elevio"
	"../variables"
)

//decisionAlgorithm making a choice every time a new order is stashed into queue remote (on button push).

//Calculates cost of new order for N_ELEVATORS, finds best elevator (lowest cost).
//The elevator with the lowest cost makes the order local.
func DecisionAlgorithm(buttonPush elevio.ButtonEvent) {
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

		cost = buttonPush.Floor - Elev.ElevState[elevator][0]
		if cost < 0 {
			cost = -cost
		}

		correctFloor = buttonPush.Floor
		correctButton = int(buttonPush.Button)

		if Elev.ElevState[elevator][2] < 0 {
			cost = cost + variables.ELEV_OFFLINE
		}

		CostArray[elevator] = cost
		fmt.Println("Elevator #: ", elevator, "%n Cost: ", cost)

		/*
			for floors := 0; floors < variables.N_FLOORS; floors++ {
				for buttons := 0; buttons < variables.N_BUTTON_TYPES-1; buttons++ {
					if queue[floors][buttons] == variables.REMOTE {
						cost = floors - Elev.ElevState[elevator][0]
						if cost < 0 {
							cost = -cost
						}
						correctFloor = floors
						correctButton = buttons

						if Elev.ElevState[elevator][2] > 0 {
							cost = cost + variables.ELEV_OFFLINE
						}
						CostArray[elevator] = cost
						fmt.Println("Elevator #: ", elevator, "%n Cost: ", cost)
					}
				}
			}
		}*/

		//Find best elevator
		var bestElev int
		bestElev = 1
		for elevator := 2; elevator < variables.N_ELEVATORS+1; elevator++ {
			if CostArray[elevator] < CostArray[bestElev] {
				bestElev = elevator
			}

		}
		fmt.Println("Elevator 1 cost :", CostArray[1])
		fmt.Println("Elevator 2 cost :", CostArray[2])
		fmt.Println("Best elevator : ", bestElev)
		fmt.Println("Elevator ID : ", Elev.ElevID)
		fmt.Println("*------------------------_*")
		fmt.Println("Elev 1 at floor: ", Elev.ElevState[1][0])
		fmt.Println("Elev 2 at floor: ", Elev.ElevState[2][0])

		//Set local in queue of best elevator
		if bestElev == Elev.ElevID {
			queue[correctFloor][correctButton] = variables.LOCAL
		}
	}
}
