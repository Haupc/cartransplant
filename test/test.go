package main

import (
	"fmt"
	"log"

	"github.com/haupc/cartransplant/utils/httputils"
)

func main() {
	requestclient := httputils.NewHttpClient()
	params := map[string]string{
		"format": "json",
		"q":      "Vu Xuan Thieu",
		// "addressdetails": "1",
	}
	requestclient.SetParams(params)
	response, err := requestclient.Get("http://localhost/nominatim/search.php?")
	if err != nil {
		log.Printf("GetCurrentAddress - Error: %v", err)
		// return nil, err
	}
	fmt.Println(string(response))
}
