package dto

import "github.com/haupc/cartransplant/grpcproto"

type Point struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func (p Point) ToGrpcPoint() *grpcproto.Point {
	return &grpcproto.Point{
		Latitude:  p.Latitude,
		Longitude: p.Longitude,
	}
}
