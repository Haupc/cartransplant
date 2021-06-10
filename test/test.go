package main

import (
	"fmt"
)

func main() {
	distance := 156796.3
	fmt.Println(float32(int32(distance/100.0)) / 10)
}
