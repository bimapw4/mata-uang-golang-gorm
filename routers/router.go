package routers

import (
	"mata-uang/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"data": "Hello World!!",
		})
	})

	router.GET("/uang", controller.GetUang)
}
