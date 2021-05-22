package main

import (
	"context"
	"fmt"

	"github.com/haupc/cartransplant/grpcproto"
)

func main() {
	p := &grpcproto.Point{
		Latitude:  "20.986387",
		Longitude: "105.793815",
	}
	r, e := client.GetGeomClient().GetCurrentAddress(context.Background(), p)
	fmt.Println(string(r.JsonResponse), e)
}
