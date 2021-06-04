package notify

import (
	"context"

	"github.com/haupc/cartransplant/grpcproto"
	"github.com/haupc/cartransplant/notify/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type notifyServer struct {
}

func NewNotifyServer() *notifyServer {
	return &notifyServer{}
}

// ------------------------------
// GetNotify get notifications by userID, limit & offset
func (n *notifyServer) GetNotify(ctx context.Context, req *grpcproto.GetNotifyRequest) (resp *grpcproto.GetNotifyResponse, err error) {
	// pre-exec check
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	notifications, err := repository.GetNotifyRepo().GetAllNotifyRepoByUserID(ctx, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &grpcproto.GetNotifyResponse{Notifications: notifications}, nil
}
