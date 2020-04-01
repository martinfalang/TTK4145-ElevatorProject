package main

import (
	"fmt"
	"net"
	"os/exec"
	"strconv"
)

func findOpenPort() (int, net.Listener) {
	connPort := 16698
	addr := ":" + strconv.Itoa(connPort)
	fmt.Println(addr)
	listener, err := net.Listen("tcp", addr)

	for err != nil {
		fmt.Printf("Port %v already in use, increments..\n", connPort)
		connPort++
		addr = ":" + strconv.Itoa(connPort)
		fmt.Println(addr)
		listener, err = net.Listen("tcp", addr)
	}
	return connPort, listener
}

// Find two open ports by opening a connection and then closing it
func getPorts() (string, string) {
	watchdogPort, watchdogListener := findOpenPort()
	elevatorPort, elevatorListener := findOpenPort()
	defer watchdogListener.Close()
	defer elevatorListener.Close()
	return strconv.Itoa(watchdogPort), strconv.Itoa(elevatorPort)
}

func main() {

	watchdogPort, elevatorPort := getPorts()
	fmt.Printf("Watchdogport: %v\n elevport: %v\n", watchdogPort, elevatorPort)

	cmdWatchdog := exec.Command("gnome-terminal", "-e", "build/watchdog -watchdogport "+watchdogPort+" -elevport "+elevatorPort)
	cmdWatchdog.Run()

	cmdElevatorHardware := exec.Command("gnome-terminal", "-e", "./SimElevatorServer --port "+elevatorPort)
	cmdElevatorHardware.Run()

	cmdElevatorSoftware := exec.Command("gnome-terminal", "-e", "build/elevator -elevport "+elevatorPort+" -watchdogport "+watchdogPort)
	cmdElevatorSoftware.Run()

}
