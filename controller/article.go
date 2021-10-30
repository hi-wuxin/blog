package controller

import (
	"blog/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ArticleController struct {
	sqlClient *gorm.DB
}

func NewArticle(db *gorm.DB) *ArticleController {

	return &ArticleController{
		sqlClient: db,
	}

}

func (this *ArticleController) AddArticle(ctx *gin.Context) {
	var params models.Article

	if err := ctx.BindJSON(&params); err != nil {
	}

}
