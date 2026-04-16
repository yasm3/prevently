package dto

type CreateUserSchema struct {
	Email string `json:"email" binding:"required,email"`
}

type CreateUserResponse struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	APIKey string `json:"apiKey"`
}

type GetUserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
