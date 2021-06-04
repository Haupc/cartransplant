package dto

import "firebase.google.com/go/messaging"

type FcmDTO struct {
	Token string            `json:"token"`
	Data  map[string]string `json:"data"`
}

func (f *FcmDTO) ToFcmMessage() *messaging.Message {
	return &messaging.Message{
		Token: f.Token,
		Data:  f.Data,
	}
}
