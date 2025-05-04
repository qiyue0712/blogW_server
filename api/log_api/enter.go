package log_api

import (
	"blogW_server/common"
	"blogW_server/common/res"
	"blogW_server/global"
	"blogW_server/models"
	"blogW_server/models/enum"
	"blogW_server/service/log_service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type LogApi struct {
}

type LogListRequest struct {
	common.PageInfo
	LogType     enum.LogType      `form:"logType"` // 日志级别 1 2 3
	Level       enum.LogLevelType `form:"level"`
	UserID      uint              `form:"userId"`
	IP          string            `gorm:"size:32" form:"ip"`
	LoginStatus bool              `form:"loginStatus"` // 登录的状态
	ServiceName string            `gorm:"size:32" form:"serviceName"`
}

type LogListResponse struct {
	models.LogModel
	UserNickname string `json:"userNickname"`
	UserAvatar   string `json:"userAvatar"`
}

func (LogApi) LogListView(c *gin.Context) {
	// 分页 查询（精确查询，模糊匹配）
	var cr LogListRequest         // 参数绑定
	err := c.ShouldBindQuery(&cr) // 参数匹配
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	// 通用查询接口封装 list_query
	list, count, err := common.ListQuery(models.LogModel{
		LogType:     cr.LogType,
		Level:       cr.Level,
		UserID:      cr.UserID,
		IP:          cr.IP,
		LoginStatus: cr.LoginStatus,
		ServiceName: cr.ServiceName,
	}, common.Options{
		PageInfo:     cr.PageInfo,
		Likes:        []string{"title"},
		Preloads:     []string{"UserModel"},
		DefaultOrder: "created_at desc",
	})
	// 响应之后的list显示 nickname avatar
	var _list = make([]LogListResponse, 0)
	for _, logModel := range list {
		_list = append(_list, LogListResponse{
			LogModel:     logModel,
			UserNickname: logModel.UserModel.Nickname,
			UserAvatar:   logModel.UserModel.Avatar,
		})
	}

	res.OkWithList(_list, int(count), c)
	return
}

func (LogApi) LogReadView(c *gin.Context) {
	var cr models.IDRequest // 日志的读取接口
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	var log models.LogModel
	err = global.DB.Take(&log, cr.ID).Error
	if err != nil {
		res.FailWithMsg("不存在的日志", c)
		return
	}
	if !log.IsRead {
		global.DB.Model(&log).Update("is_read", true)
	}
	res.OkWithMsg("日志读取成功", c)
}

func (LogApi) LogRemoveView(c *gin.Context) {
	var cr models.RemoveRequest // 日志的删除接口
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	log := log_service.GetLog(c)
	log.ShowRequest()
	log.ShowResponse()

	var logList []models.LogModel
	global.DB.Find(&logList, "id in ?", cr.IDList)

	if len(logList) > 0 {
		global.DB.Delete(&logList)
	}

	msg := fmt.Sprintf("日志删除成功，共删除%d条", len(logList))

	res.OkWithMsg(msg, c)

}
