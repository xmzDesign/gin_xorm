package models

import (
	"time"
)

type LoanRepayRecord struct {
	Id               int64     `xorm:"pk autoincr comment('序号') BIGINT(20)"`
	MemberId         int64     `xorm:"comment('用户id') index BIGINT(20)"`
	OrderNo          string    `xorm:"comment('交易单号') index VARCHAR(30)"`
	AmtPay           string    `xorm:"default 0.00 comment('金额') DECIMAL(18,2)"`
	BillItemId       int64     `xorm:"comment('账单明细Id bill_item_info.id') index BIGINT(20)"`
	PromotionCode    string    `xorm:"comment('优惠码') VARCHAR(64)"`
	ActivityCode     string    `xorm:"comment('活动码') VARCHAR(64)"`
	OidChnl          int       `xorm:"comment('交易发生渠道代码,10：android,20：ios') INT(2)"`
	BankCradId       int64     `xorm:"comment('银行卡Id') BIGINT(20)"`
	OidBiz           string    `xorm:"comment('业务代码') VARCHAR(6)"`
	VaildTime        time.Time `xorm:"comment('交易单有效期') DATETIME"`
	PayStatus        int       `xorm:"comment('状态 10：待支付 20：支付成功 30：支付失败 40：支付关闭') index INT(2)"`
	RepayStatus      int       `xorm:"comment('还款状态 10：还款中 20：还款成功 30：还款失败') index INT(2)"`
	RepaySuccessTime time.Time `xorm:"comment('还款成功时间') DATETIME"`
	PayChannel       int       `xorm:"comment('支付渠道 10:（alipay支付宝）   20 宝付') INT(2)"`
	Memo             string    `xorm:"comment('备注信息') VARCHAR(100)"`
	GmtCreate        time.Time `xorm:"comment('创建时间') DATETIME"`
	GmtModify        time.Time `xorm:"comment('更新时间') DATETIME"`
	CreateName       string    `xorm:"comment('创建人') VARCHAR(50)"`
	ModifyName       string    `xorm:"comment('修改人') VARCHAR(50)"`
}
