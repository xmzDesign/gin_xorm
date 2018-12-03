package models

import (
	"time"
)

type CommonsBoutiqueApp struct {
	Id             int       `xorm:"not null pk autoincr INT(11)"`
	Name           string    `xorm:"not null comment('应用名称') index VARCHAR(32)"`
	Telephone      string    `xorm:"not null comment('客服电话') index VARCHAR(32)"`
	Status         int       `xorm:"not null comment('状态 10 运营 90 已下架') INT(11)"`
	StatusExplains string    `xorm:"not null comment('状态说明') VARCHAR(255)"`
	GmtCreate      time.Time `xorm:"not null DATETIME"`
	GmtModify      time.Time `xorm:"not null DATETIME"`
}
