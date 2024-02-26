package watchdog

import (
	"fmt"
	"time"

	"github.com/octavio-luna/EPS_emulator/internal/constants"
)

func NewWatchdogTimer(watchdogResetChan chan bool) {
	remainingTime := time.Duration(constants.WatchDogTimeout) * time.Second
	for {
		// every time a message is received, the watchdog timer is reset
		// if the timer reaches the timeout, we should reset the system by sending a "false" message to the chan
		select {
		case <-time.After(1 * time.Second):
			//fmt.Println("Watchdog timer running. Remaining time:", remainingTime)
			remainingTime -= 1 * time.Second
			if remainingTime == 0 {
				fmt.Println("Watchdog timer expired.")

				watchdogResetChan <- false
			}
			if remainingTime == time.Duration(float64(constants.WatchDogTimeout)*float64(time.Second)*0.35) {
				//TTC peripherical reset tries to restore communication with the master. This mechanism is triggered
				//when the watchdog timer reaches 0.65 x watchdog timeout. When triggered only the communication peripheral internal
				//to the system microcontroller is reset.
				fmt.Println("TTC peripherical reset.")
				//TODO: Implement the TTC peripherical reset logic here
			}
		case <-watchdogResetChan:
			fmt.Println("Watchdog timer reset.")
			remainingTime = time.Duration(constants.WatchDogTimeout) * time.Second
		}
	}
}
