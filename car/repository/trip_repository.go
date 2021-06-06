package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/golang/glog"
	"github.com/haupc/cartransplant/auth/config"
	"github.com/haupc/cartransplant/car/model"
	"github.com/haupc/cartransplant/geometry/dto"
	"github.com/haupc/cartransplant/grpcproto"
	"gorm.io/gorm"
)

var tripRepository *tripRepo

// tripRepo interact with db
type TripRepo interface {
	CreateTrip(route dto.RoutingDTO, userID string, carID int64, maxDistance int64, beginLeaveTime, endLeaveTime time.Time, priceEachKm int64, seat int32) error
	FindTrip(from *grpcproto.Point, to *grpcproto.Point, tripType int32) ([]model.Trip, error)
	GetTripByID(tripID int64) (*model.Trip, error)
	GetTripByUserID(userID string, state, startDate, endDate int32) ([]model.Trip, error)
	GetTakenSeatByTripID(tripID int64) int32
	Create(model *model.Trip) error
	FindLastTrip(model *model.Trip) (*model.Trip, error)
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

func (r *tripRepo) Create(model *model.Trip) error {
	return r.db.Create(model).Error
}

func (r *tripRepo) FindLastTrip(tripModel *model.Trip) (*model.Trip, error) {
	result := &model.Trip{}
	if err := r.db.Where(tripModel).Order("created_at DESC").Find(result).Error; err != nil {
		log.Printf("CreateTrip query - Error: %v", err)
		return nil, err
	}
	return result, nil
}

func (r *tripRepo) GetTakenSeatByTripID(tripID int64) int32 {
	var result int32
	if err := r.db.Raw("select sum(seat) from passenger_trip where trip_id = ? and state = 2", 26).Scan(&result).Error; err != nil {
		glog.V(3).Infof("GetTakenSeatByTripID - error: %v", err)
		return 0
	}
	return result
}
func (r *tripRepo) GetTripByUserID(userID string, state, startDate, endDate int32) ([]model.Trip, error) {
	var tripModels []model.Trip
	if startDate > 0 && startDate < endDate {
		startDateTime := time.Unix(int64(startDate), 0)
		endDateTime := time.Unix(int64(endDate), 0).AddDate(0, 0, 1)
		if err := r.db.Where("user_id = ? and state = ? and (begin_leave_time between ? and ? or end_leave_time between ? and ?)", userID, state, startDateTime, endDateTime, startDateTime, endDateTime).Find(&tripModels).Error; err != nil {
			glog.V(3).Infof("GetTripByUserID - error: %v", err)
			return nil, err
		}
	} else {
		if err := r.db.Where("user_id = ? and state = ?", userID, state).Find(&tripModels).Error; err != nil {
			glog.V(3).Infof("GetTripByUserID - error: %v", err)
			return nil, err
		}
	}
	return tripModels, nil
}

func (r *tripRepo) GetTripByID(tripID int64) (*model.Trip, error) {
	var tripModel model.Trip
	if err := r.db.First(&tripModel, tripID).Error; err != nil {
		log.Printf("GetTripByID query - Error: %v", err)
		return nil, err
	}
	return &tripModel, nil
}

func (r *tripRepo) CreateTrip(route dto.RoutingDTO, userID string, carID int64, maxDistance int64, beginLeaveTime, endLeaveTime time.Time, priceEachKm int64, seat int32) error {
	lineString := makeLineString(route)
	way_json, _ := json.Marshal(route)

	query := fmt.Sprintf("insert into public.trip (user_id, car_id, max_distance, way, way_json, begin_leave_time, end_leave_time, fee_each_km, seat, state, type)  values (?, ?, ? , %s, ?, ?, ?, ?, ?, 1, 2)", lineString)
	log.Printf("CreateTrip query: %s", query)
	if err := r.db.Exec(query, userID, carID, maxDistance, way_json, beginLeaveTime, endLeaveTime, priceEachKm, seat).Error; err != nil {
		log.Printf("CreateTrip query - Error: %v", err)
		return err
	}
	return nil
}

// FindTrip conditions:
// 1: summary distance < max_distance
// 2: eta trip from -> from between ...
func (r *tripRepo) FindTrip(from *grpcproto.Point, to *grpcproto.Point, tripType int32) ([]model.Trip, error) {
	var tripModel []model.Trip
	fromPoint := makePoint(from)
	toPoint := makePoint(to)

	condition := fmt.Sprintf("ST_Distance(%s, way) + ST_Distance(%s, way)", fromPoint, toPoint)

	query := fmt.Sprintf("select * from public.trip where %s < 2000*max_distance and type = %d order by %s asc limit 10", condition, tripType, condition)
	if err := r.db.Raw(query).Find(&tripModel).Error; err != nil {
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
