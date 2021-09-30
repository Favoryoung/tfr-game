package model

import "time"

type Room struct {
	RoomNo           string         `json:"room_no" binding:"required,max=10"`    //房间号
	Master           string         `json:"master"`                               //房主
	PlayerNum        int            `json:"player_num" binding:"required,max=20"` //房间人数
	Players          map[int]string `json:"players"`                              //当前用户
	Items            []string       `json:"items" binding:"required"`
	Seconds          int            `json:"seconds" binding:"min=0"` //每次开始间隔
	Games            int            `json:"games"`                   //游戏局数
	CurrentStartTime time.Time      `json:"-"`                       //当前游戏开始时间
	ExpireIn         time.Time      `json:"-"`                       //到期时间
}
