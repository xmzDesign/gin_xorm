package models

type KernelMobilePrefix struct {
	Id       int64  `xorm:"pk comment('主键') unique BIGINT(20)"`
	Prefix   string `xorm:"not null comment('手机前缀') unique VARCHAR(32)"`
	Province string `xorm:"not null comment('省') VARCHAR(32)"`
	City     string `xorm:"not null comment('市') VARCHAR(32)"`
	Operator string `xorm:"not null comment('运营商') VARCHAR(32)"`
	AreaCode string `xorm:"not null comment('区号') VARCHAR(32)"`
	PostCode string `xorm:"not null comment('邮编') VARCHAR(32)"`
}
