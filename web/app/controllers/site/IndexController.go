package site

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(ctx *gin.Context) {
	ctx.Next()
	fmt.Println("ddddd")
	ctx.String(http.StatusOK, "dsasdf")
}
