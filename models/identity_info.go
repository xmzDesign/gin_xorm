package models

import (
	"time"
)

type IdentityInfo struct {
	Id                    int64     `xorm:"pk autoincr comment('id') BIGINT(20)"`
	MemberId              int64     `xorm:"not null comment('会员ID') index BIGINT(20)"`
	ImgPersonFace         string    `xorm:"not null default '' comment('人脸图片') VARCHAR(256)"`
	ImgIdentifyFront      string    `xorm:"not null default '' comment('身份证正面图片') VARCHAR(256)"`
	ImgIdentifyBack       string    `xorm:"not null default '' comment('身份证反面图片') VARCHAR(256)"`
	ImgIdentifyPortrait   string    `xorm:"not null default '' comment('身份证头像') VARCHAR(255)"`
	ImgIdentifyHand       string    `xorm:"not null default '' comment('手持身份证图片') VARCHAR(256)"`
	OperatorsExpireTime   time.Time `xorm:"comment('手机运营商认证有效时间') index DATETIME"`
	SesameExpireTime      time.Time `xorm:"comment('芝麻信用过期时间') index DATETIME"`
	AuditTime             time.Time `xorm:"comment('审核时间') DATETIME"`
	SesameCredit          int       `xorm:"default 0 comment('芝麻信用分') INT(11)"`
	IdentityStatus        int       `xorm:"not null default 10 comment('身份认证状态 10，未认证 20，审核中 30，审核失败 90 已认证') INT(2)"`
	SesameCreditStatus    int       `xorm:"not null default 10 comment('芝麻信用状态 10 未认证 20 已认证 30已过期') INT(2)"`
	MobileOperatorsStatus int       `xorm:"not null default 10 comment('手机运营商认证 10 未认证 ,20 已认证 30已过期') INT(2)"`
	FaceAuthStatus        int       `xorm:"default 10 comment('人脸认证 10，未认证 90, 已认证') INT(2)"`
	PersonalInfoStatus    int       `xorm:"default 10 comment('个人资料认证状态 10，未认证 20，审核中 30，审核失败 90 已认证') INT(2)"`
	Remark                string    `xorm:"not null default '' comment('备注') VARCHAR(255)"`
	ModifyName            string    `xorm:"comment('修改人') VARCHAR(50)"`
	CreateName            string    `xorm:"comment('创建人') VARCHAR(50)"`
	Education             int       `xorm:"default 90 comment('学历信息 10初中以下 20专科 30 高中 40本科 50硕士 60博士以上 99 未知') INT(2)"`
	Address               string    `xorm:"default '' comment('居住地址') VARCHAR(255)"`
	Company               string    `xorm:"default '' comment('公司名称') VARCHAR(255)"`
	CompanyAddress        string    `xorm:"default '' comment('公司地址') VARCHAR(255)"`
	GmtCreate             time.Time `xorm:"comment('创建时间') DATETIME"`
	GmtModify             time.Time `xorm:"comment('更新时间') DATETIME"`
}
