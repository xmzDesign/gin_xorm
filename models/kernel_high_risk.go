package models

type KernelHighRisk struct {
	Id      int    `xorm:"not null pk autoincr INT(11)"`
	Name    string `xorm:"not null comment('名称') VARCHAR(32)"`
	Code    string `xorm:"not null comment('编码') VARCHAR(32)"`
	Value   string `xorm:"not null comment('值') VARCHAR(32)"`
	Remarks string `xorm:"not null comment('说明') VARCHAR(255)"`
}
