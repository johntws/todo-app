package dto

type LoginReq struct {
	Username string `form:"title" binding:"required"`
	Password string `form:"description" binding:"required"`
}

type LoginRes struct {
	Token string `json:"token"`
}
