package elevator

import (
	"../elevio"
	"../variables"
)

func DecisionChooseElevator(buttonPush elevio.ButtonEvent) int {
	var CostArray [variables.N_ELEVATORS + 1]int

	for elev := 1; elev < variables.N_ELEVATORS+1; elev++ {
		CostArray[elev] = 0
	}

	for elevator := 1; elevator < variables.N_ELEVATORS+1; elevator++ {

		cost := buttonPush.Floor - Elev.ElevState[elevator][0]
		if cost < 0 {
			cost = -cost
		}

		if Elev.ElevState[elevator][1] == variables.ELEV_OFFLINE {
			cost = cost + variables.ELEV_OFFLINE
		}

		CostArray[elevator] = cost
	}

	var bestElev int
	bestElev = 1
	for elevator := 2; elevator < variables.N_ELEVATORS+1; elevator++ {
		if CostArray[elevator] < CostArray[bestElev] {
			bestElev = elevator
		}

	}
	/* -------------Debugging tools-------------------------

		fmt.Println("Elevator 1 cost :", CostArray[1])
		fmt.Println("Elevator 2 cost :", CostArray[2])
		fmt.Println("Best elevator : ", bestElev)
		fmt.Println("Elevator ID : ", Elev.ElevID)
		fmt.Println("*------------------------_*")
		fmt.Println("Elev 1 at floor: ", Elev.ElevState[1][0])
		fmt.Println("Elev 2 at floor: ", Elev.ElevState[2][0])
	---------------------------------------------------------
	*/
	return bestElev

}
