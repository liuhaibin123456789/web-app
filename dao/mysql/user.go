package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"encoding/hex"
	"go.uber.org/zap"
	"time"
)

// 把每一步数据库操作封装成函数
// 待logic层根据业务需求调用

const secret = "cold bin"

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	//sqlStr := `select count(user_id) from user where username = ?`
	//var count int64
	//if err := db.Get(&count, sqlStr, username); err != nil {
	//	return err
	//}
	//if count > 0 {
	//	return ErrorUserExist
	//}
	//return
	u := &models.UserTable{}
	err = db.Model(u).Where("user_name=?", username).Find(u).Error
	if err != nil {
		err = ErrorUserExist
	}
	return
}

// InsertUser 想数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	//// 执行SQL语句入库
	//sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	//_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	//return
	u := new(models.UserTable)
	u.UserId = user.UserID
	u.CreateTime = time.Now()
	u.UserName = user.Username
	u.Password = user.Password
	err = db.Model(&models.UserTable{}).Create(u).Error
	return
}

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	//oPassword := user.Password // 用户登录的密码
	//sqlStr := `select user_id, username, password from user where username=?`
	//err = db.Get(user, sqlStr, user.Username)
	//if err == sql.ErrNoRows {
	//	return ErrorUserNotExist
	//}
	//if err != nil {
	//	// 查询数据库失败
	//	return err
	//}
	//// 判断密码是否正确
	//password := encryptPassword(oPassword)
	//if password != user.Password {
	//	return ErrorInvalidPassword
	//}
	//return
	// 判断密码是否正确
	password := encryptPassword(user.Password)
	//查询密码
	var rPassword string
	db.Model(&models.UserTable{}).Select("password").Where("user_name=?", user.Username).Find(&rPassword)
	if password != rPassword {
		return ErrorInvalidPassword
	}
	zap.L().Debug("func Login(user *models.User) (err error) ", zap.String("right password", rPassword))
	return
}

// GetUserById 根据id获取用户信息
func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	//sqlStr := `select user_id, username from user where user_id = ?`
	//err = db.Get(user, sqlStr, uid)
	//return
	err = db.Model(&models.UserTable{}).Select("user_id", "user_name").Where("user_id=?", uid).Find(&user.UserID, &user.Username).Error
	return
}
