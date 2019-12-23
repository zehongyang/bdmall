package controller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"time"
)

//图片上传
func UploadImage(c *gin.Context)  {
	file, err := c.FormFile("Filedata")
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"message":"图片上传错误"})
		return
	}
	hash := md5.New()
	hash.Write([]byte(fmt.Sprintf("%d",time.Now().UnixNano())))
	fileName := hex.EncodeToString(hash.Sum(nil)) + filepath.Ext(file.Filename)
	err = c.SaveUploadedFile(file, "./public/static/uploads/"+fileName)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"message":"图片上传错误"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"message":"成功","data":map[string]string{"path":"/public/uploads/"+fileName}})
}
