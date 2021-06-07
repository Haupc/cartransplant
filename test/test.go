package main

import (
	"fmt"

	"github.com/haupc/cartransplant/car/service"
)

func main() {
	fmt.Println(service.GetTripService().CancelTrip("xLPrxESgCeggCEpQXlDvhfOYGyw2", 15))
}
