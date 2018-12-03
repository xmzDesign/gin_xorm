package models

import (
	"time"
)

type LoanDevicesInfo struct {
	Id        int64     `xorm:"pk autoincr BIGINT(20)"`
	LoanId    int64     `xorm:"not null comment('借款Id') unique BIGINT(20)"`
	BlackBox  string    `xorm:"TEXT"`
	AppType   string    `xorm:"comment('IOS,ANDROID') VARCHAR(255)"`
	GmtCreate time.Time `xorm:"comment('创建时间') TIME(4)"`
	GmtModify time.Time `xorm:"comment('更新时间') TIME(4)"`
}
