package main

import(
	"time"

	"./variables"
)

var WatchDogTimer *time.Timer

func WatchDogInit(){
	WatchDogTimer = time.NewTimer(0)
}

func WatchDogTimeNSeconds(){
	WatchDogTimer.Reset(variables.WATCHDOGINTERVAL * time.Second)
	<-WatchDogTimer.C	
	variables.COMMSALIVE = false
}

func WatchDogReset(){
	WatchDogTimer.Reset(variables.WATCHDOGINTERVAL * time.Second)
	variables.COMMSALIVE = true
}