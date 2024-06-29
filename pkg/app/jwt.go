package app

import (
	"blog-service/global"
	"blog-service/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	// 自定义认证信息
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	// JWT 库中预定义的， 算 JWT 规范
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

// GenerateToken 根据传入的 AppKey AppSecret 以及项目配置中设置的签发者与过期时间生成指定 token
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	// 根据 Claims 结构体创建 token 实例。
	// 第一个是加密算法 第二个是用户预定义的一些权利要求
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串 根据传入的 secret 参数不同， 签名并返回标准 token
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	// 具体的解码和校验过程
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		// Valid: 验证基于时间的声明
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, nil
}
