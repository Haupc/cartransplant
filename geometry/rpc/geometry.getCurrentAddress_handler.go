package geometry

import (
	"context"

	"github.com/haupc/cartransplant/grpcproto"
	"github.com/haupc/cartransplant/utils/httputils"
)

const currentAdressPath = "/nominatim/reverse.php"

var currentAddressUrl = locationHost + currentAdressPath

func (g *geometryServer) GetCurrentAddress(ctx context.Context, req *grpcproto.Point) (*grpcproto.JsonResponse, error) {
	// http://localhost/nominatim/reverse.php?format=json&lat=20.9785698&lon=105.7862095&zoom=21
	requestclient := httputils.NewHttpClient()
	params := map[string]string{
		"format": "json",
		"lat":    req.GetLatitude(),
		"lon":    req.GetLongitude(),
		"zoom":   "21",
	}
	requestclient.SetParams(params)
	response, err := requestclient.Get(currentAddressUrl)
	grpcResponse := &grpcproto.JsonResponse{
		JsonResponse: response,
	}
	return grpcResponse, err
}
