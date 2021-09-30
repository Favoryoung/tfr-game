package impl

import (
	"tfr-game/model"
	"time"
)

func (dao *UserRoomDao) GetRoom(roomNo string) (model.Room, bool) {
	value, ok := dao.Rooms.Load(roomNo)
	if ok {
		return value.(model.Room), true
	}

	return model.Room{}, false
}

func (dao *UserRoomDao) SaveRoom(room model.Room) model.Room {
	dao.Rooms.Store(room.RoomNo, room)
	return room
}

func (dao *UserRoomDao) UpdateRoomExpireIn(room model.Room, newExpireIn time.Time) {
	room.ExpireIn = newExpireIn
	dao.Rooms.Store(room.RoomNo, room)
}

func (dao *UserRoomDao) UpdateExpireInByRoomNo(roomNo string, expireIn time.Time) {
	room, ok := dao.GetRoom(roomNo)
	if ok {
		dao.UpdateRoomExpireIn(room, expireIn)
	}
}

