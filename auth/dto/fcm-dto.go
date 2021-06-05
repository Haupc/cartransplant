package dto

import "firebase.google.com/go/messaging"

type FcmDTO struct {
	Token string            `json:"token"`
	Data  messaging.Message `json:"data"`
}
