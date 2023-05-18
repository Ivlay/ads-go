package adsgo

type Comments struct {
	Message   string `json:"message" db:"message" binding:"required"`
	UserId    int    `json:"userId" db:"user_id"`
	AdsId     int    `json:"adsId" db:"ads_id"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}
