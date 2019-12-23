package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

var Store = base64Captcha.DefaultMemStore

//生成验证码
func GeneratorCaptcha(c *gin.Context)  {
	captcha := base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, Store)
	id, b64s, err := captcha.Generate()
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"code":500})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"msg":"success","data":gin.H{"id":id,"b64s":b64s}})
}
