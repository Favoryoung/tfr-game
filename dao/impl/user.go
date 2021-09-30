package impl

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"sync"
	"tfr-game/dao"
	"tfr-game/helper/er"
	"tfr-game/model"
	"time"
)

type UserRoomDao struct {
	Users sync.Map
	Rooms sync.Map
}

func NewUserRoomDao() dao.UserRoomDao {
	implDao := new(UserRoomDao)
	implDao.cls()
	return implDao
}

//清除过期用户 2h执行一次
func (dao *UserRoomDao) cls() {
	go func() {
		for {
			dao.clsOnce()
			time.Sleep(12 * time.Hour)
		}
	}()
}

func (dao *UserRoomDao) clsOnce() {
	fmt.Println("[过期用户清理任务] 开始执行 ...")
	dao.Users.Range(func(key, value interface{}) bool {
		user := value.(model.User)
		if user.ExpireIn.Before(time.Now()) {
			dao.DelUser(user.Uid)
			fmt.Println("[删除过期用户] " + user.Uid)
		}
		return true
	})
	fmt.Println("[过期用户清理任务] 执行完毕 ╰(￣▽￣)╭")
}

func (dao *UserRoomDao) DelUser(uid string) {
	//查询是否加入房间已加入,则先退出
	user, ok := dao.GetUser(uid)
	if !ok {
		return
	}
	if user.RoomInfo.RoomNo != "" {
		dao.OutRoom(uid, "")
	}
	dao.Users.Delete(uid)
}

func (dao *UserRoomDao) GetUser(uid string) (model.User, bool) {
	value, ok := dao.Users.Load(uid)
	if ok {
		return value.(model.User), true
	}

	return model.User{}, false
}

func (dao *UserRoomDao) OutRoom(uid string, masterUid string) (error, model.User, model.Room) {
	user, ok := dao.GetUser(uid)
	if !ok {
		return errors.New("未查询到该用户"), user, model.Room{}
	}
	if user.RoomInfo.RoomNo == "" {
		return nil, user, model.Room{}
	}

	roomNo := user.RoomInfo.RoomNo
	room, _ := dao.GetRoom(roomNo)
	if masterUid != "" && room.Master != masterUid {
		return errors.New("非房主无法执行该操作"), user, model.Room{}
	}
	delete(room.Players, user.RoomInfo.PlayerNo)

	if room.Master == uid {
		if len(room.Players) > 0 {
			for _, playerUid := range room.Players {
				newMasterUser, ok := dao.GetUser(playerUid)
				if !ok {
					room.Master = ""
				} else {
					room.Master = playerUid
					newMasterUser.RoomInfo.IsMaster = true
					dao.SaveUser(newMasterUser)
				}
				break
			}
		} else {
			room.Master = ""
		}
	}

	user.RoomInfo = model.RoomInfo{}
	dao.SaveUser(user)

	dao.SaveRoom(room)
	return nil, user, room
}

func (dao *UserRoomDao) SaveUser(user model.User) model.User {
	dao.Users.Store(user.Uid, user)
	return user
}

func newExpireIn() time.Time {
	currentTime := time.Now()
	m, _ := time.ParseDuration("24h")
	return currentTime.Add(m)
}

func (dao *UserRoomDao) UpdateUserExpireIn(u model.User) {
	expireIn := newExpireIn()
	u.ExpireIn = expireIn
	dao.SaveUser(u)

	if u.RoomInfo.RoomNo != "" {
		dao.UpdateExpireInByRoomNo(u.RoomInfo.RoomNo, expireIn)
	}
}

func newUser() model.User {
	user := model.User{}
	user.Uid = uuid.NewV4().String()
	user.ExpireIn = newExpireIn()
	return user
}

func (dao *UserRoomDao) NewUser() model.User {
	user := newUser()
	dao.Users.Store(user.Uid, user)
	return user
}

func (dao *UserRoomDao) UserList() map[string]model.User {
	users := make(map[string]model.User)

	dao.Users.Range(func(key, value interface{}) bool {
		user := value.(model.User)
		users[user.Uid] = user
		return true
	})

	return users
}

func (dao *UserRoomDao) Mine(uid string) model.UserInfo {
	user, ok := dao.GetUser(uid)
	if !ok {
		er.Panic(er.User, "用户失效,请尝试重新登录")
	}

	if user.RoomInfo.RoomNo == "" {
		return model.UserInfo{Mine: user}
	}

	room, _ := dao.GetRoom(user.RoomInfo.RoomNo)
	return model.UserInfo{Mine: user, Room: room}
}
