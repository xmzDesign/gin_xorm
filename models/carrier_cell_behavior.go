package models

import (
	"time"
)

type CarrierCellBehavior struct {
	Id          int64     `xorm:"pk autoincr BIGINT(20)"`
	OperatorZh  string    `xorm:"comment('运营商（中文）') VARCHAR(18)"`
	PhoneNum    string    `xorm:"comment('号码') index VARCHAR(22)"`
	Loc         string    `xorm:"comment('归属地') VARCHAR(18)"`
	Cmonth      string    `xorm:"comment('月份') VARCHAR(25)"`
	Cnt         int       `xorm:"comment('呼叫次数') INT(11)"`
	OutCnt      int       `xorm:"comment('主叫次数') INT(11)"`
	OutTime     float32   `xorm:"comment('主叫时间(分)') FLOAT(18,2)"`
	InCnt       int       `xorm:"comment('被叫次数') INT(11)"`
	InTime      float32   `xorm:"comment('被叫时间(分)') FLOAT(18,2)"`
	NetFlow     float32   `xorm:"comment('流量（MB）') FLOAT(18,2)"`
	SmsCnt      int       `xorm:"comment('短信数量') INT(11)"`
	TotalAmount float32   `xorm:"comment('总费用(元)') FLOAT(18,2)"`
	GmtCreate   time.Time `xorm:"DATETIME"`
	TermDate    int       `xorm:"comment('过期时间') INT(11)"`
	Operator    string    `xorm:"comment('运营商(en)') VARCHAR(18)"`
}
