package models

// 文章表

type ArticleModel struct {
	Model
	Title        string    `gorm:"size:32" json:"title"`
	Abstract     string    `gorm:"size:256" json:"abstract"`
	Content      string    `json:"content"`
	CategoryID   uint      `json:"categoryID"`                                   // 分类的id
	TagList      []string  `gorm:"type:longtext;serializer:json" json:"tagList"` // 标签列表
	Cover        string    `gorm:"size:256" json:"cover"`
	UserID       uint      `json:"userID"`
	UserModel    UserModel `gorm:"foreignKey:UserID" json:"-"`
	LookCount    int       `json:"lookCount"`
	DiggCount    int       `json:"diggCount"`
	CommentCount int       `json:"commentCount"`
	CollectCount int       `json:"collectCount"`
	OpenComment  bool      `json:"openComment"` // 是否开启评论
	Status       int8      `json:"status"`      // 状态 草稿 审核中 已发布
}
