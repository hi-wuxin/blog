package models

import "gorm.io/gorm"

// User 用户信息
type User struct {
	Uid      uint   `json:"uid" gorm:"primaryKey"`
	Name     string `json:"name"`
	Password string `json:"password"`
	LoginIp  string `json:"login_ip"`
	Updated  int64  `gorm:"autoUpdateTime"` // 使用时间戳毫秒数填充更新时间
	Created  int64  `gorm:"autoCreateTime"` // 使用时间戳秒数填充创建时间
}

type ArticleKind struct {
	gorm.Model
	Title string `json:"title"`
}

// Article 文章信息
type Article struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Title   string `json:"title"`
	Editer  string `json:"editer"`
	Uid     uint   `json:"uid"`
	Content string `json:"content"`
	IP      string `json:"ip"`
}

// Comment 评论内容
type Comment struct {
	gorm.Model
	IP      string `json:"ip"`
	Content string `json:"content"`
}
