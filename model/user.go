package model

type UserResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Birthday int64  `json:"birthday"`
	Gender   int    `json:"gender"`
	Role     int    `json:"role"`
}
