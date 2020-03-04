package main

import (
	//"github.com/TTK4145/Network-go/network/peers"
	//"github.com/TTK4145/Network-go/network/conn"
	//"github.com/TTK4145/Network-go/network/bcast"
	//"github.com/TTK4145/Network-go/network/localip"
	//"github.com/TTK4145/Network-go/network/peers"
	//"github.com/TTK4145/Network-go/driver-go/elevio"
	//"github.com/sanderfu/TTK4145-ElevatorProject/internal/datatypes"

	"time"

	"github.com/sanderfu/TTK4145-ElevatorProject/internal/hwmanager"

	"github.com/sanderfu/TTK4145-ElevatorProject/internal/networkmanager"
	"github.com/sanderfu/TTK4145-ElevatorProject/internal/ordermanager"
)

//"github.com/TTK4145/Network-go/network/peers"

func main() {

	go networkmanager.NetworkManager()

	go networkmanager.TestSendingRedundant(10)
	go networkmanager.TestReceivingRedundant(25)

	hwmanager.Init(4)

	go ordermanager.OrderManager()
	go ordermanager.ConfigureAndRunTest()
	for {
		time.Sleep(10 * time.Second)
	}

}
