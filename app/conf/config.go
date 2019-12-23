package conf

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/ini.v1"
	"log"
)

//http服务配置
type Server struct {
	ServerName string
}
//数据库配置
type DataBase struct {
	Port int64
	Host string
	UserName string
	Password string
	BaseName string
}

var (
	ServerConfig *Server
	DataBaseConfig *DataBase
	DB *gorm.DB
)

//初始配置文件
func init()  {
	cfg, err := ini.Load("./app/conf/config.ini")
	if err != nil {
		log.Fatal(err)
	}
	ServerConfig = new(Server)
	DataBaseConfig = new(DataBase)
	ServerConfig.ServerName = cfg.Section("server").Key("server_name").MustString("localhost:8888")
	DataBaseConfig.Port = cfg.Section("database").Key("port").MustInt64(3306)
	DataBaseConfig.Host = cfg.Section("database").Key("host").MustString("127.0.0.1")
	DataBaseConfig.UserName = cfg.Section("database").Key("username").MustString("root")
	DataBaseConfig.Password = cfg.Section("database").Key("password").MustString("root")
	DataBaseConfig.BaseName = cfg.Section("database").Key("basename").MustString("bdmall")
	//初始化数据库
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",DataBaseConfig.UserName,
		DataBaseConfig.Password,DataBaseConfig.Host,DataBaseConfig.Port,DataBaseConfig.BaseName))
	if err != nil {
		log.Fatal(err)
	}
	//表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "o2o_"+defaultTableName
	}
	//关闭表复数
	DB.SingularTable(true)
	DB.LogMode(true)
}
