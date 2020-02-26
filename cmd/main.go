package main

import (
	//"github.com/TTK4145/Network-go/network/peers"

	"github.com/sanderfu/TTK4145-ElevatorProject/internal/networkmanager"
)

func main() {
	go networkmanager.NetworkManager()

	go networkmanager.TestSendingRedundant(10)
	go networkmanager.TestReceivingRedundant(25)
	for {

	}
}
