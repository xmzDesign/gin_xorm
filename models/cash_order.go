package models

import (
	"time"
)

type CashOrder struct {
	Id            int64     `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	OrderNo       string    `xorm:"not null default '' comment('交易单号') VARCHAR(32)"`
	ProductId     int64     `xorm:"not null comment('产品编号') BIGINT(20)"`
	MemberId      int64     `xorm:"not null default 0 comment('用户ID') index BIGINT(11)"`
	TotalAmount   string    `xorm:"not null default 0.00 comment('总金额') DECIMAL(10,2)"`
	Amount        string    `xorm:"not null default 0.00 comment('金额') DECIMAL(10,2)"`
	Status        int       `xorm:"not null default 0 comment('状态 10 待处理 20转账成功 30结算成功 90作废') INT(11)"`
	SettleStatus  int       `xorm:"default 10 comment('结算状态 10待结算 20已结算30结算失败') INT(4)"`
	StatusExplain string    `xorm:"not null default '' comment('用户状态说明') VARCHAR(255)"`
	ApplyTime     time.Time `xorm:"not null comment('申请时间') DATETIME"`
	TransTime     time.Time `xorm:"not null comment('交易时间') DATETIME"`
	SettleTime    time.Time `xorm:"comment('结算时间') DATETIME"`
	Rate          string    `xorm:"not null default 0.0000 comment('扣率') DECIMAL(4,4)"`
	RateFee       string    `xorm:"not null default 0.00 comment('扣率手续费') DECIMAL(10,2)"`
	Fee           string    `xorm:"not null default 0.00 comment('手续费') DECIMAL(10,2)"`
	DebitCardId   int64     `xorm:"comment('j借记卡id') BIGINT(20)"`
	CreditCardId  int64     `xorm:"comment('信用卡id') BIGINT(11)"`
	TransNo       string    `xorm:"not null default '' comment('外部交易号') VARCHAR(255)"`
	TransUrl      string    `xorm:"not null default '' comment('交易跳转路径') VARCHAR(255)"`
	ChannelId     int       `xorm:"not null default 0 comment('渠道号') INT(11)"`
	TicketId      int       `xorm:"not null default 0 comment('优惠券id') INT(11)"`
	DeviceId      string    `xorm:"default '' comment('设备指纹') VARCHAR(500)"`
	LatLot        string    `xorm:"default '' comment('经纬度') VARCHAR(100)"`
	ApplyIp       string    `xorm:"not null default '' comment('申请IP') VARCHAR(32)"`
	Version       string    `xorm:"not null default '' comment('版本号') VARCHAR(30)"`
	DeviceType    string    `xorm:"not null default '' comment('10 wap 20 ios  30 android') VARCHAR(12)"`
	Remarks       string    `xorm:"not null default '' comment('备注') VARCHAR(500)"`
	CreateTime    time.Time `xorm:"not null comment('创建时间') DATETIME"`
	ModifyTime    time.Time `xorm:"not null comment('更新时间') DATETIME"`
	CreateName    string    `xorm:"not null default '' comment('创建人') VARCHAR(50)"`
	ModifyName    string    `xorm:"not null default '' comment('修改人') VARCHAR(50)"`
	WithdrawNo    string    `xorm:"default '000000' comment('提现') VARCHAR(32)"`
	SmsSeqn       string    `xorm:"comment('短信流水') VARCHAR(50)"`
}
