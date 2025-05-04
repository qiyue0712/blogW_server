package flags

import (
	"blogW_server/global"
	"blogW_server/models"
	"github.com/sirupsen/logrus"
)

// 迁移表的数据

func FlagDB() {
	err := global.DB.AutoMigrate(
		&models.UserModel{},                   // 用户表
		&models.UserConfModel{},               // 用户配置表
		&models.ArticleModel{},                // 文章表
		&models.CategoryModel{},               // 文章分类表
		&models.ArticleDiggModel{},            // 文章点赞表
		&models.CollectModel{},                // 收藏表
		&models.UserArticleCollectModel{},     // 用户文章收藏表
		&models.UserTopArticleModel{},         // 置顶文章表
		&models.ImageModel{},                  // 图片表
		&models.UserArticleLookHistoryModel{}, // 用户浏览的文章历史表
		&models.CommentModel{},                // 评论表
		&models.BannerModel{},                 // banner表
		&models.LogModel{},                    // 日志表
		&models.UserLoginModel{},              // 用户登录表
		&models.GlobalNotificationModel{},     // 全局通知表
	)
	if err != nil {
		logrus.Errorf("数据库迁移失败 %s", err)
		return
	}
	logrus.Infof("数据库迁移成功")
}
