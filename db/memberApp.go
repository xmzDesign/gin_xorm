package db

import (
	"errors"
	"gin-xorm/models"
)

type MemberAppDao struct {
}

func (MemberAppDao) Get(id int64) (models.MemberApp, error) {
	//app := new(models.MemberApp)
	var app models.MemberApp
	has, err := engine.ID(id).Get(&app)
	if err != nil {
		return app, err
	}
	if !has {
		return app, errors.New("dept not found")
	}
	return app, nil
}

//使用sql文件
func (MemberAppDao) QueryByName(name string) ([]models.MemberApp, error) {
	var memberApps []models.MemberApp
	param := make(map[string]interface{})
	param["name"] = name
	err := engine.SqlTemplateClient("memberapp.name.sql", &param).Find(&memberApps)
	if err != nil {
		return nil, err
	}
	return memberApps, nil
}

//使用xml
func (MemberAppDao) QueryByXMLName(name string) ([]models.MemberApp, error) {
	var memberApps []models.MemberApp
	param := make(map[string]interface{})
	param["name"] = name
	has, err := engine.SqlMapClient("queryByName", name).Get(&memberApps)
	if err != nil {
		return nil, err
	}
	if !has {
		return memberApps, errors.New("dept not found")
	}
	return memberApps, nil
}
