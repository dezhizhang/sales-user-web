package forms

type LoginUserForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	Password string `form:"password" json:"password" binging:"password"`
}
