package user

import (
	"github.com/gin-gonic/gin"
	"secret/model/request/user"
	"secret/model/response"
	usersrv "secret/service/user"
	"secret/utils/constex"
	"secret/utils/validation"
	"strconv"
)

func Regist(c *gin.Context) {
	var in user.Regist
	if err := c.ShouldBindJSON(&in); err != nil {
		response.RespFailed(c, constex.ERROR_PARAM_ERROR, err.Error(), "请输入正确的参数类型!")
		return
	}
	valid := validation.Validation{}
	isok_param, _ := valid.Valid(&in)
	if !isok_param {
		for _, err := range valid.Errors {
			response.RespFailed(c, constex.ERROR_PARAM_ERROR, err.Error(), err.Message)
			return
		}
	}
	userdao := usersrv.NewUserGroupDao()
	err, userId, devId := userdao.Regist(in.Mail, in.Pwd)
	if err != nil {
		response.RespFailed(c, constex.ERROR_PARAM_ERROR, err.Error(), err.Error())
		return
	}
	response.RespSuccess(c, gin.H{
		"userId": userId,
		"devId":  devId,
	}, "注册成功!")
}
func Login(c *gin.Context) {
	var in user.Login
	if err := c.ShouldBindJSON(&in); err != nil {
		response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", "请输入正确的参数类型!")
		return
	}
	valid := validation.Validation{}
	isok_param, _ := valid.Valid(&in)
	if !isok_param {
		for _, err := range valid.Errors {
			response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", err.Message)
			return
		}
	}
	userdao := usersrv.NewUserGroupDao()
	err, ret := userdao.Login(in.Mail, in.Pwd, in.Device)
	if err != nil {
		response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", err.Error())
		return
	}
	response.RespSuccess(c, ret, "登陆成功!")
}

//枚举设备
func Devices(c *gin.Context) {
	uid, _ := strconv.ParseInt(c.Param("userId"), 10, 64)
	userdao := usersrv.NewUserGroupDao()
	err, devId := userdao.Devices(uid)
	if err != nil {
		response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", err.Error())
		return
	}
	response.RespSuccess(c, devId, "ok")
}

//删除设备
func DeleteDevice(c *gin.Context) {
	var in user.DeleteDevice
	if err := c.ShouldBindJSON(&in); err != nil {
		response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", "请输入正确的参数类型!")
		return
	}
	valid := validation.Validation{}
	isok_param, _ := valid.Valid(&in)
	if !isok_param {
		for _, err := range valid.Errors {
			response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", err.Message)
			return
		}
	}
	userdao := usersrv.NewUserGroupDao()
	err := userdao.DeleteDevice(in.UserId, in.DeviceId)
	if err != nil {
		response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", err.Error())
		return
	}
	response.RespSuccess(c, "ok", "request send ok!")
}

//绑定设备
func BindDevice(c *gin.Context) {
	var in user.DeviceBind
	if err := c.ShouldBindJSON(&in); err != nil {
		response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", "请输入正确的参数类型!")
		return
	}
	valid := validation.Validation{}
	isok_param, _ := valid.Valid(&in)
	if !isok_param {
		for _, err := range valid.Errors {
			response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", err.Message)
			return
		}
	}
	userdao := usersrv.NewUserGroupDao()
	err, devId := userdao.BindDevice(in.UserId, in.DeviceId)
	if err != nil {
		response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", err.Error())
		return
	}
	response.RespSuccess(c, devId, "request send ok!")
}

//添加好友
func JoinFriend(c *gin.Context) {
	var in user.JoinFriend
	if err := c.ShouldBindJSON(&in); err != nil {
		response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", "请输入正确的参数类型!")
		return
	}
	valid := validation.Validation{}
	isok_param, _ := valid.Valid(&in)
	if !isok_param {
		for _, err := range valid.Errors {
			response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", err.Message)
			return
		}
	}
	userdao := usersrv.NewUserGroupDao()
	err := userdao.JoinFriend(in.AskId, in.FriendId)
	if err != nil {
		response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", err.Error())
		return
	}
	response.RespSuccess(c, "ok", "request send ok!")
}
func NewGroup(c *gin.Context) {
	var in user.NewGroup
	if err := c.ShouldBindJSON(&in); err != nil {
		response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", "请输入正确的参数类型!")
		return
	}
	valid := validation.Validation{}
	isok_param, _ := valid.Valid(&in)
	if !isok_param {
		for _, err := range valid.Errors {
			response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", err.Message)
			return
		}
	}
	userdao := usersrv.NewUserGroupDao()
	err := userdao.NewGroup(in.CreateId, in.GroupName, in.UserIds...)
	if err != nil {
		response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", err.Error())
		return
	}
	response.RespSuccess(c, "ok", "request send ok!")
}
func GroupJoinFriend(c *gin.Context) {
	var in user.GroupJoinFriend
	if err := c.ShouldBindJSON(&in); err != nil {
		response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", "请输入正确的参数类型!")
		return
	}
	valid := validation.Validation{}
	isok_param, _ := valid.Valid(&in)
	if !isok_param {
		for _, err := range valid.Errors {
			response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", err.Message)
			return
		}
	}
	userdao := usersrv.NewUserGroupDao()
	err := userdao.GroupJoinFriend(in.CreatId, in.GroupId, in.UserIds...)
	if err != nil {
		response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", err.Error())
		return
	}
	response.RespSuccess(c, "ok", "request send ok!")
}

func GroupKickOutFriend(c *gin.Context) {
	var in user.GroupKickOutFriend
	if err := c.ShouldBindJSON(&in); err != nil {
		response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", "请输入正确的参数类型!")
		return
	}
	valid := validation.Validation{}
	isok_param, _ := valid.Valid(&in)
	if !isok_param {
		for _, err := range valid.Errors {
			response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", err.Message)
			return
		}
	}
	userdao := usersrv.NewUserGroupDao()
	err := userdao.GroupKickOutFriend(in.GroupId, in.UserIds...)
	if err != nil {
		response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", err.Error())
		return
	}
	response.RespSuccess(c, "ok", "request send ok!")
}
func GroupDelete(c *gin.Context) {
	var in user.GroupDelete
	if err := c.ShouldBindJSON(&in); err != nil {
		response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", "请输入正确的参数类型!")
		return
	}
	valid := validation.Validation{}
	isok_param, _ := valid.Valid(&in)
	if !isok_param {
		for _, err := range valid.Errors {
			response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", err.Message)
			return
		}
	}
	userdao := usersrv.NewUserGroupDao()
	err := userdao.DeleteGroup(in.GroupId)
	if err != nil {
		response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", err.Error())
		return
	}
	response.RespSuccess(c, "ok", "request send ok!")
}
func GroupMembers(c *gin.Context) {
	groupId, _ := strconv.ParseInt(c.Param("groupId"), 10, 64)
	userdao := usersrv.NewUserGroupDao()
	err, users := userdao.GroupMembers(groupId)
	if err != nil {
		response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", err.Error())
		return
	}
	response.RespSuccess(c, users, "ok")
}
func UserGroups(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Param("userId"), 10, 64)
	userdao := usersrv.NewUserGroupDao()
	err, users := userdao.UserGroups(userId)
	if err != nil {
		response.RespFailed(c, constex.ERROR_PARAM_ERROR, "", err.Error())
		return
	}
	response.RespSuccess(c, users, "ok")
}
