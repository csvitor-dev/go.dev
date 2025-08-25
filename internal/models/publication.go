package models

import "time"

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
