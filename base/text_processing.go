package base

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"unicode"

	"github.com/haupc/cartransplant/geometry/client"
	"github.com/haupc/cartransplant/geometry/dto"
	"github.com/haupc/cartransplant/grpcproto"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var normalizer = transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

func Normalize(str string) (string, error) {
	s, _, err := transform.String(normalizer, str)
	if err != nil {
		return "", err
	}
	return strings.ToLower(s), err
}

func GetLocationName(point *grpcproto.Point) string {
	addressTo, err := client.GetGeomClient().GetCurrentAddress(context.Background(), point)
	if err != nil {
		log.Printf("ToGrpcListUserTripResponse - GetCurrentAddress error: %v", err)
		return ""
	}
	var addressParsed dto.SearchAddressResponse
	json.Unmarshal(addressTo.JsonResponse, &addressParsed)
	return addressParsed.DisplayName
}
