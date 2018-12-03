package models

import (
	"time"
)

type AgentPay struct {
	Id          int64     `xorm:"pk autoincr comment('代扣ID') BIGINT(20)"`
	OrderNo     string    `xorm:"comment('订单号') index VARCHAR(50)"`
	LoanId      int64     `xorm:"comment('借款单号') index BIGINT(20)"`
	BillId      int64     `xorm:"comment('账单Id') index BIGINT(20)"`
	BillItemId  int64     `xorm:"comment('账单Id') index BIGINT(20)"`
	AmtPay      string    `xorm:"not null comment('交易金额') DECIMAL(18,2)"`
	Status      int       `xorm:"default 10 comment('状态,10：受理成功 20：支付成功，30：支付失败 40：支付异常') index INT(2)"`
	SuccessDate time.Time `xorm:"comment('账务日期') DATETIME(4)"`
	MemberId    int64     `xorm:"not null comment('用户ID') index BIGINT(20)"`
	MemberName  string    `xorm:"comment('用户名称') VARCHAR(64)"`
	PayChannel  int       `xorm:"comment('代扣渠道') index INT(2)"`
	GmtCreate   time.Time `xorm:"comment('创建时间') DATETIME"`
	GmtModify   time.Time `xorm:"comment('更新时间') DATETIME"`
	CreateName  string    `xorm:"comment('创建人') VARCHAR(50)"`
	ModifyName  string    `xorm:"comment('修改人') VARCHAR(50)"`
}
