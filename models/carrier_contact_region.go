package models

import (
	"time"
)

type CarrierContactRegion struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	RegionLoc      string    `xorm:"comment('联系人的号码归属地') VARCHAR(18)"`
	UniqNumCnt     int       `xorm:"comment('去重后的联系人号码数量') INT(11)"`
	CallInCnt      int       `xorm:"comment('电话呼入次数') INT(11)"`
	CallOutCnt     int       `xorm:"comment('电话呼出次数') INT(11)"`
	CallInTime     float32   `xorm:"comment('总呼入时间(分)') FLOAT(18,2)"`
	CallOutTime    float32   `xorm:"comment('总呼出时间(分)') FLOAT(18,2)"`
	AvgCallInTime  float32   `xorm:"comment('平均呼入时间') FLOAT(18,2)"`
	AvgCallOutTime float32   `xorm:"comment('平均呼出时间') FLOAT(18,2)"`
	CallInCntPct   float32   `xorm:"comment('电话呼入次数百分比') FLOAT(18,2)"`
	CallOutCntPct  float32   `xorm:"comment('电话呼出次数百分比') FLOAT(18,2)"`
	CallInTimePct  float32   `xorm:"comment('电话呼入时间百分比') FLOAT(18,2)"`
	CallOutTimePct float32   `xorm:"comment('电话呼出时间百分比') FLOAT(18,2)"`
	GmtCreate      time.Time `xorm:"DATETIME"`
	TermDate       int       `xorm:"INT(11)"`
	SelfPhone      string    `xorm:"index VARCHAR(11)"`
}
