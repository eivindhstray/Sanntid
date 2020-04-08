package watchdog

import(
	"time"

	"../variables"
)

var WatchDogTimer *time.Timer

func WatchDogInit(){
	WatchDogTimer = time.NewTimer(0)
}

func WatchDogTimeNSeconds(timer chan <- bool){
	WatchDogTimer.Reset(variables.WATCHDOGINTERVAL * time.Second)
	<-WatchDogTimer.C	
	timer <- true
}

func TimerReset(timeout chan <- bool){
	WatchDogTimer.Reset(variables.WATCHDOGINTERVAL * time.Second)
	timeout <- true
}