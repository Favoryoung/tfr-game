package model

import (
	"fmt"
	"tfr-game/helper"
	"time"
)

type User struct {
	Uid      string    `json:"uid"`       //uid
	ExpireIn time.Time `json:"expire_in"` //有效期截止时间
	RoomInfo RoomInfo  `json:"room_info"` //房间相关信息
}

type RoomInfo struct {
	RoomNo   string `json:"room_no"`   //房间号
	IsMaster bool   `json:"is_master"` //是否房主
	PlayerNo int    `json:"number"`    //几号玩家
	Item     string `json:"item"`      //所持物品
}

type UserInfo struct {
	Mine User
	Room Room
}

func (u User) Show() string {
	return fmt.Sprintf("uid: %v, 有效期至:%v", u.Uid, helper.FriendTime(u.ExpireIn))
}
