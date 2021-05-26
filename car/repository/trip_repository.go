package repository

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/haupc/cartransplant/auth/config"
	"github.com/haupc/cartransplant/geometry/dto"
	"gorm.io/gorm"
)

var tripRepository *tripRepo

// tripRepo interact with db
type TripRepo interface {
	CreateTrip(route dto.RoutingDTO, userID int32) error
}

type tripRepo struct {
	db *gorm.DB
}

// GettripRepo singleton trip repo
func GettripRepo() TripRepo {
	if tripRepository == nil {
		tripRepository = &tripRepo{
			config.GetDbConnection(),
		}
	}
	return tripRepository
}

func (r *tripRepo) CreateTrip(route dto.RoutingDTO, userID int32) error {
	lineString := makeLineString(route)
	way_json, _ := json.Marshal(route)

	query := fmt.Sprintf("insert into public.trip (user_id , way, way_json)  values (? , %s, ?)", lineString)
	log.Printf("CreateTrip query: %s", query)
	if err := r.db.Raw(query, userID, way_json).Error; err != nil {
		log.Printf("CreateTrip query - Error: %v", err)
		return err
	}
	return nil
}

func makeLineString(route dto.RoutingDTO) string {
	lineString := "st_transform(st_geomfromtext('LINESTRING("
	for _, step := range route.Routes[0].Steps {
		lineString += fmt.Sprintf("%s %s, ", step.Location.Longitude, step.Location.Latitude)
	}
	lineString = lineString[:len(lineString)-2] // strip ', '
	lineString += ")', 4326), 3857)"
	return lineString
}
