package utils

import (
	"context"
	"log"

	"github.com/haupc/cartransplant/auth/config"
	"github.com/haupc/cartransplant/auth/dto"
	"github.com/haupc/cartransplant/grpcproto"
)

func GetUserInfo(userID string) *grpcproto.UserProfile {
	var myData dto.Metadata
	dsnap, err := config.GetFireStoreClient().Collection("users").Doc(userID).Get(context.Background())
	if err != nil {
		log.Println("GetUserInfo - Error: %+v", err)
		return nil
	}
	dsnap.DataTo(&myData)
	return myData.ToUserProfile()
}
