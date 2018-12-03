package models

import (
	"time"
)

type CarrierMainService struct {
	Id              int64     `xorm:"pk autoincr BIGINT(20)"`
	CompanyName     string    `xorm:"VARCHAR(64)"`
	CompanyType     string    `xorm:"comment('服务企业类型') VARCHAR(11)"`
	SelfPhone       string    `xorm:"index VARCHAR(11)"`
	TotalServiceCnt int       `xorm:"comment('总互动次数') INT(11)"`
	ServiceDetails  string    `xorm:"comment('月互动详情') TEXT"`
	InteractCnt     int       `xorm:"comment('月互动次数') INT(11)"`
	InteractMth     string    `xorm:"comment('互动月份') VARCHAR(25)"`
	TermDate        int       `xorm:"comment('过期时间') INT(11)"`
	GmtCreate       time.Time `xorm:"DATETIME"`
}
