package repository

import (
	adsgo "github.com/Ivlay/ads-go"
	"github.com/jmoiron/sqlx"
)

type CommentsPg struct {
	db *sqlx.DB
}

func NewCommentsPg(db *sqlx.DB) *CommentsPg {
	return &CommentsPg{
		db: db,
	}
}

func (r *CommentsPg) CreateComment() (adsgo.Comments, error) {
	return adsgo.Comments{}, nil
}

func (r *CommentsPg) DeleteComment(id int) error {
	return nil
}

func (r *CommentsPg) GetCommentsByAdsId(adsId int) ([]adsgo.Comments, error) {
	return []adsgo.Comments{}, nil
}
