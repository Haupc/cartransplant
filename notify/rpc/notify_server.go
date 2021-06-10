package notify

import (
	"context"
	"time"

	"firebase.google.com/go/messaging"
	"github.com/golang/glog"
	"github.com/haupc/cartransplant/auth/config"
	"github.com/haupc/cartransplant/base"
	car_repo "github.com/haupc/cartransplant/car/repository"
	"github.com/haupc/cartransplant/grpcproto"
	"github.com/haupc/cartransplant/notify/model"
	"github.com/haupc/cartransplant/notify/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type notifyServer struct {
	fcmClient          *messaging.Client
	notifyRepo         repository.NotifyRepo
	userTokenRepo      repository.UserTokenRepo
	driverProvinceRepo car_repo.DriverProvinceRepo
}

var _notifyServer *notifyServer

func NewNotifyServer() *notifyServer {
	if _notifyServer == nil {
		_notifyServer = &notifyServer{
			fcmClient:          config.GetFcmClient(),
			notifyRepo:         repository.GetNotifyRepo(),
			userTokenRepo:      repository.GetUserTokenRepo(),
			driverProvinceRepo: car_repo.GetDriverProvinceRepo(),
		}
	}
	return _notifyServer
}

// // ------------------------------
// // GetNotify get notifications by userID, limit & offset
func (n *notifyServer) GetNotify(ctx context.Context, req *grpcproto.GetNotifyRequest) (resp *grpcproto.GetNotifyResponse, err error) {

	md := base.RPCMetadataFromIncoming(ctx)
	glog.V(3).Infof("GetNotify - metadata: %v", md)
	notifications, err := n.notifyRepo.GetAllNotifyByUserID(md.UserID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	respose := &grpcproto.GetNotifyResponse{
		Notifications: []*grpcproto.NotifyMessage{},
	}
	for _, notify := range notifications {
		respose.Notifications = append(respose.Notifications, notify.ToRPCNotification())
	}
	return respose, nil
}

// ------------------------------
// AddUserToken as the func name
func (n *notifyServer) AddUserToken(ctx context.Context, req *grpcproto.AddUserTokenReq) (resp *grpcproto.AddUserTokenResp, err error) {
	// pre-exec check
	md := base.RPCMetadataFromIncoming(ctx)
	if req == nil || req.Token == "" {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if err = n.userTokenRepo.SaveUserToken(md.UserID, req.Token); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &grpcproto.AddUserTokenResp{Code: 1}, nil
}

// ------------------------------
// PushNotify save notify to db, query user tokens & push to firebase
func (n *notifyServer) PushNotify(ctx context.Context, req *grpcproto.PushNotifyReq) (resp *grpcproto.PushNotifyResp, err error) {
	// pre-exec check
	if req == nil || req.Notification == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	// save noti to db
	noti := &model.Notification{
		UserID:    req.Notification.UserID,
		Title:     req.Notification.Title,
		Message:   req.Notification.Message,
		CreatedAt: time.Now().Unix(),
		Image:     req.Notification.Image,
	}
	if err = repository.GetNotifyRepo().SaveNotification(noti); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// get all token
	tokens := n.userTokenRepo.GetAllTokenByUserID(req.Notification.UserID)
	msg := &messaging.MulticastMessage{
		Tokens: tokens,
		Notification: &messaging.Notification{
			Body:     req.Notification.Message,
			Title:    req.Notification.Title,
			ImageURL: req.Notification.Image,
		},
	}
	// send message
	_, err = n.fcmClient.SendMulticast(ctx, msg)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &grpcproto.PushNotifyResp{Code: 1}, nil
}
