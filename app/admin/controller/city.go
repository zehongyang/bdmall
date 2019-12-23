package controller

import (
	"bdmall/app/common/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//获取城市列表
func GetCities(c *gin.Context)  {
	pid, err := strconv.Atoi(c.PostForm("city_id"))
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"message":"参数错误"})
		return
	}
	err, cities := model.GetCities(int32(pid))
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"code":500,"message":"服务端错误"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"message":"成功","data":cities})
}
