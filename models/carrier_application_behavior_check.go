package models

import (
	"time"
)

type CarrierApplicationBehaviorCheck struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	CheckPoint string    `xorm:"comment('检查点') index VARCHAR(255)"`
	Score      int       `xorm:"comment('绑定个数') INT(11)"`
	Result     string    `xorm:"comment('审查结果') VARCHAR(255)"`
	Evidence   string    `xorm:"comment('证据') VARCHAR(500)"`
	Type       string    `xorm:"comment('信息监测,或者行为监测') VARCHAR(25)"`
	GmtCreate  time.Time `xorm:"DATETIME"`
	TermDate   int       `xorm:"comment('过期时间') INT(11)"`
	SelfPhone  string    `xorm:"index VARCHAR(11)"`
}
