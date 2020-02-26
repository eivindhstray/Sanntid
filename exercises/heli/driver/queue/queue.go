package queue

import (
	"./elevio"
)

type Direction int

const N_FLOORS = 4

var queue [N_FLOORS]int

func getMotorDirection() int {
	return elevio.MotorDirection
}

func setMotorDirection(value int) {
	elevio.SetMotorDirection(value)
}

func queueCheckAbove(floor int) int {
	for i := floor; i < N_FLOORS; i++ {
		if queue[floor] > 0 {
			return 1
		}
	}
	return 0
}

func queueCheckBelow(floor int) int {
	for i := 0; i < floor; i++ {
		if queue[floor] > 0 {
			return 1
		}
	}
	return 0
}

func queueAddOrder() {

}
