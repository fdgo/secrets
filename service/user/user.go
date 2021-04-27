package user

import (
	"errors"
	"github.com/jinzhu/gorm"
	"secret/initial"
	mysqluser "secret/model/mysql/user"
	"secret/model/response/user"
	"secret/utils/auth"
	"secret/utils/rand"
	"secret/utils/sign/md5"
	"secret/utils/snakesnow"
)

type UserGroupDao struct {
	db *gorm.DB
}

func NewUserGroupDao() *UserGroupDao {
	return &UserGroupDao{
		db: initial.MysqlClient(),
	}
}

//创建用户
func (ugdao *UserGroupDao) Regist(mail, pwd string) (error, int64, int64) {
	var usersrc mysqluser.TbUser
	usersrc.UserMail = mail
	tx := ugdao.db.Begin()
	err := tx.Where("user_mail=?", mail).Find(&usersrc).Error
	if err == gorm.ErrRecordNotFound {
		UserId, err := snakesnow.RandUserId()
		if err != nil {
			tx.Rollback()
			return errors.New("生成用户id失败!"), 0, 0
		}
		DevId, err := snakesnow.RandDeviceId()
		if err != nil {
			tx.Rollback()
			return errors.New("生成设备id失败!"), 0, 0
		}
		var user mysqluser.TbUser
		user.UserId = UserId
		user.UserMail = mail
		user.LoginSalt = rand.GetRandomString(6)
		user.LoginHash = md5.HashForPwd(user.LoginSalt, pwd)
		if err := tx.Save(&user).Error; err != nil {
			tx.Rollback()
			return err, 0, 0
		}
		err = tx.Model(&user).Association("Device").Append(&mysqluser.TbDevice{DeviceID: DevId}).Error
		if err != nil {
			tx.Rollback()
			return err, 0, 0
		}
		err = tx.Commit().Error
		if err != nil {
			tx.Rollback()
			return err, 0, 0
		}
		return nil, UserId, DevId
	} else {
		if err != nil {
			tx.Rollback()
			return errors.New("server is busy!"), 0, 0
		}
		tx.Rollback()
		return errors.New("user is exist!"), 0, 0
	}
}
func (ugdao *UserGroupDao) Login(mail string, pwd string, deviceId int64) (error, user.Login) {
	var useradv mysqluser.TbUser
	tx := ugdao.db.Begin()
	err := tx.Where("user_mail=?", mail).Find(&useradv).Error
	if err == gorm.ErrRecordNotFound { //邮件正确
		tx.Rollback()
		return errors.New("user is not exist!"), user.Login{}
	} else { //邮件错误
		if err != nil {
			tx.Rollback()
			return errors.New("server is busy!"), user.Login{}
		}
		if md5.HashForPwd(useradv.LoginSalt, pwd) != useradv.LoginHash {
			tx.Rollback()
			return errors.New("passwd is wrong !"), user.Login{}
		}
		if deviceId == -1 { //没有设备号
			var device []mysqluser.TbDevice
			err := tx.Where("user_id=?", useradv.UserId).Find(&device).Error
			if err != nil {
				tx.Rollback()
				return err, user.Login{}
			}
			if len(device) > 2 {
				tx.Rollback()
				return errors.New("该账户已经绑定三个设备!"), user.Login{}
			}
			DevId, err := snakesnow.RandDeviceId()
			if err != nil {
				tx.Rollback()
				return errors.New("生成设备id失败!!"), user.Login{}
			}
			err = tx.Model(&useradv).Association("Device").Append(&mysqluser.TbDevice{DeviceID: DevId}).Error
			if err != nil {
				tx.Rollback()
				return errors.New("绑定设备id失败!!"), user.Login{}
			}
			err = tx.Commit().Error
			if err != nil {
				tx.Rollback()
				return errors.New("server is busy!"), user.Login{}
			}
			return nil, user.Login{Token: auth.GenToken(useradv.UserId, DevId), DeviceId: DevId, UseradvId: useradv.UserId}
		} else { //有设备号
			var device mysqluser.TbDevice
			err := tx.Where("device_id=? and user_id=?", deviceId, useradv.UserId).Find(&device).Error
			if err == gorm.ErrRecordNotFound {
				tx.Rollback()
				return errors.New("device  is not exist!"), user.Login{}
			} else {
				if err != nil {
					tx.Rollback()
					return errors.New("server is busy!"), user.Login{}
				}
				err = tx.Commit().Error
				if err != nil {
					tx.Rollback()
					return errors.New("server is busy!"), user.Login{}
				}
				return nil, user.Login{Token: auth.GenToken(useradv.UserId, deviceId), DeviceId: deviceId, UseradvId: useradv.UserId}
			}
		}
	}
}

//创建群的同时，把用户拉进去
func (ugdao *UserGroupDao) NewGroup(creatid int64, groupname string, userids ...interface{}) error {
	GroupId, err := snakesnow.RandGroupId()
	if err != nil {
		return errors.New("创建群id失败!")
	}
	var sql string
	size := len(userids)
	for index := 0; index < size; index++ {
		if index != size-1 {
			sql += "user_id=? or "
		} else {
			sql += "user_id=?"
		}
	}
	tx := ugdao.db.Begin()
	var user []mysqluser.TbUser
	err = tx.Where(sql, userids...).Find(&user).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	var creatname string
	for _, u := range user {
		if u.UserId == creatid {
			creatname = u.UserName
		}
		err = tx.Model(&u).Association("Group").Append(mysqluser.TbGroup{GroupId: GroupId, GroupName: groupname, CreatedId: creatid, CreatedName: creatname}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	group := mysqluser.TbGroup{
		GroupId:     GroupId,
		GroupName:   groupname,
		CreatedId:   creatid,
		CreatedName: creatname,
		User:        user,
	}
	err = tx.Save(&group).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

//拉人进群
func (ugdao *UserGroupDao) GroupJoinFriend(creatid, groupid int64, users ...interface{}) error {
	var sql string
	size := len(users)
	for index := 0; index < size; index++ {
		if index != size-1 {
			sql += "user_id=? or "
		} else {
			sql += "user_id=?"
		}
	}
	tx := ugdao.db.Begin()
	var useradvs []mysqluser.TbUser
	err := tx.Where(sql, users...).Find(&useradvs).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	var group mysqluser.TbGroup
	err = tx.Where("group_id=?", groupid).Find(&group).Error
	if err == gorm.ErrRecordNotFound {
		tx.Rollback()
		return err
	} else {
		if err != nil {
			tx.Rollback()
			return err
		}
		var creatname string
		for _, u := range useradvs {
			if u.UserId == creatid {
				creatname = u.UserName
			}
			err = tx.Model(&u).Association("Group").Append(mysqluser.TbGroup{GroupId: group.GroupId, GroupName: group.GroupName, CreatedId: creatid, CreatedName: creatname}).Error
			if err != nil {
				tx.Rollback()
				return err
			}
			err = tx.Model(&group).Association("User").Append(mysqluser.TbUser{UserId: u.UserId, UserName: u.UserName}).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		err = tx.Commit().Error
		if err != nil {
			tx.Rollback()
			return err
		}
		return nil
	}
}

//获取所有设备号
func (ugdao *UserGroupDao) Devices(userId int64) (error, []int64) {
	var device []mysqluser.TbDevice
	err := ugdao.db.Where("user_id=?", userId).Find(&device).Error
	if err != nil {
		return err, []int64{}
	}
	var ret []int64
	for _, ele := range device {
		ret = append(ret, ele.DeviceID)
	}
	return nil, ret
}

//删除设备号
func (ugdao *UserGroupDao) DeleteDevice(userId, deviceId int64) error {
	err := ugdao.db.Where("device_id=? and user_id=? ", deviceId, userId).Unscoped().Delete(&mysqluser.TbDevice{}).Error
	if err != nil {
		return err
	}
	return nil
}

//换绑设备号
func (ugdao *UserGroupDao) BindDevice(userId, deviceId int64) (error, user.Login) {
	var asker mysqluser.TbUser
	asker.UserId = userId
	err := ugdao.db.Model(&asker).Association("Device").Append(&mysqluser.TbDevice{DeviceID: deviceId}).Error
	if err != nil {
		return err, user.Login{}
	}
	return nil, user.Login{Token: auth.GenToken(userId, deviceId), DeviceId: deviceId, UseradvId: userId}
}

//添加好友
func (ugdao *UserGroupDao) JoinFriend(askerId, friendId int64) error {
	var asker mysqluser.TbUser
	asker.UserId = askerId
	err := ugdao.db.Model(&asker).Association("Friend").Append(&mysqluser.TbFriend{FriendID: friendId}).Error
	if err != nil {
		return err
	}
	return nil
}

//剔除出群
func (ugdao *UserGroupDao) GroupKickOutFriend(groupid int64, users ...interface{}) error {
	var sql string
	size := len(users)
	for index := 0; index < size; index++ {
		if index != size-1 {
			sql += "user_id=? or "
		} else {
			sql += "user_id=?"
		}
	}
	tx := ugdao.db.Begin()
	var useradvs []mysqluser.TbUser
	err := tx.Where(sql, users...).Find(&useradvs).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	var group mysqluser.TbGroup
	err = tx.Where("group_id=?", groupid).Find(&group).Error
	if err == gorm.ErrRecordNotFound {
		tx.Rollback()
		return err
	} else {
		if err != nil {
			tx.Rollback()
			return err
		}
		for _, u := range useradvs {
			err = tx.Model(&u).Association("Group").Delete(mysqluser.TbGroup{GroupId: group.GroupId}).Error
			if err != nil {
				tx.Rollback()
				return err
			}
			err = tx.Model(&group).Association("User").Delete(mysqluser.TbUser{UserId: u.UserId}).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		err = tx.Commit().Error
		if err != nil {
			tx.Rollback()
			return err
		}
		return nil
	}
}

//删除群
func (ugdao *UserGroupDao) DeleteGroup(groupid int64) error {
	tx := ugdao.db.Begin()
	var group mysqluser.TbGroup
	err := tx.Where("group_id=?", groupid).Find(&group).Error
	if err == gorm.ErrRecordNotFound {
		tx.Rollback()
		return err
	} else {
		if err != nil {
			tx.Rollback()
			return err
		}
		var useradvs []mysqluser.TbUser
		err = tx.Model(&group).Association("User").Find(&useradvs).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		for _, u := range useradvs {
			err = tx.Model(&u).Association("Group").Delete(mysqluser.TbGroup{GroupId: group.GroupId}).Error
			if err != nil {
				tx.Rollback()
				return err
			}
			err = tx.Model(&group).Association("User").Delete(mysqluser.TbUser{UserId: u.UserId}).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		err = tx.Where("group_id=?", groupid).Delete(&mysqluser.TbGroup{}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		err = tx.Commit().Error
		if err != nil {
			tx.Rollback()
			return err
		}
		return nil
	}
}

//查看某人有多少群
func (ugdao *UserGroupDao) UserGroups(userid int64) (error, []mysqluser.TbGroup) {
	tx := ugdao.db.Begin()
	var useradv mysqluser.TbUser
	err := tx.Where("user_id=?", userid).Find(&useradv).Error
	if err == gorm.ErrRecordNotFound {
		tx.Rollback()
		return err, []mysqluser.TbGroup{}
	} else {
		if err != nil {
			tx.Rollback()
			return err, []mysqluser.TbGroup{}
		}
		var groups []mysqluser.TbGroup
		err = tx.Model(&useradv).Association("Group").Find(&groups).Error
		if err != nil {
			tx.Rollback()
			return err, []mysqluser.TbGroup{}
		}
		err = tx.Commit().Error
		if err != nil {
			tx.Rollback()
			return err, []mysqluser.TbGroup{}
		}
		return nil, groups
	}
}

//查看群里成员
func (ugdao *UserGroupDao) GroupMembers(groupId int64) (error, []mysqluser.TbUser) {
	var group mysqluser.TbGroup
	tx := ugdao.db.Begin()
	err := tx.Where("group_id=?", groupId).Find(&group).Error
	if err == gorm.ErrRecordNotFound {
		tx.Rollback()
		return err, []mysqluser.TbUser{}
	} else {
		if err != nil {
			tx.Rollback()
			return err, []mysqluser.TbUser{}
		}
		var users []mysqluser.TbUser
		err := ugdao.db.Model(&group).Association("User").Find(&users).Error
		if err != nil {
			tx.Rollback()
			return err, []mysqluser.TbUser{}
		}
		err = tx.Commit().Error
		if err != nil {
			tx.Rollback()
			return err, []mysqluser.TbUser{}
		}
		return nil, users
	}
}
