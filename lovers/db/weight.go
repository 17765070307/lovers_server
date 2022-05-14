package db

import (
	"time"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type WeightRecord struct {
	Id         int64     `gorm:"type:int;not null;column:id;comment:主键;autoIncrement" json:"id"`
	UserCode   string    `gorm:"type:varchar(32);not null;column:user_code;comment:用户编号" json:"user_code"`
	Weight     float32   `gorm:"type:double;null;column:weight;comment:体重记录" json:"weight"`
	RecordDate time.Time `gorm:"type:timestamp;null;column:record_date;comment:记录日期" json:"record_date"`
	CreatedAt  time.Time `gorm:"type:timestamp;null;column:create_time;comment:创建时间" json:"create_time"`
	UpdatedAt  time.Time `gorm:"type:timestamp;null;column:modify_time;comment:更新时间" json:"modify_time"`
}

func (WeightRecord) TableName() string {
	return "weight_record"
}

// QueryWeightRecord
// @Description: 查询体重记录
// @param ctx 上下文信息
// @param start 开始时间
// @param end 结束时间
// @return []WeightRecord 体重信息
// @return error 异常信息
func QueryWeightRecord(ctx *gin.Context, start, end string) ([]WeightRecord, error) {
	var records []WeightRecord
	var result *gorm.DB
	if len(start) == 0 && len(end) == 0 {
		result = DB.Find(&records)
	} else if len(start) == 0 {
		result = DB.Where("record_date <= ?", end).Find(&records)
	} else if len(end) == 0 {
		result = DB.Where("record_date >= ?", start).Find(&records)
	} else {
		result = DB.Where("record_date between ? and ?", start, end).Find(&records)
	}
	if result.Error != nil {
		return []WeightRecord{}, result.Error
	}
	return records, nil
}

// InsertWeightRecord
// @Description: 插入体重数据
// @param ctx 上下文
// @param record 记录信息
// @return error 异常信息
func InsertWeightRecord(ctx *gin.Context, record WeightRecord) error {
	result := DB.Create(&record)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
