package validate

import (
	"bdmall/app/common/model"
	"bdmall/app/conf"
	"github.com/go-playground/validator/v10"
)

type CategoryForm struct {
	ID int64 `form:"id" load:"ID"`
	Name string `form:"name" validate:"required,max=30,CheckCategoryNameExists" load:"Name"`
	Pid int64 `form:"pid" load:"Pid"`
}

//验证分类名称是否存在
func (c *CategoryForm) CheckCategoryNameExists(fl validator.FieldLevel) bool {
	var count int
	var err error
	if c.ID == 0 {
		err = conf.DB.Model(&model.Category{}).Where("name = ?", fl.Field().String()).Count(&count).Error
		if err != nil || count == 0{
			return  true
		}
	}else{
		err = conf.DB.Model(&model.Category{}).Where("name = ? and id <> ?", fl.Field().String(), c.ID).Count(&count).Error
		if err != nil || count == 0{
			return  true
		}
	}
	return false
}