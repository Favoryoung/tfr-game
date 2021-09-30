package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tfr-game/dao"
	"tfr-game/model"
)

func User(userRoomDao dao.UserRoomDao) gin.HandlerFunc {
	return func(c *gin.Context) {
		uidCookie, err := c.Request.Cookie("uid")
		var uid string
		if err != nil {
			uid = setNewUser(c, userRoomDao.NewUser())
		} else {
			oldUid := uidCookie.Value

			//更新
			if oldUser, ok := userRoomDao.GetUser(oldUid); ok {
				userRoomDao.UpdateUserExpireIn(oldUser)
				uid = oldUser.Uid
				c.SetCookie("uid", uid, 15*86400, "/", "", false, false)
				fmt.Println("用户激活" + oldUser.Show())
			} else {
				uid = setNewUser(c, userRoomDao.NewUser())
			}
		}

		c.Set("uid", uid)
		c.Next()
	}
}

func setNewUser(c *gin.Context, user model.User) string {
	c.SetCookie("uid", user.Uid, 15*86400, "/", "", false, false)
	fmt.Println("新增用户" + user.Show())
	return user.Uid
}
