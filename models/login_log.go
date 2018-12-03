package models

import (
	"time"
)

type LoginLog struct {
	Id             int64     `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	MemberId       int64     `xorm:"not null default 0 comment('会员id') index BIGINT(20)"`
	LoginMobile    string    `xorm:"not null default '' comment('登录账号名') index VARCHAR(20)"`
	LoginTime      time.Time `xorm:"comment('登录时间') DATETIME"`
	LoginIp        string    `xorm:"not null comment('登录ip') index VARCHAR(20)"`
	LonLat         string    `xorm:"not null default '' comment('经纬度') VARCHAR(50)"`
	Type           string    `xorm:"not null default '' comment('android 安卓  ios 苹果') VARCHAR(30)"`
	DeviceId       string    `xorm:"not null default '' comment('设备号') index VARCHAR(120)"`
	Mac            string    `xorm:"not null default '' comment('设备mac') VARCHAR(30)"`
	Version        string    `xorm:"not null default '' comment('版本号') VARCHAR(30)"`
	Imei           string    `xorm:"not null default '' comment('串号IMEI') VARCHAR(30)"`
	Sn             string    `xorm:"not null default '' comment('序列号') VARCHAR(30)"`
	Brand          string    `xorm:"not null default '' comment('设备厂商') VARCHAR(30)"`
	Model          string    `xorm:"not null default '' comment('设备名称') VARCHAR(30)"`
	Sim            string    `xorm:"not null default '' comment('sim卡电话卡') VARCHAR(30)"`
	IsRoot         int       `xorm:"not null default 0 comment('是否root权限') INT(1)"`
	NetType        string    `xorm:"not null default '' comment('网络状态 2G 3G 4G') VARCHAR(30)"`
	LonLatCountry  string    `xorm:"comment('经纬度国家') VARCHAR(32)"`
	LonLatProvince string    `xorm:"comment('经纬度省') VARCHAR(32)"`
	LonLatCity     string    `xorm:"comment('经纬度市') VARCHAR(32)"`
	LonLatDistrict string    `xorm:"comment('经纬度区域') VARCHAR(32)"`
	IpProvince     string    `xorm:"comment('省') VARCHAR(32)"`
	IpCity         string    `xorm:"comment('市') VARCHAR(32)"`
}
