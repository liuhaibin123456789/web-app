package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"web_app/tool"

	"web_app/model"
)

// GDB mysql数据库的操作对象
var GDB *gorm.DB

func Mysql() (err error) {
	//获取配置
	user := tool.V.GetString("mysql.user")
	pwd := tool.V.GetString("mysql.password")
	h := tool.V.GetString("mysql.host")
	p := tool.V.GetString("mysql.port")
	db := tool.V.GetString("mysql.dbname")
	dsn := user + ":" + pwd + "@tcp(" + h + ":" + p + ")/" + db + "?charset=utf8mb4&parseTime=True&loc=Local"
	gdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {

		return err
	}
	GDB = gdb

	err = createTables()
	return err
}

func createTables() error {
	if !GDB.Migrator().HasTable(&model.User{}) {
		err := GDB.AutoMigrate(&model.User{})
		if err != nil {
			return err
		}
	}
	if !GDB.Migrator().HasTable(&model.Community{}) {
		err := GDB.AutoMigrate(&model.Community{})
		if err != nil {
			return err
		}
	}
	if !GDB.Migrator().HasTable(&model.Post{}) {
		err := GDB.AutoMigrate(&model.Post{})
		if err != nil {
			return err
		}
	}
	return nil
}
