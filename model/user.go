package model

// 用户信息返回

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Role     int    `json:"role"`
	Gender   int    `json:"gender"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Password string `json:"password"`
}

type UserResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Birthday int64  `json:"birthday"`
	Gender   int    `json:"gender"`
	Role     int    `json:"role"`
}

// 用户登录参数

type LoginUserForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	Password string `form:"password" json:"password" binging:"required,min=3,max=10"`
}
