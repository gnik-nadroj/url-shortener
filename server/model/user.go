package model

type User struct {
	ID       string `json:"userId"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
