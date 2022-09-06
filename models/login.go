package models

type LoginRes struct {
	Token string `json:"token"`
	UserInfo UserInfo `json:"userInfo"`

}
