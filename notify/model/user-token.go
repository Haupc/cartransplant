package model

type UserToken struct {
	ID     int64  `json:"id"`
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}

func (u UserToken) TableName() string {
	return "user_token"
}
