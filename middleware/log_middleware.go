package middleware

import (
	"blogW_server/service/log_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseWriter struct { // 1. 定义增强的ResponseWriter
	gin.ResponseWriter        // 继承原始响应写入器
	Body               []byte // 新增字段用于存储响应内容
	Head               http.Header
}

// 拿到Head方法

func (w *ResponseWriter) Header() http.Header {
	return w.Head
}

// 2. 重写Write方法

func (w *ResponseWriter) Write(data []byte) (int, error) {
	w.Body = append(w.Body, data...)    // 捕获数据到Body字段
	return w.ResponseWriter.Write(data) // 继续原始写入流程
}

func LogMiddleware(c *gin.Context) {
	log := log_service.NewActionLogByGin(c) // 操作日志建立
	log.SetRequest(c)                       // 请求中间件
	c.Set("log", log)                       // 请求->视图 log传递

	res := &ResponseWriter{
		ResponseWriter: c.Writer,          // 创建自定义Writer
		Head:           make(http.Header), // 创建自定义Head
	}
	c.Writer = res // 替换上下文中的Writer
	c.Next()
	// 响应中间件
	log.SetResponse(res.Body)
	log.SetResponseHeader(res.Head)
	log.MiddlewareSave() // 保存入库

}
