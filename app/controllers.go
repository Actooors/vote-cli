package app

import (
	"github.com/gin-gonic/gin"
	"vote-cli/tools"
)

func ClearHandler(ctx *gin.Context) {
	dropTables()
	CreateTables()
	err := InitInsert()
	if err != nil {
		tools.Response(ctx, "FAIL", err.Error())
		return
	}
	tools.Response(ctx, "SUCCESS", nil)
}
