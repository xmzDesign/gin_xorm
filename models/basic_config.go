package models

import (
	"time"
)

type BasicConfig struct {
	Id              int       `xorm:"not null pk autoincr INT(11)"`
	LinkConfig      int       `xorm:"default 0 comment('关联的配置项, basic_config.id') INT(11)"`
	Cgroup          string    `xorm:"comment('配置组') unique(idx_group_key) VARCHAR(255)"`
	Ckey            string    `xorm:"comment('key') unique(idx_group_key) VARCHAR(255)"`
	Cvalue          string    `xorm:"comment('value') VARCHAR(255)"`
	Cdesc           string    `xorm:"comment('描述') VARCHAR(255)"`
	CreateName      string    `xorm:"comment('创建人') VARCHAR(255)"`
	ModifyName      string    `xorm:"comment('修改人') VARCHAR(255)"`
	GmtCreate       time.Time `xorm:"comment('创建时间') DATETIME"`
	GmtModify       time.Time `xorm:"comment('修改时间') DATETIME"`
	ExternalVisable int       `xorm:"default 0 comment('是否外部可见') TINYINT(1)"`
}
