package service

type TaskRequest struct {
	UserId      int    `json:"userId" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Status      string `json:"status,omitempty"`
}

type UserRequest struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
