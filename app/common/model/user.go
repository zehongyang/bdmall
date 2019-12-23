package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID int64
	Order int32
	Status int32
	Username string
	Password string
	Code string
	LastLoginIp string
	Email string
	Mobile string
	CreateTime time.Time
}

//钩子函数
func (c *User) BeforeCreate(scope *gorm.Scope) error {
	c.CreateTime = time.Now()
	return nil
}
