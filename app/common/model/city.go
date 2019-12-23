package model

import (
	"bdmall/app/conf"
	"time"
)

type City struct {
	ID int64 `json:"id"`
	Order int32 `json:"order"`
	Pid int32 `json:"pid"`
	Status int64 `json:"status"`
	Name string `json:"name"`
	Uname string `json:"uname"`
	CreateTime time.Time `json:"create_time"`
}

//根据pid获取城市
func GetCities(pid int32) (error,[]City) {
	var cities []City
	err := conf.DB.Where("pid=?", pid).Find(&cities).Error
	if err != nil {
		return err,nil
	}
	return nil,cities
}