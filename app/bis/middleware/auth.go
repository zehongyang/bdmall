package middleware

import (
	"bdmall/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

//登陆验证
func LoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := utils.SessionStore.Get(c.Request, "user")
		//忽略验证路径
		ignorePath := []string{"/bis/login","/bis/register","/bis/add"}
		var flag bool
		for _, v := range ignorePath {
			if c.FullPath() == v {
				flag = true
				break
			}
		}
		if !flag && session.Values["username"] == nil {
			c.Redirect(http.StatusFound,"/bis/login")
			return
		}
		c.Next()
	}
}
