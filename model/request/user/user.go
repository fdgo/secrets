package user

import (
	"secret/utils/validation"
)

type Regist struct {
	Mail   string `valid:"Required;" json:"mail"`
	Pwd    string `valid:"Required;" json:"pwd"`
	Name   string
	Mobile string
	VfCode string
}

func (in *Regist) Valid(v *validation.Validation) {
	//if !model.VerifyMobile(in.Mobile) {
	//	v.SetError("手机号", "请输入正确的手机号!")
	//}
	//if err := model.VerifyPwd(in.Pwd); err != nil {
	//	v.SetError("密码", err.Error())
	//}
	//if err := model.VerifyVfcode(in.VfCode); err != nil {
	//	v.SetError("验证码", err.Error())
	//}
}

type Login struct {
	Mail   string `valid:"Required;" json:"mail"`
	Device int64  `valid:"Required;" json:"device"`
	Mobile string
	VfCode string
	Pwd    string `valid:"Required;" json:"pwd"`
}

func (in *Login) Valid(v *validation.Validation) {
	//if !model.VerifyMobile(in.Mobile) {
	//	v.SetError("手机号", "请输入正确的手机号!")
	//}
	//if err := model.VerifyPwd(in.Pwd); err != nil {
	//	v.SetError("密码", err.Error())
	//}
	//if err := model.VerifyVfcode(in.VfCode); err != nil {
	//	v.SetError("验证码", err.Error())
	//}
}

type JoinFriend struct {
	AskId    int64 `valid:"Required;" json:"askId"`
	FriendId int64 `valid:"Required;" json:"friendId"`
}

func (in *JoinFriend) Valid(v *validation.Validation) {
	//if !model.VerifyMobile(in.Mobile) {
	//	v.SetError("手机号", "请输入正确的手机号!")
	//}
	//if err := model.VerifyPwd(in.Pwd); err != nil {
	//	v.SetError("密码", err.Error())
	//}
	//if err := model.VerifyVfcode(in.VfCode); err != nil {
	//	v.SetError("验证码", err.Error())
	//}
}

type DeviceBind struct {
	UserId   int64 `valid:"Required;" json:"userId"`
	DeviceId int64 `valid:"Required;" json:"deviceId"`
}

func (in *DeviceBind) Valid(v *validation.Validation) {
	//if !model.VerifyMobile(in.Mobile) {
	//	v.SetError("手机号", "请输入正确的手机号!")
	//}
	//if err := model.VerifyPwd(in.Pwd); err != nil {
	//	v.SetError("密码", err.Error())
	//}
	//if err := model.VerifyVfcode(in.VfCode); err != nil {
	//	v.SetError("验证码", err.Error())
	//}
}

type DeleteDevice struct {
	UserId   int64 `valid:"Required;" json:"userId"`
	DeviceId int64 `valid:"Required;" json:"deviceId"`
}

func (in *DeleteDevice) Valid(v *validation.Validation) {
	//if !model.VerifyMobile(in.Mobile) {
	//	v.SetError("手机号", "请输入正确的手机号!")
	//}
	//if err := model.VerifyPwd(in.Pwd); err != nil {
	//	v.SetError("密码", err.Error())
	//}
	//if err := model.VerifyVfcode(in.VfCode); err != nil {
	//	v.SetError("验证码", err.Error())
	//}
}

type NewGroup struct {
	CreateId  int64         `valid:"Required;" json:"createId"`
	UserIds   []interface{} `valid:"Required;" json:"userIds"`
	GroupName string        `valid:"Required;" json:"groupName"`
}

func (in *NewGroup) Valid(v *validation.Validation) {
	//if !model.VerifyMobile(in.Mobile) {
	//	v.SetError("手机号", "请输入正确的手机号!")
	//}
	//if err := model.VerifyPwd(in.Pwd); err != nil {
	//	v.SetError("密码", err.Error())
	//}
	//if err := model.VerifyVfcode(in.VfCode); err != nil {
	//	v.SetError("验证码", err.Error())
	//}
}

type GroupJoinFriend struct {
	CreatId int64         `valid:"Required;" json:"createId"`
	GroupId int64         `valid:"Required;" json:"groupId"`
	UserIds []interface{} `valid:"Required;" json:"userIds"`
}

func (in *GroupJoinFriend) Valid(v *validation.Validation) {
	//if !model.VerifyMobile(in.Mobile) {
	//	v.SetError("手机号", "请输入正确的手机号!")
	//}
	//if err := model.VerifyPwd(in.Pwd); err != nil {
	//	v.SetError("密码", err.Error())
	//}
	//if err := model.VerifyVfcode(in.VfCode); err != nil {
	//	v.SetError("验证码", err.Error())
	//}
}

type GroupKickOutFriend struct {
	GroupId int64         `valid:"Required;" json:"groupId"`
	UserIds []interface{} `valid:"Required;" json:"userIds"`
}

func (in *GroupKickOutFriend) Valid(v *validation.Validation) {
	//if !model.VerifyMobile(in.Mobile) {
	//	v.SetError("手机号", "请输入正确的手机号!")
	//}
	//if err := model.VerifyPwd(in.Pwd); err != nil {
	//	v.SetError("密码", err.Error())
	//}
	//if err := model.VerifyVfcode(in.VfCode); err != nil {
	//	v.SetError("验证码", err.Error())
	//}
}

type GroupDelete struct {
	UserId    int64 `valid:"Required;" json:"userId"`
	GroupName string
	GroupId   int64 `valid:"Required;" json:"groupId"`
}

func (in *GroupDelete) Valid(v *validation.Validation) {
	//if !model.VerifyMobile(in.Mobile) {
	//	v.SetError("手机号", "请输入正确的手机号!")
	//}
	//if err := model.VerifyPwd(in.Pwd); err != nil {
	//	v.SetError("密码", err.Error())
	//}
	//if err := model.VerifyVfcode(in.VfCode); err != nil {
	//	v.SetError("验证码", err.Error())
	//}
}
