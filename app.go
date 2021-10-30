package main

import (
	"blog/config"
	"blog/controller"
	"blog/util"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

var Config struct {
	DataBases struct {
		Mysql config.Mysql
		Redis config.Redis
	}
	Oss config.OSS
}

func main() {

	b, err := ioutil.ReadFile("config/config.toml")
	if err != nil {
		panic(err)
	}
	err = toml.Unmarshal(b, &Config)
	if err != nil {
		panic(err)
	}
	blog := util.NewBlogHttp(gin.Default(), Config.DataBases.Mysql)
	article := controller.NewArticle(blog.SqlClient)
	blog.GET("/demo", article.AddArticle)
	blog.LoadHTMLGlob("web/*")
	blog.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", "flysnow_org")
	})
	blog.GET("/about", func(ctx *gin.Context) {
		ctx.HTML(200, "about.html", "flysnow_org")
	})
	//后台登陆验证
	admin := blog.Group("/admin")
	//验证登陆信息
	admin.Use(func(context *gin.Context) {

	})
	//后台静态页面
	admin.Group("/web")
	blog.Run(":8080")

}
