package user

import (
	"github.com/gin-gonic/gin"
	"web/config/core"
)

func Login(ctx *gin.Context) {
	core.Jwt().LoginHandler(ctx)
}
