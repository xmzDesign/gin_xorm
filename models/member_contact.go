package models

import (
	"time"
)

type MemberContact struct {
	Id               int64     `xorm:"pk autoincr BIGINT(20)"`
	MemberId         int64     `xorm:"comment('user_id') index BIGINT(20)"`
	DisplayName      string    `xorm:"comment('联系人姓名') index VARCHAR(255)"`
	TimesContacted   int       `xorm:"comment('通话次数') INT(11)"`
	LastTimeContacts time.Time `xorm:"comment('最后一次联系时间') DATETIME"`
	ContactNo        string    `xorm:"comment('联系号码') VARCHAR(100)"`
	CallDuration     int       `xorm:"comment('通话时长') INT(11)"`
	GmtCreate        time.Time `xorm:"comment('创建时间') DATETIME"`
	GmtModify        time.Time `xorm:"comment('修改时间') DATETIME"`
}
