package main

import (
	"fmt"

	"./elevio"
	"./queue"
)

func main() {

	numFloors := 4

	elevio.Init("localhost:15657", numFloors)

	var d elevio.MotorDirection = elevio.MD_Up
	//elevio.SetMotorDirection(d)

	drv_buttons := make(chan elevio.ButtonEvent)
	drv_floors := make(chan int)
	drv_obstr := make(chan bool)
	drv_stop := make(chan bool)

	go elevio.PollButtons(drv_buttons)
	go elevio.PollFloorSensor(drv_floors)
	go elevio.PollObstructionSwitch(drv_obstr)
	go elevio.PollStopButton(drv_stop)

	for i := 0; i < queue.N_FLOORS; i++ {
		for j := 0; j < 2; j++ {
			queue.Delete_order(i, j)
		}
	}

	for {
		select {
		case a := <-drv_buttons:
			fmt.Printf("%+v\n", a)
			elevio.SetButtonLamp(a.Button, a.Floor, true)
			queue.Add_order(a.Floor, int(a.Button))
			dir := queue.Direction_to_travel(a.Floor)
			d = elevio.MotorDirection(dir)
			fmt.Printf("%+v\n", d)
			elevio.SetMotorDirection(d)

		case a := <-drv_floors:
			fmt.Printf("%+v\n", a)
			elevio.SetFloorIndicator(a)
			queue.Delete_order(a, 0)
			queue.Delete_order(a, 1)
			dir := queue.Direction_to_travel(a)
			d = elevio.MotorDirection(dir)
			fmt.Printf("%+v\n", d)
			elevio.SetMotorDirection(d)

		case a := <-drv_obstr:
			fmt.Printf("%+v\n", a)
			if a {
				elevio.SetMotorDirection(elevio.MD_Stop)
			} else {
				elevio.SetMotorDirection(d)
			}

		case a := <-drv_stop:
			fmt.Printf("%+v\n", a)
			for f := 0; f < numFloors; f++ {
				for b := elevio.ButtonType(0); b < 3; b++ {
					elevio.SetButtonLamp(b, f, false)
				}
			}
		}
	}
}
