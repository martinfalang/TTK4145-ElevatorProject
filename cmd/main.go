package main

import (
	//"github.com/TTK4145/Network-go/network/peers"
	//"github.com/TTK4145/Network-go/network/conn"
	//"github.com/TTK4145/Network-go/network/bcast"
	//"github.com/TTK4145/Network-go/network/localip"
	//"github.com/TTK4145/Network-go/network/peers"
	//"github.com/TTK4145/Network-go/driver-go/elevio"
	//"github.com/sanderfu/TTK4145-ElevatorProject/internal/datatypes"

	"github.com/sanderfu/TTK4145-ElevatorProject/internal/hwmanager"

	"github.com/sanderfu/TTK4145-ElevatorProject/internal/networkmanager"
)

func main() {

	go networkmanager.NetworkManager()

	go networkmanager.TestSendingRedundant(10)
	go networkmanager.TestReceivingRedundant(25)

	hwmanager.Init(4)

	for {

	}

}
