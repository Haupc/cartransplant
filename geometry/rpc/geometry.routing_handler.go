package geometry

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/golang/glog"
	"github.com/haupc/cartransplant/geometry/dto"
	"github.com/haupc/cartransplant/geometry/helper"
	"github.com/haupc/cartransplant/grpcproto"
	"github.com/haupc/cartransplant/utils/httputils"
)

const routingPath = "/route/v1/car/%s,%s;%s,%s"

func (g *geometryServer) GetRouting(ctx context.Context, req *grpcproto.RouteRequest) (*grpcproto.JsonResponse, error) {
	// http://localhost:5000/route/v1/car/105.786201,20.978652;105.794931,20.980682

	path := fmt.Sprintf(routingPath, req.From.Longitude, req.From.Latitude, req.To.Longitude, req.To.Latitude)
	routingUrl := routingHost + path
	requestclient := httputils.NewHttpClient()
	params := map[string]string{
		"steps":        "true",
		"alternatives": "true",
	}
	requestclient.SetParams(params)
	response, err := requestclient.Get(routingUrl)
	if err != nil {
		glog.Errorf("GetRouting - Error: %v", err)
		return nil, err
	}

	var routingResponse dto.RouteResponse
	err = json.Unmarshal(response, &routingResponse)
	if err != nil {
		glog.Errorf("GetRouting Decode Json - Error: %v", err)
		return nil, err
	}
	responseObject := helper.RouteResponseToDTO(&routingResponse)
	rpcResponse, err := json.Marshal(responseObject)
	if err != nil {
		glog.Errorf("GetRouting Encode Json - Error: %v", err)
		return nil, err
	}
	grpcResponse := &grpcproto.JsonResponse{
		JsonResponse: rpcResponse,
	}
	return grpcResponse, err
}
