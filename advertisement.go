package adsgo

import "github.com/lib/pq"

type Advertisement struct {
	Id          int            `json:"id" db:"id"`
	Title       string         `json:"title" db:"title" binding:"required"`
	Description string         `json:"description" db:"description" binding:"required"`
	Images      pq.StringArray `json:"images" db:"images"`
	CreatedAt   string         `json:"created_at" db:"created_at"`
	UpdatedAt   string         `json:"updated_at" db:"updated_at"`
	UserId      int            `json:"userId" db:"user_id,pk"`
}
