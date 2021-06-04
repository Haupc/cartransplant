package main

import (
	"fmt"

	"github.com/haupc/cartransplant/car/service"
)

func main() {
	fmt.Println(service.GetTripService().ListDriverTrip("xLPrxESgCeggCEpQXlDvhfOYGyw2", 1))
}
