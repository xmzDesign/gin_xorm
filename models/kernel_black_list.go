package models

import (
	"time"
)

type KernelBlackList struct {
	Id        int       `xorm:"not null pk autoincr INT(11)"`
	Mobile    string    `xorm:"VARCHAR(32)"`
	Pid       string    `xorm:"VARCHAR(32)"`
	Name      string    `xorm:"VARCHAR(32)"`
	Source    string    `xorm:"VARCHAR(255)"`
	Remarks   string    `xorm:"VARCHAR(255)"`
	GmtCreate time.Time `xorm:"DATETIME"`
}
