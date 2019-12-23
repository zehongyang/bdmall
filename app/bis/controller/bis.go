package controller

import (
	"bdmall/app/common/model"
	"bdmall/app/conf"
	"bdmall/app/queue"
	"bdmall/utils"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//商户登陆页面展示
func BisLoginIndex(c *gin.Context)  {
	c.HTML(http.StatusOK,"bis/bis/login.html",gin.H{})
}

//商户注册页面
func BisRegister(c *gin.Context)  {
	//获取城市列表
	_, cities := model.GetCities(0)
	//获取分类列表
	_, categories := model.GetCategories(0)
	c.HTML(http.StatusOK,"bis/bis/register.html",gin.H{"cities":cities,"categories":categories})
}

//添加商铺
func BisAdd(c *gin.Context)  {
	var bis model.Bis
	err := c.ShouldBind(&bis)
	if err != nil {
		utils.ResErrors(c,gin.H{"code":0,"msg":err,"url":"/bis/register","wait":3})
		return
	}
	var bisAccount model.BisAccount
	err = c.ShouldBind(&bisAccount)
	if err != nil {
		utils.ResErrors(c,gin.H{"code":0,"msg":err,"url":"/bis/register","wait":3})
		return
	}
	var bisLocation model.BisLocation
	err = c.ShouldBind(&bisLocation)
	if err != nil {
		utils.ResErrors(c,gin.H{"code":0,"msg":err,"url":"/bis/register","wait":3})
		return
	}
	bis.CityPath = strconv.Itoa(int(bis.CityId))+","+c.PostForm("se_city_id")
	bisLocation.CategoryPath = strings.Join(c.PostFormArray("cates"),",")
	tx := conf.DB.Begin()
	defer func() {
		if r := recover();r != nil {
			tx.Rollback()
		}
	}()
	if err = tx.Create(&bis).Error;err != nil{
		tx.Rollback()
		utils.ResErrors(c,gin.H{"code":0,"msg":err,"url":"/bis/register","wait":3})
		return
	}
	bisAccount.Code = utils.GenerateString(10)
	bisAccount.BisId = bis.ID
	bisAccount.Password = fmt.Sprintf("%x",md5.Sum([]byte(fmt.Sprintf("%x%x", md5.Sum([]byte(bisAccount.Password)),
		md5.Sum([]byte(bisAccount.Code))))))
	if err = tx.Create(&bisAccount).Error;err != nil {
		tx.Rollback()
		utils.ResErrors(c,gin.H{"code":0,"msg":err,"url":"/bis/register","wait":3})
		return
	}
	longitude, latitude, ok := utils.GetXyByAddress(bisLocation.Address)
	if !ok {
		utils.ResErrors(c,gin.H{"code":0,"msg":"地址不合法","url":"/bis/register","wait":3})
		return
	}
	bisLocation.Xpoint = longitude
	bisLocation.Ypoint = latitude
	bisLocation.BisId = bis.ID
	bisLocation.CityPath = strings.Join([]string{c.PostForm("city_id"),c.PostForm("se_city_id")},",")
	if err = tx.Create(&bisLocation).Error;err != nil {
		tx.Rollback()
		utils.ResErrors(c,gin.H{"code":0,"msg":err,"url":"/bis/register","wait":3})
		return
	}
	if err = tx.Commit().Error;err != nil {
		utils.ResErrors(c,gin.H{"code":0,"msg":err,"url":"/bis/register","wait":3})
		return
	}
	mp := make(map[string]string)
	mp["to"] = bis.Email
	mp["subject"] = "商户入驻申请"
	mp["content"] = fmt.Sprintf("<a href='/bis/status?id=%d'>查看状态</a>",bis.ID)
	queue.EmailChan <- mp
	utils.ResSuccess(c,gin.H{"code":1,"msg":"商户申请成功","url":"/bis/login","wait":3})
	return
}

//商户登陆
func BisLogin(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	var bisAccount model.BisAccount
	err := conf.DB.Where("username = ?", username).First(&bisAccount).Error
	if err != nil {
		utils.ResErrors(c,gin.H{"code":0,"msg":err,"url":"/bis/login","wait":3})
		return
	}
	pwd := fmt.Sprintf("%x",md5.Sum([]byte(fmt.Sprintf("%x%x",md5.Sum([]byte(password)),md5.Sum([]byte(bisAccount.Code))))))
	if pwd != bisAccount.Password {
		utils.ResErrors(c,gin.H{"code":0,"msg":"用户名或密码错误","url":"/bis/login","wait":3})
		return
	}
	if bisAccount.Status != 1 {
		utils.ResErrors(c,gin.H{"code":0,"msg":"用户状态补正常","url":"/bis/login","wait":3})
		return
	}
	bisAccount.LastLoginIp = c.Request.RemoteAddr
	bisAccount.LastLoginTime = time.Now()
	if err = conf.DB.Save(&bisAccount).Error;err != nil {
		utils.ResErrors(c,gin.H{"code":0,"msg":err,"url":"/bis/login","wait":3})
		return
	}
	session, _ := utils.SessionStore.Get(c.Request, "user")
	session.Values["username"] = username
	session.Values["id"] = bisAccount.ID
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		utils.ResErrors(c,gin.H{"code":0,"msg":err,"url":"/bis/login","wait":3})
		return
	}
	utils.ResSuccess(c,gin.H{"code":1,"msg":"登陆成功","url":"/bis/index","wait":3})
	return
}

//商户后台首页
func BisIndex(c *gin.Context)  {
	session, _ := utils.SessionStore.Get(c.Request, "user")
	c.HTML(http.StatusOK,"bis/bis/index.html",gin.H{"username":session.Values["username"]})
}

//退出登陆
func BisLogout(c *gin.Context)  {
	session, _ := utils.SessionStore.Get(c.Request, "user")
	session.Values["username"] = nil
	session.Values["id"] = nil
	err := session.Save(c.Request, c.Writer)
	if err != nil {
		utils.ResErrors(c,gin.H{"code":0,"msg":err,"url":"/bis/index","wait":3})
		return
	}
	c.Redirect(http.StatusFound,"/bis/login")
}