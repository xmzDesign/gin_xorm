package models

type CardBin2 struct {
	Id           int    `xorm:"not null pk autoincr INT(20)"`
	BankName     string `xorm:"comment('标记') VARCHAR(55)"`
	CardName     string `xorm:"VARCHAR(55)"`
	AtmStatus    int    `xorm:"comment('0:不适用 1：适用') INT(2)"`
	PosStatus    int    `xorm:"comment('0:不适用  1：适用') INT(2)"`
	CardLength   string `xorm:"comment('主账号长度') VARCHAR(20)"`
	CardNo       string `xorm:"comment('主账号') VARCHAR(55)"`
	BinLength    string `xorm:"comment('发卡行标识长度') index VARCHAR(20)"`
	Bin          string `xorm:"comment('发卡行标识') index VARCHAR(20)"`
	CardTypeName string `xorm:"VARCHAR(55)"`
	CardType     string `xorm:"default '10' comment('卡类型 10:借记卡 20：贷记卡  30:预付费卡  40:准贷记卡') VARCHAR(10)"`
	BankCode     string `xorm:"VARCHAR(55)"`
	RealBankCode string `xorm:"comment('真实bankCode') VARCHAR(55)"`
	TermDate     int    `xorm:"comment('失效时间') INT(20)"`
}
