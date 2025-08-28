package models

import (
	"strings"
	"time"
)

type Publication struct {
	Id             uint64    `json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Content        string    `json:"content,omitempty"`
	Likes          uint64    `json:"likes"`
	AuthorId       uint64    `json:"author_id,omitempty"`
	AuthorNickName string    `json:"author_nickname,omitempty"`
	CreatedOn      time.Time `json:"created_on,omitzero"`
	UpdatedOn      time.Time `json:"updated_on,omitzero"`
}

// NewPub: creates a new publication instance
func NewPub(title, content string, authorId uint64) (Publication, error) {
	now := time.Now()
	pub := Publication{
		Title:     strings.TrimSpace(title),
		Content:   strings.TrimSpace(content),
		AuthorId:  authorId,
		CreatedOn: now,
		UpdatedOn: now,
	}

	return pub, nil
}
