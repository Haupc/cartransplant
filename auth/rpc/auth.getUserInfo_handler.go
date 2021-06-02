package auth

import (
	"context"
	"log"

	"github.com/haupc/cartransplant/auth/config"
	"github.com/haupc/cartransplant/auth/dto"
	"github.com/haupc/cartransplant/grpcproto"
)

func (a *authServer) GetUserInfo(ctx context.Context, req *grpcproto.GetUserInfoRequest) (*grpcproto.UserProfile, error) {
	var myData dto.Metadata
	dsnap, err := config.GetFireStoreClient().Collection("users").Doc(req.GetUserID()).Get(context.Background())
	if err != nil {
		log.Printf("GetUserInfo - Error: %+v", err)
		return nil, err
	}
	dsnap.DataTo(&myData)
	return myData.ToUserProfile(), nil
}
