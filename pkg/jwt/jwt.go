package jwt

import (
	"bluebell/dao/redis"
	"errors"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	MySecret          = []byte("123456")
	ErrorWrongToken   = errors.New("token格式错误")
	ErrorExpiredToken = errors.New("token过期")
)

//由于前端没有适配改进过后的双token策略，暂时不使用双token AccessTokenExpireDuration的时间延长作为普通token
const (
	AccessTokenExpireDuration  = time.Hour * 24 * 7
	RefreshTokenExpireDuration = time.Hour * 24 * 30
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(userID int64) (aToken string, err error) {
	// 创建一个我们自己的声明的数据
	c := MyClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(
				time.Duration(viper.GetInt("auth.jwt_expire")) * time.Hour).Unix(), // 过期时间
			Issuer: "cold bin", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	aToken, err = token.SignedString(MySecret)
	//将accessToken放到redis里，对应user_id
	err = redis.Set(strconv.FormatInt(userID, 10), aToken, AccessTokenExpireDuration)
	return
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, ErrorWrongToken
}

// GenTwoToken 生成JWT
func GenTwoToken(userID int64) (aToken, rToken string, err error) {
	//生成accessToken，时间短，保存有用信息
	claim := MyClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessTokenExpireDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "cold bin",
		},
	}
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(MySecret)
	if err != nil {
		return "", "", err
	}
	//生成refreshToken，只需要部门官方字段即可，时间长
	rToken, err = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.StandardClaims{ExpiresAt: time.Now().Add(RefreshTokenExpireDuration).Unix(), Issuer: "cold bin"}).SignedString(MySecret)
	if err != nil {
		return "", "", err
	}
	//将accessToken放到redis里，对应user_id
	err = redis.Set(strconv.FormatInt(userID, 10), aToken, AccessTokenExpireDuration)
	if err != nil {
		return "", "", err
	}
	return aToken, rToken, nil
}

// ParseTokenV2 解析JWT
func ParseTokenV2(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		if v.Errors == jwt.ValidationErrorExpired {
			return &MyClaims{}, ErrorExpiredToken
		}
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, ErrorWrongToken
}

// RefreshToken 适用于access_token过期，需要获取新的access_token
func RefreshToken(userId int64, auth string) (newAToken, newRToken string, err error) {
	//拆解auth，拿到refresh_token
	rToken := strings.TrimPrefix(auth, "Bearer ")
	//refresh_token无效直接返回
	if _, err1 := jwt.Parse(rToken, func(t *jwt.Token) (interface{}, error) {
		return MySecret, nil
	}); err1 != nil {
		err = err1
		return
	}
	return GenTwoToken(userId)
}
