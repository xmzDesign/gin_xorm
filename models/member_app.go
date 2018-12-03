package models

import (
	"time"
)

type MemberApp struct {
	Id           int64     `xorm:"pk autoincr BIGINT(20)"`
	MemberId     int64     `xorm:"not null comment('会员ID') index BIGINT(20)"`
	PackageName  string    `xorm:"comment('包名') VARCHAR(255)"`
	AppName      string    `xorm:"not null comment('app名称') index VARCHAR(100)"`
	StartupTime  time.Time `xorm:"comment('启动时间') DATETIME"`
	RunningKeep  int       `xorm:"comment('运行时长') INT(11)"`
	BackResident int       `xorm:"comment('常驻后台, 0:否, 1:是') INT(11)"`
	GmtCreate    time.Time `xorm:"not null DATETIME"`
	GmtModify    time.Time `xorm:"not null DATETIME"`
}
