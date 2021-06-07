package model

import (
	"github.com/haupc/cartransplant/grpcproto"
)

type Notification struct {
	UserID    string `json:"user_id"`
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Message   string `json:"message"`
	CreatedAt int64  `json:"created_at"`
	Image     string `json:"image"`
}

func (n Notification) ToRPCNotification() *grpcproto.NotifyMessage {
	return &grpcproto.NotifyMessage{
		UserID:      n.UserID,
		CreatedTime: n.CreatedAt,
		Title:       n.Title,
		Message:     n.Message,
		Image:       n.Image,
	}
}
