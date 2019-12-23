package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//后台系统首页
func Index(c *gin.Context)  {
	c.HTML(http.StatusOK,"admin/index/index.html",gin.H{})
}

//系统欢迎页
func Welcome(c *gin.Context)  {
	c.HTML(http.StatusOK,"admin/index/welcome.html",gin.H{})
}