package model

import "time"

type BisAccount struct {
	ID int64
	BisId int64
	IsMain int32
	Order int32
	Status int64
	Username string `form:"username"`
	Password string `form:"password"`
	Code string
	LastLoginIp string
	LastLoginTime time.Time
	CreateTime time.Time
}
