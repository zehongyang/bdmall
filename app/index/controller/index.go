package controller

import (
	"bdmall/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context)  {
	session, _ := utils.SessionStore.Get(c.Request, "home_user")
	username := session.Values["username"]
	id := session.Values["id"]
	var isLogin bool
	if username != nil {
		isLogin = true
	}
	c.HTML(http.StatusOK,"home/index/index.html",gin.H{"username":username,"id":id,"isLogin":isLogin})
}
