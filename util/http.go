package util

import (
	"blog/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	sqlLink = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=UTC"
)

type Response struct {
	ErrCode int         `json:"err_code"`
	Data    interface{} `json:"data"`
	ErrMsg  string      `json:"err_msg"`
}

type BlogContext struct {
	gin.Context
}

type BlogHttp struct {
	*gin.Engine
	SqlClient *gorm.DB
}

func NewBlogHttp(g *gin.Engine, sqlConfig config.Mysql) *BlogHttp {
	dsn := fmt.Sprintf(sqlLink, sqlConfig.Name, sqlConfig.Password, sqlConfig.Addr, sqlConfig.Port, sqlConfig.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &BlogHttp{
		g,
		db,
	}
}
