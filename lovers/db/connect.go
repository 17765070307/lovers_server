package db

import (
	"fmt"

	"github.com/sirupsen/logrus"
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

func init() {
	path := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", userName, password, ip, port, dbName)
	DB, _ = gorm.Open(mysql.Open(path), &gorm.Config{})
	DB.Table("user").Create(&User{})
	DB.Table("weight_record").Create(&WeightRecord{})
	logrus.Infof("init gorm config successfully, path=%s", path)
}
