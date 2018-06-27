package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//当code为FAIL时，data为string类型返回给前端的消息
func Response(ctx *gin.Context, code string, data interface{}) {
	if code == "FAIL" {
		data = data.(string)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": data,
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": nil,
		"data":    data,
	})
}
