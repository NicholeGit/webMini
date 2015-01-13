package model

type UserData struct {
	Id       int64
	Name     string
	PassWord []byte
	Salt     []byte
}
