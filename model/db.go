package model

import (
	"log"
)

type Model struct {
}

var (
	cache map[string]UserData
	id    int64
)

func init() {
	cache = make(map[string]UserData, 64)

}

func (this *Model) GetLoginData(username string) (*UserData, error) {
	userData, ok := cache[username]
	if !ok {
		return &userData, ErrNotFound
	}
	return &userData, nil
}

func (this *Model) CreateUser(username string, hahSum, salt []byte) error {
	userData := UserData{
		Name:     username,
		PassWord: hahSum,
		Salt:     salt,
		Id:       id,
	}
	id++
	cache[username] = userData
	log.Println(cache)
	return nil
}

func (this *Model) UserGet(userID int64) (*UserData, error) {
	for _, k := range cache {
		if k.Id == userID {
			return &k, nil
		}
	}
	return nil, ErrNotFound
}
