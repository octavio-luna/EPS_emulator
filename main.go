package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	UART "github.com/octavio-luna/EPS_emulator/UART"
	"github.com/octavio-luna/EPS_emulator/internal/comm"
	"github.com/octavio-luna/EPS_emulator/internal/constants"
	"github.com/octavio-luna/EPS_emulator/internal/watchdog"
	"github.com/octavio-luna/EPS_emulator/state"
)

/*
Definitions:

PDU: Power Distribution Unit
PBU: Power Battery Unit
PCU: Power Conditioning Unit
PIU: ICEPSv2 Power Integrated Unit

IMTQv3: ISIS Magnetorquer v3

IMDU: ISIS Motor Drive Unit

It is recommended to use the “0x02 (0x03) – No-operation” command when using the IVID wildcard (0x00)

BID: Board ID is used  to identify a particular board within a multi-board system, such as the IMEPS, which will likely
contain multiple boards of the same system type. The board identifier is a sequentially numbered value starting from 1.
*/

// Startup function
func startup() (reader *os.File, writer *comm.Writer, watchdogResetChan chan bool, err error) {
	reader, writer, err = comm.New()
	watchdogResetChan = make(chan bool, 1)

	return reader, writer, watchdogResetChan, err
}

// Operation function
func operation(reader *os.File, writer *comm.Writer, watchdogResetChan chan bool) {
	eps := state.NewEps()

	go watchdog.NewWatchdogTimer(watchdogResetChan)

	eps.E.SetOpMode(constants.OpModeNominal)
	for {
		select {
		case <-watchdogResetChan:
			fmt.Println("Watchdog timer reset.")
			//TODO: Implement the system reset logic here
		default:
			scanner := bufio.NewScanner(reader)
			fmt.Println("Reading from FIFO...")
			if scanner.Scan() {
				// Reset the watchdog timer
				watchdogResetChan <- true

				// Process the received message
				message := scanner.Bytes()
				UARTMessage, err := UART.ParseUARTMessage(message)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error parsing message: %v\n", err)
					break
				}
				fmt.Printf("Received message: %+v\n", UARTMessage)

				// Process the received command and generate a response
				response, err := eps.ProcessCommand(UARTMessage)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error processing command: %v\n", err)
				}
				rsp := UART.ConstructUARTMessage(response)
				fmt.Printf("Generated response: \n%+v\n", response)

				// Write the response to the FIFO in a byte format

				//TODO: Check if I should write the response to the FIFO as a string or as a byte array
				if _, err := writer.Write([]byte(rsp)); err != nil {
					fmt.Fprintf(os.Stderr, "Error writing to FIFO: %v\n", err)
				}
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintf(os.Stderr, "Error reading from FIFO: %v\n", err)
			}
		}

		// Sleep for a while before the next poll
		fmt.Println("Waiting for new messages...")
		time.Sleep(time.Second)
	}
}

// Main function
func main() {
	reader, writer, watchdogResetChan, err := startup()
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	defer writer.Close()

	operation(reader, writer, watchdogResetChan)
}
