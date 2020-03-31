package costAlgorithm

const N_ELEVATORS = 3

var COST int

//Information for each elevator on position and direction
const N_INFO = 2

var elevatorList [N_ELEVATORS][N_INFO]int

func getElevatorPos(elevator int) int {
	return elevatorList[elevator][0]
}

func getElevatorDir(elevator int) int {
	return elevatorList[elevator][1]
}

func calculateCost(elevator int, newOrder string) int {
	COST = 0
	if getElevatorPos(elevator) != newOrder.pos {
		COST = COST + 1
	}
	if getElevatorDir(elevator) == 1 && getElevatorPos(elevator) > newOrder.pos {
		COST = COST + 100
	}
	if getElevatorDir(elevator) == -1 && getElevatorPos(elevator) < newOrder.pos {
		COST = COST + 100
	}
	if getElevatorDir(elevator) == 0 {
		COST = COST + 10
	}
	return COST
}

func bestSuitedElevator(newOrder string) int {
	currentBestCost := 1000
	for i := 0; i < 3; i++ {
		tempCost := calculateCost(i, newOrder)
		if tempCost < currentBestCost {
			bestSuited := i
		}

	}
	return bestSuited
}

//Tanken her er i bunn og grunn å assigne ulike queuer til ulike heiser.
//Denne mappen skal holde N queuer (i vårt tilfelle 2) for så å markere jobber som sendes
//som queuer til hver enkelt heis.
//Sett fra driver (hver enkelt heis) kan dette bare sendes som en melding og oppdatere lokal
//queue.

//Gi meg gjerne en lyd om hva dere tenker Magz og Andy <3
