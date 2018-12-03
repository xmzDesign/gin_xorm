package models

import (
	"time"
)

type CarrierDeliverAddress struct {
	Id              int64     `xorm:"pk autoincr BIGINT(20)"`
	Address         string    `xorm:"VARCHAR(255)"`
	Lng             float32   `xorm:"comment('经度') FLOAT(18,2)"`
	Lat             float32   `xorm:"comment('纬度') FLOAT(18,2)"`
	PredictAddrType string    `xorm:"comment('地址类型') VARCHAR(64)"`
	BeginDate       string    `xorm:"comment('开始送货时间') VARCHAR(64)"`
	EndDate         string    `xorm:"comment('结束送货时间') VARCHAR(64)"`
	TotalAmount     float32   `xorm:"comment('总送货金额') FLOAT(18,2)"`
	TotalCount      int       `xorm:"comment('总送货次数') INT(11)"`
	Receiver        string    `xorm:"comment('收货人列表') TEXT"`
	Rcount          int       `xorm:"comment('收货次数') INT(11)"`
	Amount          float32   `xorm:"comment('收货金额') FLOAT(18,2)"`
	Rname           string    `xorm:"comment('收货人名称') VARCHAR(18)"`
	PhoneNumList    string    `xorm:"comment('收货人号码') TEXT"`
	TermDate        int       `xorm:"comment('过期时间') INT(11)"`
	GmtCreate       time.Time `xorm:"DATETIME"`
	SelfPhone       string    `xorm:"index VARCHAR(11)"`
}
