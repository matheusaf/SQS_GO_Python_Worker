package model

type MessageModel struct {
	Channel string `json:"channel" binding:"required"`
	Message string `json:"message" binding:"required"`
}
