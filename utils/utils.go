package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"go-web/config"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// jwt密钥
var jwtSecret = []byte(config.Cfg.JWT.Secret)

type CustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// 加密密码
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

// 校验检查密码
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// 生成 Token
func GenerateToken(userID uint) (string, error) {
	// 从配置文件中解析 Token 过期时间
	expireDuration, err := time.ParseDuration(config.Cfg.JWT.Expire)
	if err != nil {
		return "", err
	}
	// 构造自定义 Claims
	claims := &CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			// Token 过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
			// Token 签发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	// 使用 HS256 算法生成 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用密钥进行签名并返回 Token 字符串
	return token.SignedString(jwtSecret)
}

// 解析 Token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	claims := &CustomClaims{}

	// 解析 token 并验证签名
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		// 确保签名算法是 HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token 签名算法错误")
		}
		// 返回密钥给库，让它验证签名
		return jwtSecret, nil
	})
	// 检查解析是否出错
	if err != nil {
		return nil, err
	}

	// token 无效或已过期
	if !token.Valid || (claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now())) {
		return nil, errors.New("token 无效或已过期")
	}

	return claims, nil
}
