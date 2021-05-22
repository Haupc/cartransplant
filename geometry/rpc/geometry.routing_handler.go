package geometry

import (
	"context"
	"fmt"

	"github.com/haupc/cartransplant/grpcproto"
	"github.com/haupc/cartransplant/utils/httputils"
)

var routingPath = "/route/v1/car/%s,%s;%s,%s"

func (g *geometryServer) GetRouting(ctx context.Context, req *grpcproto.RouteRequest) (*grpcproto.JsonResponse, error) {
	// http://localhost:5000/route/v1/car/105.786201,20.978652;105.794931,20.980682

	routingPath = fmt.Sprintf(routingPath, req.From.Longitude, req.From.Latitude, req.To.Longitude, req.To.Latitude)
	routingUrl := routingHost + routingPath
	requestclient := httputils.NewHttpClient()
	// params := map[string]string{
	// 	"format": "json",
	// 	"lat":    req.GetLatitude(),
	// 	"lon":    req.GetLongitude(),
	// 	"zoom":   "21",
	// }
	// requestclient.SetParams(params)
	response, err := requestclient.Get(routingUrl)
	grpcResponse := &grpcproto.JsonResponse{
		JsonResponse: response,
	}
	return grpcResponse, err
}
