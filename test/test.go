package main

import (
	"fmt"

	"github.com/haupc/cartransplant/car/repository"
)

func main() {
	result, err := repository.GetPassengerTripRepo().FindPendingTrip(182, 0, 0, 10, 0)
	fmt.Println(result, err)
}
