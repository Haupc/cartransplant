package geometry

import (
	"context"
	"encoding/json"

	"github.com/golang/glog"
	"github.com/haupc/cartransplant/geometry/dto"
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
	if err != nil {
		glog.Errorf("GetCurrentAddress - Error: %v", err)
		return nil, err
	}

	var responseObj []dto.SearchAddressRawResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		glog.Errorf("GetCurrentAddress - Error: %v", err)
		return nil, err
	}
	normalizeObj := []dto.SearchAddressResponse{}
	for _, r := range responseObj {
		normalizeObj = append(normalizeObj, r.Normalize())
	}
	byteResponse, err := json.Marshal(normalizeObj)
	grpcResponse := &grpcproto.JsonResponse{
		JsonResponse: byteResponse,
	}
	return grpcResponse, err
}
