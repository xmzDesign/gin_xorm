package models

import (
	"time"
)

type BankCard4Factor struct {
	Id          int64     `xorm:"pk autoincr BIGINT(20)"`
	RealName    string    `xorm:"comment('姓名') index(i_yaosu) VARCHAR(100)"`
	BankCard    string    `xorm:"not null comment('银行卡号') index(i_yaosu) VARCHAR(100)"`
	Mobile      string    `xorm:"comment('预留手机号') index(i_yaosu) VARCHAR(100)"`
	CertNo      string    `xorm:"comment('身份证号') index(i_yaosu) VARCHAR(100)"`
	Pass        int       `xorm:"comment('是否通过四要素校验') TINYINT(1)"`
	Description string    `xorm:"comment('渠道返回结果') VARCHAR(200)"`
	GmtCreate   time.Time `xorm:"DATETIME"`
	GmtModify   time.Time `xorm:"DATETIME"`
}
