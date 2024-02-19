package model

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type Post struct {
	Base
	PostAPI
	User *User `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID"`
}

type PostAPI struct {
	UserID     *uuid.UUID `json:"user_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
	Tweet      string     `json:"tweet,omitempty" gorm:"not null"`
	PictureUrl *string    `json:"picture_url,omitempty"`
}

type PostRequest struct {
	UserID  *uuid.UUID            `form:"user_id,omitempty"`
	Tweet   string                `form:"tweet,omitempty"`
	Picture *multipart.FileHeader `form:"picture"`
}
