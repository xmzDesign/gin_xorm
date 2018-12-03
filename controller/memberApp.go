package controller

import (
	"encoding/json"
	"fmt"
	"gin-xorm/db"
	"github.com/gin-gonic/gin"
	"github.com/lunny/log"
	"net/http"
	"strconv"
)

type MemberAppController struct {
	MemberAppDao db.MemberAppDao
}

func (mc MemberAppController) GetById(c *gin.Context) {
	param := c.Param("id")
	int, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		fmt.Print("类型转换错误")
		fmt.Print(err)
	}
	app, err := mc.MemberAppDao.Get(int)
	if err != nil {
		log.Fatal("获取memberapp异常", err)
	}
	log.Info("memberApp:{}", app)
	bytes, _ := json.Marshal(app)
	c.String(http.StatusOK, string(bytes))
}
