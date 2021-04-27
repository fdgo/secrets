package router

import (
	"github.com/gin-gonic/gin"
	"secret/control/user"
	"secret/utils/auth"
	"secret/utils/config"
	"secret/utils/middle"
)

func UserRouter(rf *gin.RouterGroup) {
	ru := rf.Group("/user")
	{
		ru.POST("/regist", user.Regist)              //注册
		ru.POST("/login", user.Login)                //登陆
		ru.GET("/devices/:userId", user.Devices)     //获取所有的设备号
		ru.POST("/device/delete", user.DeleteDevice) //删除设备号
		rus := ru.Use(middle.Auth(&auth.JwtToken{ //
			SigningKey: []byte(config.GloCfg.GetString("jwt.secretKey")),
		}))
		rus.POST("/device/bind", user.BindDevice)                  //绑定设备
		rus.POST("/friend/join", user.JoinFriend)                  //添加好友
		rus.POST("/group/new", user.NewGroup)                      //新建群拉好友
		rus.POST("/group/friend/join", user.GroupJoinFriend)       //群加好友
		rus.POST("/group/friend/kichout", user.GroupKickOutFriend) //组里删除好友
		rus.POST("/group/delete", user.GroupDelete)                //解散群
	}
}
