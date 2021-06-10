package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/haupc/cartransplant/car/repository"
	"github.com/haupc/cartransplant/grpcproto"
	"github.com/haupc/cartransplant/notify/client"
)

var _notifyService *notifyService

type NotifyService interface {
	NotifyNewTrip(startPoint *grpcproto.Point)
	NotifyNewUserRegisterTrip(driverID string, driverTripID int32)
	NotifyDriverTakeTrip(userID string, userTripID int32)
}

type notifyService struct {
	provinceRepo repository.ProvinceRepo
}

func GetNotifyService() NotifyService {
	if _notifyService == nil {
		_notifyService = &notifyService{
			provinceRepo: repository.GetProvinceRepo(),
		}
	}
	return _notifyService
}

func (s *notifyService) NotifyNewTrip(startPoint *grpcproto.Point) {
	province := s.provinceRepo.GetProvinceByPoint(startPoint)
	pushNotiRQ := &grpcproto.PushNotifyReq{
		Notification: &grpcproto.NotifyMessage{
			CreatedTime: time.Now().Unix(),
			Title:       "Có chuyến đi mới",
			Message:     fmt.Sprintf("Có chuyến đi mới trong vùng hoạt động của bạn: %s", province.Name),
			Image:       "https://thaygiangcomai.com/wp-content/uploads/2019/05/car-icon.png",
			Topic:       province.Topic,
		},
	}
	_, err := client.GetNotifyClient().PushNotifyToTopic(context.Background(), pushNotiRQ)
	if err != nil {
		log.Printf("NotifyNewTrip - Error: %v", err)
	}
}

func (s *notifyService) NotifyNewUserRegisterTrip(driverID string, driverTripID int32) {
	pushNotiRQ := &grpcproto.PushNotifyReq{
		Notification: &grpcproto.NotifyMessage{
			CreatedTime: time.Now().Unix(),
			Title:       "Có hành khách đi mới",
			Message:     fmt.Sprintf("Có chuyến đi %d của bạn có thêm hành khách mới", driverTripID),
			Image:       "https://thaygiangcomai.com/wp-content/uploads/2019/05/car-icon.png",
			UserID:      driverID,
		},
	}
	_, err := client.GetNotifyClient().PushNotify(context.Background(), pushNotiRQ)
	if err != nil {
		log.Printf("NotifyNewTrip - Error: %v", err)
	}
}

func (s *notifyService) NotifyDriverTakeTrip(userID string, userTripID int32) {
	pushNotiRQ := &grpcproto.PushNotifyReq{
		Notification: &grpcproto.NotifyMessage{
			CreatedTime: time.Now().Unix(),
			Title:       "Đã tìm thấy driver",
			Message:     fmt.Sprintf("Có chuyến đi %d của bạn đã có tài xế nhận chuyến", userTripID),
			Image:       "https://thaygiangcomai.com/wp-content/uploads/2019/05/car-icon.png",
			UserID:      userID,
		},
	}
	_, err := client.GetNotifyClient().PushNotify(context.Background(), pushNotiRQ)
	if err != nil {
		log.Printf("NotifyNewTrip - Error: %v", err)
	}
}
