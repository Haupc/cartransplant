package dto

type Address struct {
	Road        string `json:"road"`
	Suburb      string `json:"suburb"`
	City        string `json:"city"`
	County      string `json:"county"`
	Postcode    string `json:"postcode"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
}

type SearchAddressRawResponse struct {
	Boundingbox []string `json:"boundingbox,omitempty"`
	Class       string   `json:"class,omitempty"`
	DisplayName string   `json:"display_name,omitempty"`
	Importance  float64  `json:"importance,omitempty"`
	Lat         string   `json:"lat,omitempty"`
	Licence     string   `json:"licence,omitempty"`
	Lon         string   `json:"lon,omitempty"`
	OsmID       int      `json:"osm_id,omitempty"`
	OsmType     string   `json:"osm_type,omitempty"`
	PlaceID     int      `json:"place_id,omitempty"`
	Type        string   `json:"type,omitempty"`
	Address     Address  `json:"address,omitempty"`
}

type SearchAddressResponse struct {
	DisplayName string `json:"display_name"`
	Lat         string `json:"latitude"`
	Lon         string `json:"longitude"`
	OsmID       int    `json:"osm_id"`
	OsmType     string `json:"osm_type"`
}

func (s SearchAddressRawResponse) Normalize() SearchAddressResponse {
	return SearchAddressResponse{
		DisplayName: s.DisplayName,
		Lat:         s.Lat,
		Lon:         s.Lon,
		OsmID:       s.OsmID,
		OsmType:     s.OsmType,
	}
}
