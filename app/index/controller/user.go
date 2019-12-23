package controller

import (
	"bdmall/app/common/model"
	"bdmall/app/conf"
	"bdmall/utils"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLoginIndex(c *gin.Context)  {
    c.HTML(http.StatusOK,"home/user/login.html",gin.H{})
}

func UserRegisterIndex(c *gin.Context)  {
	c.HTML(http.StatusOK,"home/user/register.html",gin.H{})
}

//用户注册
func UserRegister(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	id := c.PostForm("id")
	code := c.PostForm("code")
	email := c.PostForm("email")
	ok := Store.Verify(id, code, true)
	if !ok {
		utils.ResErrors(c,gin.H{"code":0,"msg":"验证码不匹配","url":"/user/register","wait":3})
		return
	}
	var user model.User
	user.Username = username
	user.Password = password
	user.Email = email
	user.Code = utils.GenerateString(10)
	user.Password = fmt.Sprintf("%x",md5.Sum([]byte(fmt.Sprintf("%x%s",md5.Sum([]byte(user.Password)),user.Code))))
	if err := conf.DB.Create(&user).Error;err != nil{
		utils.ResErrors(c,gin.H{"code":0,"msg":err,"url":"/user/register","wait":3})
		return
	}
	c.Redirect(http.StatusFound,"/user/login")
}

//用户登陆
func UserLogin(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	var user model.User
	err := conf.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		utils.ResErrors(c,gin.H{"code":0,"msg":"用户名或密码错误","url":"/user/login","wait":3})
		return
	}
	if user.ID <= 0 {
		utils.ResErrors(c,gin.H{"code":0,"msg":"用户名或密码错误","url":"/user/login","wait":3})
		return
	}
	pwd := fmt.Sprintf("%x",md5.Sum([]byte(fmt.Sprintf("%x%s",md5.Sum([]byte(password)),user.Code))))
	if pwd != user.Password {
		utils.ResErrors(c,gin.H{"code":0,"msg":"用户名或密码错误","url":"/user/login","wait":3})
		return
	}
	//保存session
	session, _ := utils.SessionStore.Get(c.Request, "home_user")
	session.Values["username"] = username
	session.Values["id"] = user.ID
	_ = session.Save(c.Request, c.Writer)
	c.Redirect(http.StatusFound,"/")
}

//用户退出
func UserLogOut(c *gin.Context)  {
	session, _ := utils.SessionStore.Get(c.Request, "home_user")
	session.Values["username"] = nil
	session.Values["id"] = nil
	_ = session.Save(c.Request, c.Writer)
	c.Redirect(http.StatusFound,"/")
}