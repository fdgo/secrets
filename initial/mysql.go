package initial

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"secret/model/mysql/user"
	"sync"
	"time"
)

var mysqlLock sync.Mutex
var mysqlinstance *gorm.DB

func InitMysql() *gorm.DB {
	db, err := gorm.Open("mysql", "root:000000@tcp(192.168.204.139:3306)/secret?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return nil
	}
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(1000)
	db.LogMode(true)
	db.SingularTable(true)
	db.DB().SetConnMaxLifetime(time.Duration(3000) * time.Second)
	{
		//db.DropTableIfExists(&model.TbUseradv{},  &model.TbOrder{})
		db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='用户表'").AutoMigrate(&user.TbUser{})
		db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='爱好表'").AutoMigrate(&user.TbGroup{})
		db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='朋友表'").AutoMigrate(&user.TbFriend{})
		db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='设备表'").AutoMigrate(&user.TbDevice{})
	}
	mysqlinstance = db
	return mysqlinstance
}

// 得到唯一的主库实例
func MysqlClient() *gorm.DB {
	if mysqlinstance != nil {
		return mysqlinstance
	}
	mysqlLock.Lock()
	defer mysqlLock.Unlock()
	if mysqlinstance != nil {
		return mysqlinstance
	}
	return InitMysql()
}
