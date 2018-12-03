package models

import (
	"time"
)

type Member struct {
	Id                 int64     `xorm:"pk autoincr comment('id') BIGINT(20)"`
	Mobile             string    `xorm:"not null default '' comment('手机号') index VARCHAR(11)"`
	NickName           string    `xorm:"not null default '' comment('昵称') VARCHAR(20)"`
	Sex                int       `xorm:"not null default 0 comment('性别    10男性，20女性 ，90未知性别') INT(2)"`
	RealName           string    `xorm:"not null comment('真实姓名') index VARCHAR(20)"`
	IdCard             string    `xorm:"not null comment('身份证号') index VARCHAR(18)"`
	Authentication     int       `xorm:"not null comment('实名认证状态 10，未认证 ；20，审核中  30审核失败    90已认证') INT(11)"`
	HeadPortrait       string    `xorm:"not null default '' comment('头像') VARCHAR(256)"`
	RegTime            time.Time `xorm:"not null comment('注册时间') DATETIME"`
	ConsumeFlag        int       `xorm:"not null default 0 comment('是否消费 0,未消费 ;1,已消费') INT(1)"`
	FirstConsumeTime   time.Time `xorm:"not null comment('首次消费时间') DATETIME"`
	LastLoginTime      time.Time `xorm:"not null comment('最后登录时间') DATETIME"`
	StartLevel         int       `xorm:"not null default 0 comment('星级') INT(2)"`
	OpenId             string    `xorm:"not null VARCHAR(30)"`
	RegIp              string    `xorm:"not null VARCHAR(20)"`
	RegLonLat          string    `xorm:"not null comment('注册经纬度') VARCHAR(50)"`
	Status             int       `xorm:"not null comment('10 正常，  20 异常  ，90冻结') INT(2)"`
	BindBankcardFlag   int       `xorm:"not null default 10 comment('银行卡标识 10,未绑定银行卡 20，已绑定银行卡') INT(2)"`
	BindCreditcardFlag int       `xorm:"not null default 10 comment('信用卡绑定标识 10，未绑定 ;20，已绑定') INT(2)"`
	GmtCreate          time.Time `xorm:"comment('创建时间') DATETIME"`
	GmtModify          time.Time `xorm:"comment('更新时间') DATETIME"`
	CreateName         string    `xorm:"comment('创建人') VARCHAR(50)"`
	ModifyName         string    `xorm:"comment('修改人') VARCHAR(50)"`
	LonLatCountry      string    `xorm:"comment('经纬度地区') VARCHAR(32)"`
	LonLatProvince     string    `xorm:"comment('经纬度省') VARCHAR(32)"`
	LonLatCity         string    `xorm:"comment('经纬度市') VARCHAR(32)"`
	LonLatDistrict     string    `xorm:"comment('经纬度地区') VARCHAR(32)"`
	IpProvince         string    `xorm:"comment('ip省') VARCHAR(32)"`
	IpCity             string    `xorm:"comment('ip市') VARCHAR(32)"`
	MemberType         string    `xorm:"default 'prod' comment('账号类型,  prod:生产, test:测试') VARCHAR(10)"`
	IsValid            int       `xorm:"default 10 comment('10:有效, 20:无效') INT(11)"`
	Source             string    `xorm:"comment('渠道来源') VARCHAR(32)"`
	DeviceId           string    `xorm:"comment('设备id') VARCHAR(64)"`
	RiskType           int       `xorm:"default 0 INT(2)"`
}
