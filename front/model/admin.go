package model

type Admin struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}
