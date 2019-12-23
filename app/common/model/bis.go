package model

import "time"

type Bis struct {
	ID int64
	CityId int64 `form:"city_id"`
	Money float64
	Order int32
	Status int32
	Name string `form:"name"`
	Email string `form:"email"`
	Logo string `form:"logo"`
	LicenceLogo string `form:"licence_logo"`
	CityPath string
	Desc string `form:"desc"`
	BankInfo string `form:"bank_info"`
	BankName string `form:"bank_name"`
	BankUser string `form:"bank_user"`
	Faren string `form:"faren"`
	FarenTel string `form:"faren_tel"`
	CreateTime time.Time
}
