package main

import (
	"encoding/json"

	"github.com/sanderfu/TTK4145-ElevatorProject/internal/ordermanager"
	"github.com/sanderfu/TTK4145-ElevatorProject/internal/watchdog"

	//"github.com/TTK4145/Network-go/network/peers"
	//"github.com/TTK4145/Network-go/network/conn"
	//"github.com/TTK4145/Network-go/network/bcast"
	//"github.com/TTK4145/Network-go/network/localip"
	//"github.com/TTK4145/Network-go/network/peers"
	//"github.com/TTK4145/Network-go/driver-go/elevio"

	"fmt"
	"os"
	"time"

	"github.com/sanderfu/TTK4145-ElevatorProject/internal/datatypes"
	"github.com/sanderfu/TTK4145-ElevatorProject/internal/fsm"
	"github.com/sanderfu/TTK4145-ElevatorProject/internal/hwmanager"
	"github.com/sanderfu/TTK4145-ElevatorProject/internal/networkmanager"
)

func main() {
	fmt.Println("The Process ID is: ", os.Getpid())
	readConfig("./config.json")

	args := os.Args[1:]
	var lastPID string
	var resuming bool
	if len(args) > 0 {
		lastPID = args[0]
		resuming = true
	} else {
		lastPID = "NONE"
		resuming = false
	}
	fmt.Println("PID to resume from: ", lastPID)
	go watchdog.SenderNode()
	go networkmanager.NetworkManager()

	go ordermanager.OrderManager(resuming, lastPID)
	//go ordermanager.ConfigureAndRunTest()

	go hwmanager.HardwareManager()
	go fsm.FSM()

	for {
		time.Sleep(10 * time.Second)
	}

}

func readConfig(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&datatypes.Config)
	if err != nil {
		fmt.Println(err)
	}
}
