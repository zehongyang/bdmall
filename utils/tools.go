package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"math/rand"
	"net/http"
	"reflect"
	"strings"
	"time"
)

//百度地图根据地址获取坐标返回对象
type MapLocationRes struct {
	Status int `json:"status"`
	Result struct{
		Location struct{
			Lng float64 `json:"lng"`
			Lat float64 `json:"lat"`
		} `json:"location"`
	} `json:"result"`
}

//结构体字段映射
func LoadStruct(src interface{},dst interface{}) (err error) {
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
		if !ok || !dstVal.FieldByName(tag).CanSet() {
			continue
		}
		dstVal.FieldByName(tag).Set(srcVal.FieldByName(tag))
	}
	return
}


//成功跳转
func ResSuccess(c *gin.Context,h gin.H)  {
	c.HTML(http.StatusOK,"common/tpl/jump.html",h)
}

//错误跳转
func ResErrors(c *gin.Context,h gin.H)  {
	c.HTML(http.StatusOK,"common/tpl/jump.html",h)
}

//时间格式化
func TimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

//百度api根据地名获取经纬度
func GetXyByAddress(address string) (longitude float64,latitude float64,ok bool) {
	var url = fmt.Sprintf("http://api.map.baidu.com/geocoding/v3/?address=%s&output=json&ak=Q9mFwBA5MsoRuXiUzBsNbdgLVjLbL9QW",address)
	resp, err := http.Get(url)
	if err != nil {
		ok = false
		return
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ok = false
		return
	}
	var res MapLocationRes
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		ok = false
		return
	}
	if res.Status == 0 {
		longitude = res.Result.Location.Lng
		latitude = res.Result.Location.Lat
		ok = true
	}
	return
}

//邮件发送
func SendEmail(to string,subject string,content string) error {
	m := gomail.NewMessage()
	m.SetHeader("From","15808381785@163.com")
	m.SetHeader("To",to)
	m.SetHeader("Subject",subject)
	m.SetBody("text/html",content)
	dialer := gomail.NewDialer("smtp.163.com", 25, "15808381785@163.com", "yangzehong123")
	return dialer.DialAndSend(m)
}

//生成随机字符串
func GenerateString(length int) string {
	if length <= 0  {
		return  ""
	}
	str := []byte{'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z',
		'A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z','0','1','2',
	'3','4','5','6','7','8','9'}
	rand.Seed(time.Now().UnixNano())
	var sb  strings.Builder
	for i := 0;i < length ;i++  {
		index := rand.Intn(len(str))
		if err := sb.WriteByte(str[index]);err != nil {
			return ""
		}
	}
	return sb.String()
}

