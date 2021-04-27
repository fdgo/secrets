package middle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
	"secret/initial"
	"secret/model/response"
	"secret/utils/auth"
	"secret/utils/constex"
	"secret/utils/limit"
	"secret/utils/timex"
	"strings"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("interview/access-Control-Allow-Origin", "*")                            //跨域
		ctx.Header("interview/access-Control-Allow-Headers", "Token,Content-Type")          //必须的请求头
		ctx.Header("interview/access-Control-Allow-Methods", "OPTIONS,PUT,POST,GET,DELETE") //接收的请求方法
	}
}
func NoRoute() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == "OPTIONS" {
			ctx.JSON(200, nil)
		}
	}
}
func Recover(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				DebugStack := ""
				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					DebugStack += v + "<br/>"
				}
				c.JSON(http.StatusOK, struct {
					Code   int
					Name   string
					Time   string
					Host   string
					Uri    string
					Method string
					Agent  string
					Stack  string
				}{
					Code:   500,
					Name:   name,
					Time:   timex.TimeCurrString(),
					Host:   c.Request.Host,
					Uri:    c.Request.RequestURI,
					Method: c.Request.Method,
					Agent:  c.Request.UserAgent(),
					Stack:  DebugStack})
			}
		}()
		c.Next()
	}
}
// 中间件，用令牌桶限制请求频率
func Limit(c *gin.Context) {
	limit.LimitHandler(c)
}
func Auth(token *auth.JwtToken) gin.HandlerFunc {
	return func(c *gin.Context) {
		istokenok, msg, sub, code := token.Decode(c.Request, c.Writer)
		if !istokenok {
			response.RespFailed(c, code, nil, "请先登录!"+msg)
			c.Abort()
			return
		}
		result := initial.RedisClient().Get(constex.REDIS_USER_TOKEN + sub.Id)
		if result == nil || result.Val() == "" {
			response.RespFailed(c, constex.ERROR_TOKEN_NEED, nil, "您的token异常!")
			c.Abort()
			return
		}
		c.Request.Header.Set("city", c.Request.Header.Get("city"))
		c.Request.Header.Set("lat", c.Request.Header.Get("lat"))
		c.Request.Header.Set("lgt", c.Request.Header.Get("lgt"))
		c.Request.Header.Set("X-Head-Uuid", sub.Id)
		c.Request.Header.Set("X-Head-TimeStamp", timex.TimeInt64ToTimeString(sub.ExpiresAt))
		c.Next()
	}
}
//{"city":"深圳市","lat":30.664649,"lgt":104.570084}
