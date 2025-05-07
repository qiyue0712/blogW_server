package image_api

import (
	"blogW_server/common"
	"blogW_server/common/res"
	"blogW_server/global"
	"blogW_server/models"
	"blogW_server/service/log_service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ImageApi struct {
}

type ImageListResponse struct {
	models.ImageModel
	WebPath string `json:"webPath"`
}

func (ImageApi) ImageListView(c *gin.Context) {
	var cr common.PageInfo
	c.ShouldBindQuery(&cr)
	_list, count, _ := common.ListQuery(models.ImageModel{}, common.Options{
		PageInfo: cr,
		Likes:    []string{"filename"},
	})
	var list = make([]ImageListResponse, 0)
	for _, model := range _list {
		list = append(list, ImageListResponse{
			ImageModel: model,
			WebPath:    model.WebPath(),
		})
	}
	res.OkWithList(list, count, c)

} // 列表和删除

func (ImageApi) ImageRemoveView(c *gin.Context) {
	var cr models.RemoveRequest // 删除接口
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	log := log_service.GetLog(c)
	log.ShowRequest()
	log.ShowResponse()

	var list []models.ImageModel
	global.DB.Find(&list, "id in ?", cr.IDList)

	var successCount, errCount int64
	if len(list) > 0 {
		successCount = global.DB.Delete(&list).RowsAffected
	}
	errCount = int64(len(list)) - successCount
	msg := fmt.Sprintf("操作成功，成功%d 失败%d", successCount, errCount)

	res.OkWithMsg(msg, c)
}
