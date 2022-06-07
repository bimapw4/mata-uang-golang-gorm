package main

import (
	"mata-uang/database"
	"mata-uang/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Database()

	router := gin.Default()
	routers.Router(router)

	router.Run()
}
