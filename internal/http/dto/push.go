package dto

type CreatePushRequest struct {
	Message string `json:"message" binding:"required"`
}
