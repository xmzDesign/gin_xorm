package models

import (
	"time"
)

type CarrierTripInfo struct {
	Id                 int64     `xorm:"pk autoincr BIGINT(20)"`
	TripLeave          string    `xorm:"comment('出发地') VARCHAR(64)"`
	TripDest           string    `xorm:"comment('目的地') VARCHAR(64)"`
	TripTransportation string    `xorm:"comment('多种出行交通工具') TEXT"`
	TripPerson         string    `xorm:"comment('多个同行人') TEXT"`
	TripType           string    `xorm:"comment('出行时间类型') VARCHAR(64)"`
	SelfPhone          string    `xorm:"index VARCHAR(11)"`
	TripStartTime      string    `xorm:"comment('出行开始时间') VARCHAR(25)"`
	TripEndTime        string    `xorm:"comment('出行结束时间') VARCHAR(25)"`
	TripDataSource     string    `xorm:"comment('是通过哪些数据源分析出来的') TEXT"`
	TermDate           int       `xorm:"INT(11)"`
	GmtCreate          time.Time `xorm:"DATETIME"`
}
