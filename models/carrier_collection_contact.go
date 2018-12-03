package models

import (
	"time"
)

type CarrierCollectionContact struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	ContactName    string    `xorm:"comment('联系人姓名') index VARCHAR(18)"`
	BeginDate      string    `xorm:"comment('和联系人最早联系的时间') VARCHAR(25)"`
	EndDate        string    `xorm:"comment('和联系人最晚联系的时间') VARCHAR(25)"`
	TotalCount     int       `xorm:"comment('申请人为该联系人购货总次数') INT(11)"`
	TotalAmount    float32   `xorm:"comment('申请人为该联系人购货总金额') FLOAT(18,2)"`
	ContactDetails string    `xorm:"comment('呼叫人信息统计') TEXT"`
	PhoneNum       string    `xorm:"comment('联系人电话') VARCHAR(18)"`
	PhoneNumLoc    string    `xorm:"comment('联系人归属地') VARCHAR(64)"`
	CallCnt        int       `xorm:"comment('呼叫次数') INT(11)"`
	CallLen        float32   `xorm:"comment('呼叫时长') FLOAT(18,2)"`
	CallOutCnt     int       `xorm:"comment('主叫次数') INT(11)"`
	CallOutLen     float32   `xorm:"comment('主叫时长') FLOAT(18,2)"`
	CallInCnt      int       `xorm:"comment('被叫次数') INT(11)"`
	SmsCnt         int       `xorm:"comment('短信发送和接受次数') INT(11)"`
	TransStart     string    `xorm:"comment('在呼叫记录里面最早出现的时间') VARCHAR(64)"`
	TransEnd       string    `xorm:"comment('在呼叫记录里面最晚出现的时间') VARCHAR(64)"`
	TermDate       int       `xorm:"INT(11)"`
	GmtCreate      time.Time `xorm:"DATETIME"`
	SelfPhone      string    `xorm:"index VARCHAR(11)"`
}
