package model

import (
	"time"

	"github.com/haupc/cartransplant/grpcproto"
)

type Notification struct {
	UserID    string    `json:"user_id"`
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	Image     string    `json:"image"`
}

func (n Notification) ToRPCNotification() *grpcproto.NotifyMessage {
	return &grpcproto.NotifyMessage{
		UserID:      n.UserID,
		CreatedTime: n.CreatedAt.Unix(),
		Title:       n.Title,
		Message:     n.Message,
		Image:       n.Image,
	}
}
