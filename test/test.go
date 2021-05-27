package main

import (
	"fmt"
	"time"
)

func main() {
	t, err := time.Parse(time.RFC3339, "1994-11-05T13:15:30Z07:00")
	fmt.Println(t, err)
}
