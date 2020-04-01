package hwmanager

import (
	"github.com/TTK4145/driver-go/elevio"
	"github.com/sanderfu/TTK4145-ElevatorProject/internal/channels"
	"github.com/sanderfu/TTK4145-ElevatorProject/internal/configuration"
	"github.com/sanderfu/TTK4145-ElevatorProject/internal/datatypes"
)

////////////////////////////////////////////////////////////////////////////////
// Private variables
////////////////////////////////////////////////////////////////////////////////

var numberOfFloors int

////////////////////////////////////////////////////////////////////////////////
// Public functions
////////////////////////////////////////////////////////////////////////////////

func HardwareManager() {
	hwInit(configuration.Flags.ElevatorPort)

	go pollCurrentFloor()
	go pollHWORder()
	go updateOrderLights()
}

func SetElevatorDirection(dir int) {
	elevio.SetMotorDirection(elevio.MotorDirection(dir))
}

func SetDoorOpenLamp(value bool) {
	elevio.SetDoorOpenLamp(value)
}

////////////////////////////////////////////////////////////////////////////////
// Private functions
////////////////////////////////////////////////////////////////////////////////

func hwInit(port string) {
	numberOfFloors = configuration.Config.NumberOfFloors

	addr := ":" + port
	elevio.Init(addr, numberOfFloors)

	for floor := 0; floor < numberOfFloors; floor++ {
		setAllLightsAtFloor(floor, false)
	}
	SetDoorOpenLamp(false)

	// signal that HW init is finished
	channels.HMInitStatusFhmTfsm <- true
}

func pollCurrentFloor() {
	floorSensorChan := make(chan int)
	go elevio.PollFloorSensor(floorSensorChan)

	for {
		floor := <-floorSensorChan
		elevio.SetFloorIndicator(floor)
		channels.CurrentFloorFhmTfsm <- floor
	}
}

func pollHWORder() {
	btnChan := make(chan elevio.ButtonEvent)
	go elevio.PollButtons(btnChan)

	for {
		btnValue := <-btnChan
		hwOrder := datatypes.Order{
			Floor:     btnValue.Floor,
			OrderType: int(btnValue.Button),
		}
		channels.OrderFhmTom <- hwOrder
	}
}

func updateOrderLights() {
	for {
		select {
		case orderComplete := <-channels.ClearLightsFomThm:
			setAllLightsAtFloor(orderComplete.Floor, false)
		case orderRegistered := <-channels.SetLightsFomThm:
			elevio.SetButtonLamp(elevio.ButtonType(orderRegistered.OrderType),
				orderRegistered.Floor, true)
		}
	}
}

func setAllLightsAtFloor(floor int, value bool) {
	for btn := datatypes.OrderUp; btn <= datatypes.OrderInside; btn++ {
		if !(floor == 0 && btn == datatypes.OrderDown) &&
			!(floor == numberOfFloors-1 && btn == datatypes.OrderUp) {
			elevio.SetButtonLamp(elevio.ButtonType(btn), floor, value)
		}
	}
}
