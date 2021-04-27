package user

import "time"

type TbUser struct {
	UserId    int64      `gorm:"primary_key;comment:'用户id';" json:"user_id"`
	CreatedAt time.Time  `gorm:"comment:'创建时间'" json:"-"`
	UpdatedAt time.Time  `gorm:"comment:'更新时间'" json:"-"`
	DeletedAt *time.Time `gorm:"comment:'删除时间'" sql:"index" json:"-"`
	UserName  string     `gorm:"comment:'用户名字';type:varchar(32);comment:'用户名字';not null;" json:"user_name"`
	UserMail  string     `gorm:"comment:'用户mail';type:varchar(32);comment:'用户邮件';not null;" json:"user_mail"`
	LoginHash string     `gorm:"column:login_hash;type:varchar(32);default:'';not null;comment:'密码hash'" json:"login_hash"`
	LoginSalt string     `gorm:"column:login_salt;type:varchar(16);default:'';not null;comment:'密码盐'" json:"login_salt"`
	Friend    []TbFriend `gorm:"ForeignKey:UserID;" json:"friend"`
	Device    []TbDevice `gorm:"ForeignKey:UserID;" json:"device"`
	Group     []TbGroup  `gorm:"many2many:tb_user_group;column:group;comment:'组'" json:"group"`
}
type TbFriend struct {
	FriendID  int64      `gorm:"primary_key;comment:'friend_id';" json:"friend_id"`
	CreatedAt time.Time  `gorm:"comment:'创建时间'" json:"-"`
	UpdatedAt time.Time  `gorm:"comment:'更新时间'" json:"-"`
	DeletedAt *time.Time `gorm:"comment:'删除时间'" sql:"index" json:"-"`
	OrderName string     `gorm:"comment:'订单名字';type:varchar(32);default:''; not null;"`
	UserID    int64      `gorm:"comment:'用户id'" json:"user_id"`
}
type TbDevice struct {
	DeviceID  int64      `gorm:"primary_key;comment:'device_id';" json:"device_id"`
	CreatedAt time.Time  `gorm:"comment:'创建时间'" json:"-"`
	UpdatedAt time.Time  `gorm:"comment:'更新时间'" json:"-"`
	DeletedAt *time.Time `gorm:"comment:'删除时间'" sql:"index" json:"-"`
	OrderName string     `gorm:"comment:'订单名字';type:varchar(32);default:''; not null;"`
	UserID    int64      `gorm:"comment:'用户id'" json:"user_id"`
}

type TbGroup struct {
	GroupId     int64      `gorm:"primary_key;comment:'group_id';" json:"group_id"`
	CreatedAt   time.Time  `gorm:"comment:'创建时间'" json:"-"`
	UpdatedAt   time.Time  `gorm:"comment:'更新时间'" json:"-"`
	DeletedAt   *time.Time `gorm:"comment:'删除时间'" sql:"index" json:"-"`
	GroupName   string     `gorm:"comment:'组名称';type:varchar(32);default:''; not null;"`
	CreatedId   int64      `gorm:"comment:'组创建者id';type:varchar(32);default:''; not null;"`
	CreatedName string     `gorm:"comment:'组创建者名称';type:varchar(32);default:''; not null;"`
	User        []TbUser   `gorm:"many2many:tb_group_user;column:user;comment:'用户'" json:"user"` //多对多
}
