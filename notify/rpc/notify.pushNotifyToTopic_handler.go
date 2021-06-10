package notify

import (
	"context"
	"time"

	"firebase.google.com/go/messaging"
	"github.com/haupc/cartransplant/grpcproto"
	"github.com/haupc/cartransplant/notify/model"
	"github.com/haupc/cartransplant/notify/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (n *notifyServer) PushNotifyToTopic(ctx context.Context, req *grpcproto.PushNotifyReq) (*grpcproto.Bool, error) {
	driverIDs, err := n.driverProvinceRepo.GetAllDriverIDByTopic(req.Notification.Topic)
	if err != nil {
		return nil, err
	}
	for _, driverID := range driverIDs {
		// save noti to db
		noti := &model.Notification{
			UserID:    driverID,
			Title:     req.Notification.Title,
			Message:   req.Notification.Message,
			CreatedAt: time.Now().Unix(),
			Image:     req.Notification.Image,
		}
		if err = repository.GetNotifyRepo().SaveNotification(noti); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	msg := &messaging.Message{
		Topic: req.Notification.Topic,
		Notification: &messaging.Notification{
			Body:     req.Notification.Message,
			Title:    req.Notification.Title,
			ImageURL: req.Notification.Image,
		},
	}
	// send message
	_, err = n.fcmClient.Send(ctx, msg)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &grpcproto.Bool{Value: true}, nil
}
