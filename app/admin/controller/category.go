package controller

import (
	"bdmall/app/admin/validate"
	"bdmall/app/common/model"
	"bdmall/app/conf"
	"bdmall/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"strconv"
	"strings"
)


//分类列表页面
func CategoryIndex(c *gin.Context)  {
	pid, err := strconv.Atoi(c.DefaultQuery("pid", "0"))
	if err != nil {
		utils.ResErrors(c,gin.H{"code":0,"msg":err,"url":"/admin/welcome","wait":3})
	}
	err, categories := model.GetCategories(int64(pid))
	if err != nil {
		utils.ResErrors(c,gin.H{"code":0,"msg":err,"url":"/admin/welcome","wait":3})
	}
	c.HTML(http.StatusOK,"admin/category/index.html",gin.H{"categories":categories})
}

//新增分类页面展示
func CategoryCreate(c *gin.Context)  {
	err, categories := model.GetCategories(0)
	if err != nil {
		utils.ResErrors(c,gin.H{"code":0,"msg":err,"url":"/admin/welcome","wait":3})
	}
	c.HTML(http.StatusOK,"admin/category/add.html",gin.H{"categories":categories})
}

//添加分类
func CategoryStore(c *gin.Context)  {
	var categoryForm validate.CategoryForm
	err := c.ShouldBind(&categoryForm)
	if err != nil {
		utils.ResErrors(c,gin.H{"code":0,"msg":err,"url":"/admin/category/create","wait":3})
		return
	}
	//设置中文提示,注册验证方法
	validate := validator.New()
	err = validate.RegisterValidation("CheckCategoryNameExists", categoryForm.CheckCategoryNameExists)
	if err != nil {
		utils.ResErrors(c,gin.H{"code":0,"msg":err,"url":"/admin/category/create","wait":3})
		return
	}
	zh_ch := zh.New()
	uni := ut.New(zh_ch)
	trans, _ := uni.GetTranslator("zh")
	err = zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		utils.ResErrors(c,gin.H{"code":0,"msg":err,"url":"/admin/category/create","wait":3})
		return
	}
	err = validate.Struct(&categoryForm)
	if err != nil {
		var errMsg []string
		for _, err := range err.(validator.ValidationErrors) {
			if err.Tag() == "CheckCategoryNameExists" {
				errMsg = append(errMsg,"Name分类名称已存在")
			}else{
				errMsg = append(errMsg,err.Translate(trans))
			}
		}
		msg := strings.Join(errMsg,";")
		utils.ResErrors(c,gin.H{"code":0,"msg":msg,"url":"/admin/category/create","wait":3})
		return
	}
	//数据加载
	var category model.Category
	err = utils.LoadStruct(categoryForm, &category)
	if err != nil {
		utils.ResErrors(c,gin.H{"code":0,"msg":err,"url":"/admin/category/create","wait":3})
		return
	}
	//保存数据
	err = conf.DB.Create(&category).Error
	if err != nil {
		utils.ResErrors(c,gin.H{"code":0,"msg":err,"url":"/admin/category/create","wait":3})
		return
	}
	utils.ResSuccess(c,gin.H{"code":1,"msg":"新增分类成功","url":"/admin/category/create","wait":3})
}

//分类编辑页面
func CategoryEdit(c *gin.Context)  {
	var category model.Category
	conf.DB.First(&category,c.Query("id"))
	_, categories := model.GetCategories(0)
	c.HTML(http.StatusOK,"admin/category/edit.html",gin.H{"category":category,"categories":categories})
}

//分类更新
func CategoryUpdate(c *gin.Context)  {
	//绑定表单数据
	var categoryForm validate.CategoryForm
	err := c.ShouldBind(&categoryForm)
	if err != nil {
		utils.ResErrors(c,gin.H{"code":0,"msg":err,"url":"/admin/category/edit","wait":3})
		return
	}
	//表单验证
	validate := validator.New()
	err = validate.RegisterValidation("CheckCategoryNameExists", categoryForm.CheckCategoryNameExists)
	if err != nil {
		utils.ResErrors(c,gin.H{"code":0,"msg":err,"url":"/admin/category/edit","wait":3})
		return
	}
	zh_ch := zh.New()
	uni := ut.New(zh_ch)
	trans, _ := uni.GetTranslator("zh")
	err = zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		utils.ResErrors(c,gin.H{"code":0,"msg":err,"url":"/admin/category/create","wait":3})
		return
	}
	err = validate.Struct(&categoryForm)
	if err != nil {
		var errMsg []string
		for _, err := range err.(validator.ValidationErrors) {
			if err.Tag() == "CheckCategoryNameExists" {
				errMsg = append(errMsg,"Name分类名称已存在")
			}else{
				errMsg = append(errMsg,err.Translate(trans))
			}
		}
		msg := strings.Join(errMsg,";")
		utils.ResErrors(c,gin.H{"code":0,"msg":msg,"url":"/admin/category/edit","wait":3})
		return
	}
	//模型转换
	var category model.Category
	err = utils.LoadStruct(categoryForm, &category)
	if err != nil {
		utils.ResErrors(c,gin.H{"code":0,"msg":err,"url":"/admin/category/edit","wait":3})
		return
	}
	//更新数据
	err = conf.DB.Save(&category).Error
	if err != nil {
		utils.ResErrors(c,gin.H{"code":0,"msg":err,"url":"/admin/category/edit","wait":3})
		return
	}
	utils.ResSuccess(c,gin.H{"code":1,"msg":"修改分类成功","url":"/admin/category/edit","wait":3})
}

//获取分类列表
func GetCategories(c *gin.Context)  {
	cid, err := strconv.Atoi(c.PostForm("cid"))
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"message":"参数错误"})
		return
	}
	err, categories := model.GetCategories(int64(cid))
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"code":500,"message":"服务端错误"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"message":"成功","data":categories})
}
