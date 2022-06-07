package controller

import (
	"mata-uang/database"
	"mata-uang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUang(ctx *gin.Context) {
	var uangs []models.Uang
	id := ctx.Query("id")
	base := database.DB.Debug()

	if id != "" {
		base = base.Where("id = ?", id)
	}

	base.Find(&uangs)

	ctx.JSON(http.StatusOK, gin.H{
		"data": uangs,
	})
}
