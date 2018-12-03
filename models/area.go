package models

type Area struct {
	Pinyin     string `xorm:"VARCHAR(32)"`
	Lng        string `xorm:"VARCHAR(32)"`
	Level      string `xorm:"VARCHAR(32)"`
	ParentId   string `xorm:"VARCHAR(32)"`
	AreaCode   string `xorm:"VARCHAR(32)"`
	Name       string `xorm:"VARCHAR(32)"`
	MergerName string `xorm:"VARCHAR(32)"`
	CityCode   string `xorm:"VARCHAR(32)"`
	ShortName  string `xorm:"VARCHAR(32)"`
	Id         string `xorm:"VARCHAR(32)"`
	ZipCode    string `xorm:"VARCHAR(32)"`
	Lat        string `xorm:"VARCHAR(32)"`
}
