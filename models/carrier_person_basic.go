package models

import (
	"time"
)

type CarrierPersonBasic struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	RealName       string    `xorm:"comment('姓名') VARCHAR(11)"`
	CertNo         string    `xorm:"comment('身份证号') VARCHAR(18)"`
	Gender         string    `xorm:"comment('性别') VARCHAR(2)"`
	Constellation  string    `xorm:"comment('星座') VARCHAR(5)"`
	Province       string    `xorm:"comment('省份') VARCHAR(18)"`
	City           string    `xorm:"comment('城市') VARCHAR(18)"`
	SelfPhone      string    `xorm:"index VARCHAR(11)"`
	Age            string    `xorm:"comment('年龄') VARCHAR(3)"`
	PhoneLocation  string    `xorm:"default '' comment('归属地') VARCHAR(25)"`
	Pstatus        int       `xorm:"comment('身份证解析状态,是否有效') TINYINT(1)"`
	RegTime        string    `xorm:"VARCHAR(64)"`
	GmtCreate      time.Time `xorm:"DATETIME"`
	TermDate       int       `xorm:"INT(11)"`
	OriginProvince string    `xorm:"comment('籍贯省份') VARCHAR(15)"`
}
