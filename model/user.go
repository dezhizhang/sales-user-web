package model

// 用户信息返回

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
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password string `form:"password" json:"password" binging:"required,min=3,max=10"`
}
