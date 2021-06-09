package notify

import (
	"context"
	"log"

	"github.com/haupc/cartransplant/grpcproto"
)

func (n *notifyServer) UnRegisterTokenToTopic(ctx context.Context, req *grpcproto.RegisterTokenToTopicRequest) (*grpcproto.Bool, error) {
	log.Printf("RegisterTokenToTopic\n")
	tokens := n.userTokenRepo.GetAllTokenByUserID(req.UserID)
	result, err := n.fcmClient.UnsubscribeFromTopic(ctx, tokens, req.Topic)
	if err != nil {
		log.Printf("RegisterTokenToTopic - Error: %v", err)
		return nil, err
	}
	log.Printf("RegisterTokenToTopic - result: %v\n", result)
	return &grpcproto.Bool{Value: true}, nil
}
