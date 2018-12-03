package models

import (
	"time"
)

type CommonsMobileLocation struct {
	Id        int       `xorm:"not null pk autoincr comment('主键') INT(11)"`
	Mobile    string    `xorm:"comment('手机号') index VARCHAR(32)"`
	Province  string    `xorm:"comment('省') VARCHAR(32)"`
	City      string    `xorm:"comment('市') VARCHAR(32)"`
	Company   string    `xorm:"comment('运营商') VARCHAR(32)"`
	CardType  string    `xorm:"comment('cardtype') VARCHAR(32)"`
	AreaCode  string    `xorm:"comment('区号') VARCHAR(32)"`
	GmtCreate time.Time `xorm:"comment('创建时间') DATETIME"`
}
