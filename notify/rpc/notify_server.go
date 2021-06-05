package notify

import (
	"context"

	"firebase.google.com/go/messaging"
	"github.com/haupc/cartransplant/auth/config"
	"github.com/haupc/cartransplant/base"
	"github.com/haupc/cartransplant/grpcproto"
	"github.com/haupc/cartransplant/notify/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type notifyServer struct {
	fcmClient *messaging.Client
}

var _notifyServer *notifyServer

func NewNotifyServer() *notifyServer {
	if _notifyServer == nil {
		_notifyServer = &notifyServer{
			fcmClient: config.GetFcmClient(),
		}
	}
	return _notifyServer
}

// ------------------------------
// GetNotify get notifications by userID, limit & offset
func (n *notifyServer) GetNotify(ctx context.Context, req *grpcproto.GetNotifyRequest) (resp *grpcproto.GetNotifyResponse, err error) {
	// pre-exec check
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	md := base.RPCMetadataFromIncoming(ctx)

	notifications, err := repository.GetNotifyRepo().GetAllNotifyRepoByUserID(ctx, md.UserID, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &grpcproto.GetNotifyResponse{Notifications: notifications}, nil
}

// ------------------------------
// AddUserToken as the func name
func (n *notifyServer) AddUserToken(ctx context.Context, req *grpcproto.AddUserTokenReq) (resp *grpcproto.AddUserTokenResp, err error) {
	// pre-exec check
	md := base.RPCMetadataFromIncoming(ctx)
	if req == nil || req.Token == "" {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	token := &grpcproto.UserToken{
		UserID: md.UserID,
		Token:  req.Token,
	}

	if err = repository.GetNotifyRepo().SaveUserToken(ctx, token); err != nil {
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
	if err = repository.GetNotifyRepo().SaveNotification(ctx, req.Notification); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// get all token
	tokens, err := repository.GetNotifyRepo().GetAllTokenByUserID(ctx, req.Notification.UserID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// make firebase message
	// create token list
	deviceTokens := make([]string, 0, len(tokens))
	for i := range tokens {
		deviceTokens = append(deviceTokens, tokens[i].Token)
	}
	msg := &messaging.MulticastMessage{
		Tokens: deviceTokens,
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
