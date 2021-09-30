package dao

import "tfr-game/model"

type UserRoomDao interface {
	UserList() map[string]model.User
	GetUser(uid string) (model.User, bool)
	UpdateUserExpireIn(u model.User)
	NewUser() model.User
	Mine(uid string) model.UserInfo
	//GetById(id uint64) *model.User
}
