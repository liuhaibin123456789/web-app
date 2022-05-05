package mysql

import (
	"bluebell/models"
	"bluebell/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//var db *sqlx.DB

//// Init 初始化MySQL连接
//func Init(cfg *setting.MySQLConfig) (err error) {
//	// "user:password@tcp(host:port)/dbname"
//	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
//	db, err = sqlx.Connect("mysql", dsn)
//	if err != nil {
//		return
//	}
//	db.SetMaxOpenConns(cfg.MaxOpenConns)
//	db.SetMaxIdleConns(cfg.MaxIdleConns)
//	return
//}
//
//// Close 关闭MySQL连接
//func Close() {
//	_ = db.Close()
//}
// GDB mysql数据库的操作对象

var db *gorm.DB

func Init(cfg *setting.MySQLConfig) (err error) {
	//获取配置
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		return err
	}

	err = createTables()
	return err
}

func createTables() error {
	if !db.Migrator().HasTable(&models.UserTable{}) {
		err := db.AutoMigrate(&models.UserTable{})
		if err != nil {
			return err
		}
	}
	if !db.Migrator().HasTable(&models.PostTable{}) {
		err := db.AutoMigrate(&models.PostTable{})
		if err != nil {
			return err
		}
	}
	if !db.Migrator().HasTable(&models.CommunityTable{}) {
		err := db.AutoMigrate(&models.CommunityTable{})
		if err != nil {
			return err
		}
	}
	return nil
}
