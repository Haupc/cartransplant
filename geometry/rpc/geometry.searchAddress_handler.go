package geometry

import (
	"context"

	"github.com/golang/glog"
	"github.com/haupc/cartransplant/grpcproto"
	"github.com/haupc/cartransplant/utils/httputils"
)

const searchAdressPath = "/nominatim/search.php"

var searchAddressUrl = locationHost + searchAdressPath

func (g *geometryServer) SearchAddress(ctx context.Context, req *grpcproto.SearchRequest) (*grpcproto.JsonResponse, error) {
	// http://localhost/nominatim/reverse.php?format=json&lat=20.9785698&lon=105.7862095&zoom=21
	requestclient := httputils.NewHttpClient()
	params := map[string]string{
		"format": "json",
		"q":      req.GetQuery(),
		// "addressdetails": "1",
	}
	requestclient.SetParams(params)
	response, err := requestclient.Get(searchAddressUrl)
	if err != nil {
		glog.Errorf("GetCurrentAddress - Error: %v", err)
		return nil, err
	}
	grpcResponse := &grpcproto.JsonResponse{
		JsonResponse: response,
	}
	return grpcResponse, err
}
