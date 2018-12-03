package models

import (
	"time"
)

type CardInfo struct {
	Id         int64     `xorm:"pk autoincr comment('id') BIGINT(20)"`
	MemberId   int64     `xorm:"not null comment('用户id') index BIGINT(20)"`
	Owner      string    `xorm:"not null default '' comment('持卡人') VARCHAR(50)"`
	Cardno     string    `xorm:"not null default '' comment('卡号(银行卡，信用卡)') index VARCHAR(50)"`
	Idno       string    `xorm:"not null default '' comment('身份证号') index VARCHAR(50)"`
	CardName   string    `xorm:"not null default '' comment('银行名称') VARCHAR(50)"`
	CardImg    string    `xorm:"not null default '' comment('卡图标') VARCHAR(100)"`
	Cvn2       string    `xorm:"not null default '' comment('安全码') VARCHAR(300)"`
	Mobile     string    `xorm:"not null default '' comment('手机号') index VARCHAR(20)"`
	CardType   int       `xorm:"not null default 10 comment('卡类型 10,银行卡;20,信用卡 ') INT(1)"`
	Status     int       `xorm:"not null default 10 comment('卡状态 20,正常;10,作废 ') INT(1)"`
	SortVal    int       `xorm:"not null default 0 comment('排序') INT(1)"`
	MasterFlag int       `xorm:"not null default 10 comment('主卡判别 ,10普通 ,20主卡') INT(1)"`
	Branch     string    `xorm:"not null default '' comment('支行名称') VARCHAR(255)"`
	BranchNo   string    `xorm:"not null default '' comment('支行编号') VARCHAR(255)"`
	ValidDate  string    `xorm:"not null default '' comment('信用卡有效日期') VARCHAR(300)"`
	BankCode   string    `xorm:"not null default '' comment('银行编号') VARCHAR(10)"`
	LlAgreeNo  string    `xorm:"comment('连连签约号') VARCHAR(32)"`
	GmtCreate  time.Time `xorm:"comment('创建时间') DATETIME"`
	GmtModify  time.Time `xorm:"comment('更新时间') DATETIME"`
	CreateName string    `xorm:"comment('创建人') VARCHAR(50)"`
	ModifyName string    `xorm:"comment('修改人') VARCHAR(50)"`
	IsDefault  int       `xorm:"default 0 TINYINT(1)"`
	DueDay     int       `xorm:"default 0 comment('还款日') TINYINT(2)"`
	BillDay    int       `xorm:"comment('账单日') TINYINT(2)"`
}
