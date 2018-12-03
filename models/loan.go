package models

import (
	"time"
)

type Loan struct {
	Id                    int64     `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	MemberId              int64     `xorm:"not null comment('用户ID') index BIGINT(20)"`
	MemberName            string    `xorm:"not null comment('用户名称') index index VARCHAR(64)"`
	IdentifyCard          string    `xorm:"comment('身份证号') index VARCHAR(30)"`
	Telphone              string    `xorm:"comment('手机号码') VARCHAR(20)"`
	BankCardNo            string    `xorm:"comment('银行卡号') VARCHAR(40)"`
	LoanReason            string    `xorm:"not null comment('借款原因') VARCHAR(512)"`
	AmtLoan               string    `xorm:"default 0.00 comment('借款金额') DECIMAL(18,2)"`
	Stags                 int       `xorm:"not null comment('分期数') INT(11)"`
	LoanDays              int       `xorm:"not null comment('借款天数') INT(11)"`
	ProductCode           string    `xorm:"not null default '' comment('借款产品编号  10 大额贷款  20小额贷款') index VARCHAR(50)"`
	Status                int       `xorm:"not null comment('状态 10，待风控审核 （收单）20，风控审核拒绝 30，待提现审核 40，提现审核通过 50，提现审核拒绝 60，提现处理中   70提现成功   80，提现失败 91，已退款   92，已还款 99，已关闭') INT(11)"`
	BankCardId            int64     `xorm:"comment('银行卡ID') BIGINT(20)"`
	PayTime               time.Time `xorm:"comment('放款成功时间') DATETIME"`
	RiskAuditId           int64     `xorm:"comment('风控审核人') BIGINT(20)"`
	RiskAuditName         string    `xorm:"comment('风控审核人姓名') VARCHAR(128)"`
	RiskAuditRejectReason string    `xorm:"comment('风控审核拒绝原因') VARCHAR(512)"`
	RiskAuditTime         time.Time `xorm:"comment('风控审核时间') DATETIME"`
	WithdrawBillno        int64     `xorm:"comment('提现交易单号') BIGINT(20)"`
	WithdrawTime          time.Time `xorm:"comment('提现申请时间') DATETIME"`
	RatesNo               string    `xorm:"not null comment('费率编号') VARCHAR(256)"`
	CreditAuditCost       string    `xorm:"default 0.00 comment('额度审核费') DECIMAL(18,2)"`
	CreditInspectCost     string    `xorm:"default 0.00 comment('征信查询费') DECIMAL(18,2)"`
	ManageCost            string    `xorm:"default 0.00 comment('管理费') DECIMAL(18,2)"`
	LoanInterestCost      string    `xorm:"default 0.00 comment('利息') DECIMAL(18,2)"`
	IntroduceCost         string    `xorm:"default 0.00 comment('介绍费') DECIMAL(18,2)"`
	RiskGrade             string    `xorm:"comment('用户风险等级') VARCHAR(10)"`
	ActivityCode          string    `xorm:"comment('活动码') VARCHAR(100)"`
	CouponNo              string    `xorm:"comment('红包编号') VARCHAR(36)"`
	LoanProvider          string    `xorm:"default 'hd' comment('贷款方 hd:禾刀 qt:其他') VARCHAR(20)"`
	DeviceId              string    `xorm:"comment('设备指纹') VARCHAR(500)"`
	RiskReportId          string    `xorm:"comment('风控报告id') VARCHAR(200)"`
	LatLot                string    `xorm:"comment('经纬度') VARCHAR(100)"`
	ApplyIp               string    `xorm:"comment('申请IP') VARCHAR(32)"`
	Version               string    `xorm:"comment('版本号') VARCHAR(30)"`
	YsbContractNo         string    `xorm:"comment('赢生宝签约号') VARCHAR(50)"`
	LlContractNo          string    `xorm:"comment('连连签约号') VARCHAR(50)"`
	Remarks               string    `xorm:"comment('备注') VARCHAR(500)"`
	GmtCreate             time.Time `xorm:"comment('创建时间') DATETIME"`
	GmtModify             time.Time `xorm:"comment('更新时间') DATETIME"`
	CreateName            string    `xorm:"comment('创建人') VARCHAR(50)"`
	ModifyName            string    `xorm:"comment('修改人') VARCHAR(50)"`
	BlackBox              string    `xorm:"comment('同盾设备指纹移动端') VARCHAR(5000)"`
	AppType               string    `xorm:"comment('IOS,ANDROID') VARCHAR(12)"`
	RiskStatus            int       `xorm:"default 10 comment('风控状态  10，正常 20,运营商数据缺失') INT(2)"`
	DeviceType            int       `xorm:"INT(20)"`
}
