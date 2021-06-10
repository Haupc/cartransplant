package main

import (
	"encoding/json"
	"fmt"

	"github.com/haupc/cartransplant/grpcproto"
)

func main() {
	mjson := []byte(`{"type":[],"date":1623258000}`)
	var obj grpcproto.FindPendingTripRequest
	fmt.Println(json.Unmarshal(mjson, &obj), obj)
}
