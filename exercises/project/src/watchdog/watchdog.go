package watchdog

import(
	"time"

	"../variables"
)

var WatchDogTimer *time.Timer

func WatchDogInit(){
	WatchDogTimer = time.NewTimer(0)
}

func WatchDogTimeNSeconds(timeout chan <- bool){
	WatchDogTimer.Reset(variables.WATCHDOGINTERVAL * time.Second)
	<-WatchDogTimer.C	
	timeout <- true
}

func WatchDogReset(timeout chan <- bool){
	WatchDogTimer.Reset(variables.WATCHDOGINTERVAL * time.Second)
	timeout <- true
}