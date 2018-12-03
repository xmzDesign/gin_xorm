package models

import (
	"time"
)

type LoanWithdraw struct {
	Id                int64     `xorm:"pk autoincr BIGINT(20)"`
	OidNo             string    `xorm:"not null comment('提现表序号') index VARCHAR(20)"`
	OutOrderId        int64     `xorm:"not null comment('外部单号') index BIGINT(64)"`
	OidChnl           string    `xorm:"not null comment('渠道') VARCHAR(2)"`
	OidBiz            string    `xorm:"not null comment('业务代码') VARCHAR(6)"`
	AmtWithdraw       string    `xorm:"default 0.00 comment('提现金额') DECIMAL(18,2)"`
	AmtLoan           string    `xorm:"default 0.00 comment('手续费') DECIMAL(18,2)"`
	PayChnl           int       `xorm:"comment('提现渠道  10,lianlian 20 支付宝') INT(2)"`
	Memo              string    `xorm:"comment('付款备注') VARCHAR(255)"`
	Status            int       `xorm:"comment('状态,10：待审核,20：提现审核通过,等待提现处理  30：审核拒绝，40：提现成功 50：提现失败 60：提现异常 70：提现处理中 ') INT(11)"`
	AuditUserId       int64     `xorm:"comment('审核人ID') BIGINT(20)"`
	AuditUserName     string    `xorm:"comment('审核人姓名') VARCHAR(100)"`
	AuditRejectReason string    `xorm:"comment('审核拒绝理由') VARCHAR(100)"`
	AuditTime         time.Time `xorm:"comment('审核时间') DATETIME"`
	CustId            int64     `xorm:"not null comment('收款客户号') index BIGINT(20)"`
	CustAcctno        int64     `xorm:"not null BIGINT(20)"`
	CustBankId        int64     `xorm:"not null comment('收款客户银行Id') BIGINT(20)"`
	CustName          string    `xorm:"not null comment('收款客户名称') VARCHAR(64)"`
	IdentityCard      string    `xorm:"not null comment('身份证') VARCHAR(20)"`
	Telphone          string    `xorm:"comment('手机号') VARCHAR(20)"`
	PayCustAcctno     int64     `xorm:"not null comment('付款账户号') BIGINT(20)"`
	PayCustId         int64     `xorm:"not null comment('付款用户Id') BIGINT(20)"`
	AcctTime          time.Time `xorm:"comment('账务日期') DATETIME"`
	Version           string    `xorm:"comment('版本号') VARCHAR(20)"`
	GmtCreate         time.Time `xorm:"comment('创建时间') DATETIME"`
	GmtModify         time.Time `xorm:"comment('更新时间') DATETIME"`
	CreateName        string    `xorm:"comment('创建人') VARCHAR(50)"`
	ModifyName        string    `xorm:"comment('修改人') VARCHAR(50)"`
}
