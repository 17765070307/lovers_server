package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	userName = "root"
	password = "qwer5684002"
	ip       = "cdb-kc99r1uy.bj.tencentcdb.com"
	port     = "10073"
	dbName   = "Lovers"
)

type Table interface {
	TableName() string
}

var DB *gorm.DB

//func init() {
//	path := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", userName, password, ip, port, dbName)
//	DB, _ = sql.Open("mysql", path)
//	//defer DB.Close()
//	//if err != nil {
//	//	fmt.Println(err.Error())
//	//}
//	DB.SetConnMaxLifetime(100)
//	DB.SetMaxIdleConns(10)
//	if err := DB.Ping(); err != nil {
//		fmt.Println("opon database fail")
//		return
//	}
//	fmt.Println("connnect success")
//}

func init() {
	path := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", userName, password, ip, port, dbName)
	DB, _ = gorm.Open(mysql.Open(path), &gorm.Config{})
	DB.Table("user").Create(&User{})
	DB.Table("weight_record").Create(&WeightRecord{})
}
