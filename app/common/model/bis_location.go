package model

import "time"

type BisLocation struct {
	ID int64
	Xpoint float64
	Ypoint float64
	BisId int64
	IsMain int64
	CityId int64 `form:"city_id"`
	CategoryId int64 `form:"category_id"`
	Order int32
	Status int32
	Name string `form:"name"`
	Logo string `form:"logo"`
	Address string `form:"address"`
	Tel string `form:"tel"`
	Contact string `form:"contact"`
	OpenTime string `form:"open_time"`
	Content string `form:"content"`
	CityPath string
	CategoryPath string
	BankInfo string `form:"bank_info"`
	CreateTime time.Time
}
