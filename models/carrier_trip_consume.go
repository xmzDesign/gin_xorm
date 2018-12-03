package models

import (
	"time"
)

type CarrierTripConsume struct {
	Id          int64     `xorm:"pk autoincr BIGINT(20)"`
	OrderDate   string    `xorm:"comment('下单时间') VARCHAR(64)"`
	FlightSpend float32   `xorm:"comment('机票消费汇总') FLOAT(18,2)"`
	TrainSpend  float32   `xorm:"comment('火车票消费汇总') FLOAT(18,2)"`
	HotelSpend  float32   `xorm:"comment('酒店消费汇总') FLOAT(18,2)"`
	Ccount      int       `xorm:"comment('月度消费频次(/人次)') INT(11)"`
	TermDate    int       `xorm:"INT(11)"`
	GmtCreate   time.Time `xorm:"DATETIME"`
	SelfPhone   string    `xorm:"index VARCHAR(11)"`
}
