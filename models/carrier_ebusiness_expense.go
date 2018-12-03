package models

import (
	"time"
)

type CarrierEbusinessExpense struct {
	Id        int64     `xorm:"pk autoincr BIGINT(20)"`
	TransMth  string    `xorm:"comment('汇总月份') VARCHAR(64)"`
	AllAmount float32   `xorm:"comment('总购物金额') FLOAT(18,2)"`
	AllCount  int       `xorm:"comment('总购物次数') INT(11)"`
	Category  string    `xorm:"comment('商品的品类分析') TEXT"`
	TermDate  int       `xorm:"comment('过期时间') INT(11)"`
	GmtCreate time.Time `xorm:"DATETIME"`
	SelfPhone string    `xorm:"index VARCHAR(11)"`
}
