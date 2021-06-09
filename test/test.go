package main

import (
	"fmt"

	"github.com/haupc/cartransplant/car/repository"
)

func main() {
	fmt.Println(repository.GetDriverProvinceRepo().BatchDelete("13w", []int32{1, 13}))
}
