package main

import (
	"bdmall/app/conf"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/mojocn/base64Captcha"
	"net/http"
	"reflect"
)

type CategoryForm struct {
	Name string `load:"Name"`
	Order int64 `load:"Order"`
	Status int64 `load:"Status"`
}

type Category struct {
	Name string
	Order int64
	Status int64
}

var store = base64Captcha.DefaultMemStore


func geneCa(w http.ResponseWriter,r *http.Request)  {
	c := base64Captcha.NewCaptcha(base64Captcha.NewDriverChinese(80,240,4,base64Captcha.OptionShowHollowLine,20,"ni",nil,[]string{"3Dumb.ttf"}), store)
	id, b64s, err := c.Generate()
	fmt.Println(id)
	fmt.Println(b64s)
	fmt.Println(err)
}


func main() {
	err := conf.DB.Table("o2o_bis_account").Where("id = ? and status - ? >= 0", 1, 2).UpdateColumn("status", gorm.Expr("status - ?", 2)).Error
	fmt.Println(err)
}


func load(src,dst interface{}) (err error) {
	defer func() {
		if r := recover();r != nil {
			err = fmt.Errorf("%v",r)
		}
	}()
	srcType := reflect.TypeOf(src)
	srcVal := reflect.ValueOf(src)
	dstType := reflect.TypeOf(dst).Elem()
	dstVal := reflect.ValueOf(dst).Elem()
	for i := 0;i < srcType.NumField() ;i++  {
		tag := srcType.Field(i).Tag.Get("load")
		_, ok := dstType.FieldByName(tag)
		if !ok || !dstVal.FieldByName(tag).CanSet(){
			continue
		}
		dstVal.FieldByName(tag).Set(srcVal.FieldByName(tag))
	}
	return
}
