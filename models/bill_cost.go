package models

import (
	"time"
)

type BillCost struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	BillId        int64     `xorm:"default 0 comment('bill_id') index BIGINT(20)"`
	BillItemId    int64     `xorm:"default 0 comment('bill_item_id') index BIGINT(20)"`
	MemberId      int64     `xorm:"default 0 comment('用户ID') index BIGINT(20)"`
	Amount        string    `xorm:"default 0.00 comment('金额') DECIMAL(18,2)"`
	CostUniqueKey string    `xorm:"comment('费用唯一key（日期+子账单ID+类型）') unique index VARCHAR(30)"`
	BillCostType  int       `xorm:"not null default 0 comment('类型, 10:逾期费, 20:逾期减免  30:利息减免 40:管理费减免 50:优惠 60:代扣减免') index INT(2)"`
	GmtCreate     time.Time `xorm:"comment('创建时间') DATETIME"`
	GmtModify     time.Time `xorm:"comment('更新时间') DATETIME"`
	CreateName    string    `xorm:"comment('创建人') VARCHAR(50)"`
	ModifyName    string    `xorm:"comment('修改人') VARCHAR(50)"`
}
