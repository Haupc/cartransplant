package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/haupc/cartransplant/auth/config"
	"github.com/haupc/cartransplant/base"
	"github.com/haupc/cartransplant/car/model"
	"github.com/haupc/cartransplant/geometry/dto"
	"github.com/haupc/cartransplant/grpcproto"
	"gorm.io/gorm"
)

var tripRepository *tripRepo

// tripRepo interact with db
type TripRepo interface {
	CreateTrip(route dto.RoutingDTO, userID int32, beginLeaveTime, endLeaveTime time.Time) error
	FindTrip(from *grpcproto.Point, to *grpcproto.Point, beginLeaveTime int64, endLeaveTime int64, opt int32) ([]model.Trip, error)
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

func (r *tripRepo) CreateTrip(route dto.RoutingDTO, userID int32, beginLeaveTime, endLeaveTime time.Time) error {
	lineString := makeLineString(route)
	way_json, _ := json.Marshal(route)

	query := fmt.Sprintf("insert into public.trip (user_id , way, way_json, begin_leave_time, end_leave_time)  values (? , %s, ?, ?, ?)", lineString)
	log.Printf("CreateTrip query: %s", query)
	if err := r.db.Exec(query, userID, way_json, beginLeaveTime, endLeaveTime).Error; err != nil {
		log.Printf("CreateTrip query - Error: %v", err)
		return err
	}
	return nil
}

func (r *tripRepo) FindTrip(from *grpcproto.Point, to *grpcproto.Point, beginLeaveTime int64, endLeaveTime int64, opt int32) ([]model.Trip, error) {
	var tripModel []model.Trip
	fromPoint := makePoint(from)
	toPoint := makePoint(to)
	beginTime := time.Unix(beginLeaveTime, 0)
	endTime := time.Unix(endLeaveTime, 0)
	var condition string
	switch opt {
	case base.NEAR_START:
		condition = fmt.Sprintf("ST_Distance(%s, way)", fromPoint)
	case base.NEAR_END:
		condition = fmt.Sprintf("ST_Distance(%s, way)", toPoint)
	default:
		condition = fmt.Sprintf("ST_Distance(%s, way) + ST_Distance(%s, way)", fromPoint, toPoint)
	}
	query := fmt.Sprintf("select * from public.trip where begin_leave_time between ? and ?  or end_leave_time between ? and ? order by %s asc limit 10", condition)
	if err := r.db.Raw(query, beginTime, endTime, beginTime, endTime).Find(&tripModel).Error; err != nil {
		log.Printf("CreateTrip query - Error: %v", err)
		return tripModel, err
	}
	log.Println("FindTrip - result:", tripModel)
	return tripModel, nil
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

func makePoint(p *grpcproto.Point) string {
	return fmt.Sprintf("st_transform(st_geomfromtext('POINT(%s %s)', 4326), 3857)", p.Longitude, p.Latitude)
}
