package controller

import (
	"github.com/gin-gonic/gin"
	"tfr-game/dao"
	"tfr-game/helper"
)

//UserHandler 用户控制器
type UserHandler struct {
	UserRoomDao dao.UserRoomDao
}

func NewUserHandler(dao dao.UserRoomDao) *UserHandler {
	return &UserHandler{
		UserRoomDao: dao,
	}
}

func (h *UserHandler) Users(c *gin.Context) {
	users := h.UserRoomDao.UserList()
	helper.Success(c, users)
}

func (h *UserHandler) Mine(c *gin.Context) {
	userInfo := h.UserRoomDao.Mine(c.GetString("uid"))
	helper.Success(c, userInfo)
}
