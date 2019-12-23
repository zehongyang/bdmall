package model

import (
	"bdmall/app/conf"
	"github.com/jinzhu/gorm"
	"time"
)

type Category struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Pid int64 `json:"pid"`
	Order int64 `json:"order"`
	Status int64 `json:"status"`
	CreateTime time.Time `json:"create_time"`
}

//获取分类列表
func GetCategories(pid int64) (err error,categories []Category) {
	err = conf.DB.Where("pid=?", pid).Find(&categories).Error
	return
}

//新增分类前钩子函数
func (c *Category) BeforeCreate(scope *gorm.Scope) error {
	c.CreateTime = time.Now()
	return nil
}