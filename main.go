package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/:program", func(context *gin.Context) {

	})

	_ = r.Run()
}
