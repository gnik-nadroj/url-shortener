package service

import (
	"server/common"
	"server/data_access"
	"github.com/gin-gonic/gin"
)


func SetupServer() {
    router := gin.Default()
    store := data_access.NewURLStore()

    router.GET("/:short", func(c *gin.Context) {
       redirect(c, store)
    })

    router.GET("/stats", func(c *gin.Context) {
        getAllShortenUrl(c, store)
    })

    router.POST("/shortener", func(c *gin.Context) {
        shortener(c, store)
    })

    port := common.GetEnv("SERVER_PORT")

    router.Run(":" + port)
}