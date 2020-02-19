package queue

import (
	"fmt"
)

type Direction int

const N_FLOORS = 4

var queue [N_FLOORS][2]int

func Add_order(floor int, dir int) {

	if floor == 0 {
		dir = 1
	}
	if floor == N_FLOORS-1 {
		dir = 0
	}
	queue[floor][dir] = 1
	print_queue()
}

func Delete_order(floor int, dir int) {
	queue[floor][dir] = 0

}

func order_at_floor(floor int, dir int) bool {
	status := queue[floor][dir] == 1
	return status
}

func order_list_empty(floor int, dir int) bool {
	status := true
	for i := 0; i < N_FLOORS; i++ {
		if (queue[floor][1] == 1) || (queue[i][0] == 1) {
			status = false
		}
	}
	return status
}

func Direction_to_travel(floor int) int {
	direction := 0
	if orders_above(floor) {
		direction = 1
	} else if orders_below(floor) {
		direction = -1
	}
	return direction
}
func orders_below(floor int) bool {
	status := false
	for i := 0; i < floor; i++ {
		if (queue[i][1] == 1) || (queue[i][0] == 1) {
			status = true
		}
	}
	return status
}

func orders_above(floor int) bool {
	status := false
	for i := floor; i < N_FLOORS; i++ {
		if queue[i][0] == 1 || queue[i][1] == 1 {
			status = true
		}
	}
	return status
}

func print_queue() {
	for i := 0; i < N_FLOORS; i++ {
		for j := 0; j < 2; j++ {

			fmt.Printf("%+v", queue[i][j])

		}
	}
	fmt.Printf("\n")

}
