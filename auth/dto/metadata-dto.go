package dto

import "github.com/haupc/cartransplant/grpcproto"

type Metadata struct {
	UserID   string `json:"user_id,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Email    string `json:"email,omitempty"`
	FullName string `json:"full_name,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

func (m *Metadata) ToUserProfile() *grpcproto.UserProfile {
	return &grpcproto.UserProfile{
		UserID:   m.UserID,
		Avatar:   m.Avatar,
		Email:    m.Email,
		FullName: m.FullName,
		Phone:    m.Phone,
	}
}
