package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strings"
	"tests/lib"
)

type CodeInfo struct {
	Code string
}

// UserCode 模拟验证码验证信息 map{func:{username:{code}}}
var UserCode = map[string]map[string]*CodeInfo{
	"Login": {
		"admin": {
			Code: "123456",
		},
		"test": {
			Code: "123456",
		},
	},
}

func Code() gin.HandlerFunc {
	return func(c *gin.Context) {
		input := struct {
			Username string `json:"username"`
			Code     string `json:"code"`
		}{}
		if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			c.Abort()
			return
		}
		if err := lib.VerifyMulti(input, lib.UserVerify, lib.CodeVerify); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		if info, ok := UserCode[funcName(c)][input.Username]; !ok || info.Code != input.Code {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Code"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func funcName(c *gin.Context) string {
	name := strings.Split(c.HandlerName(), ".")[2]
	return strings.Split(name, "-")[0]
}
