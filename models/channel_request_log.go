package models

import (
	"time"
)

type ChannelRequestLog struct {
	Id          int64     `xorm:"pk autoincr BIGINT(20)"`
	OrderNo     string    `xorm:"comment('订单号') VARCHAR(50)"`
	ChannelCode string    `xorm:"not null comment('渠道编号') index VARCHAR(50)"`
	ChannelName string    `xorm:"not null comment('渠道名称') VARCHAR(50)"`
	BillItemId  int64     `xorm:"comment('子账单Id') index BIGINT(20)"`
	OrderAmt    string    `xorm:"DECIMAL(20,2)"`
	Request     string    `xorm:"comment('请求数据') VARCHAR(2000)"`
	Response    string    `xorm:"comment('响应数据') VARCHAR(2000)"`
	AcceptCode  string    `xorm:"comment('受理状态码') index VARCHAR(50)"`
	ProcessCode string    `xorm:"comment('处理状态码') index VARCHAR(50)"`
	RetMsg      string    `xorm:"comment('返回信息') VARCHAR(500)"`
	MemberId    int64     `xorm:"comment('用户编号') BIGINT(20)"`
	GmtCreate   time.Time `xorm:"DATETIME"`
	GmtModify   time.Time `xorm:"DATETIME"`
}
