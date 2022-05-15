package dtos

import "time"

type UserToken struct {
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	UserID       uint
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AccessUUID   string `json:"access_uuid"`
	RefreshUUID  string `json:"refresh_uuid"`
}

func NewUserToken(userId uint, at, rt, auuid, ruuid string) *UserToken {
	return &UserToken{
		AccessToken:  at,
		RefreshToken: rt,
		AccessUUID:   auuid,
		RefreshUUID:  ruuid,
		UserID:       userId,
	}
}
