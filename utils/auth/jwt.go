package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"secret/initial"
	"secret/utils/config"
	"secret/utils/constex"
	"secret/utils/stringex"
	"secret/utils/timex"
	"strconv"
	"strings"
	"time"
)

var (
	TokenExpired     error = errors.New("Token过期!")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("token错误!")
	TokenInvalid     error = errors.New("token错误!")
)

type JwtToken struct {
	SigningKey []byte
}

func (j *JwtToken) CreateToken(claims jwt.StandardClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}
func (j *JwtToken) ParseToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}
func (j *JwtToken) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

//Decode 解码
func (j *JwtToken) Decode(r *http.Request, w http.ResponseWriter) (bool, string, *jwt.StandardClaims, int) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return false, "请求未携带token，无权限访问", nil, constex.ERROR_TOKEN_NEED
	}
	if !strings.HasPrefix(token, "Bearer ") {
		return false, "请求未携带token，无权限访问", nil, constex.ERROR_TOKEN_NEED
	}
	if stringex.Length(token) < 128 {
		return false, "请求未携带超过128位的token参数，无权限访问", nil, constex.ERROR_TOKEN_ERROR
	}
	token = stringex.SubString(token, stringex.Length("Bearer "), stringex.Length(token)-stringex.Length("Bearer "))
	sub, err := j.ParseToken(token)
	if sub == nil {
		return false, err.Error(), nil, constex.ERROR_TOKEN_ERROR
	}
	return true, "token验证成功!", sub, 200
}
func GenToken(user_id , device int64) string {
	extime := config.GloCfg.GetInt64("jwt.exptime")
	claims := jwt.StandardClaims{Id: fmt.Sprintf("%d", user_id), Audience: fmt.Sprintf("%d", device), ExpiresAt: extime + timex.TimeCurrInt64()}
	jwtoken := &JwtToken{SigningKey: []byte(config.GloCfg.GetString("jwt.secretKey"))}
	token, _ := jwtoken.CreateToken(claims)
	exptime, _ := time.ParseDuration("+" + strconv.FormatInt(extime, 10) + "s")
	initial.RedisClient().Set(constex.REDIS_USER_TOKEN+fmt.Sprintf("%d", user_id), token, exptime).Err()
	return token
}
