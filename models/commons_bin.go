package models

type CommonsBin struct {
	Id         string `xorm:"not null pk default '' VARCHAR(255)"`
	Orgno      string `xorm:"unique(idx_bin) VARCHAR(255)"`
	Bankname   string `xorm:"VARCHAR(255)"`
	Cardtype   string `xorm:"VARCHAR(255)"`
	Cardname   string `xorm:"VARCHAR(255)"`
	Cardlength string `xorm:"unique(idx_bin) VARCHAR(255)"`
	Binlegth   string `xorm:"VARCHAR(255)"`
	Bin        string `xorm:"unique(idx_bin) VARCHAR(255)"`
	Bankcode   string `xorm:"VARCHAR(255)"`
}
