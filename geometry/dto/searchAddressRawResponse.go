package dto

type SearchAddressRawResponse struct {
	Boundingbox []string `json:"boundingbox"`
	Class       string   `json:"class"`
	DisplayName string   `json:"display_name"`
	Importance  float64  `json:"importance"`
	Lat         string   `json:"lat"`
	Licence     string   `json:"licence"`
	Lon         string   `json:"lon"`
	OsmID       int      `json:"osm_id"`
	OsmType     string   `json:"osm_type"`
	PlaceID     int      `json:"place_id"`
	Type        string   `json:"type"`
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
