package models

type BankBranchNo struct {
	Id         int64  `xorm:"pk autoincr BIGINT(20)"`
	BranchName string `xorm:"comment('支行名称') VARCHAR(55)"`
	BranchNo   string `xorm:"comment('联行号') VARCHAR(55)"`
	BankNo     int    `xorm:"comment('银行识别号') index INT(11)"`
	AreaCode   int    `xorm:"comment('区域编码') index INT(11)"`
}
