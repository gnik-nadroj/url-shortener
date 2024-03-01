package service

import (
	"server/common"
	"server/data_access"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)


func SetupServer() {
    router := gin.Default()
    urlStore := data_access.NewURLStore()
    userStore := data_access.NewUserStore()

    config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true 
	router.Use(cors.New(config))

    sessionStore, _ := redis.NewStore(10, "tcp", common.GetEnv(common.RedisUrl), common.GetEnv(common.RedisPassword), []byte(common.GetEnv(common.SessionSecret)))
	router.Use(sessions.Sessions("mysession", sessionStore))

    router.POST("/signup", func(ctx *gin.Context) {
        signUp(ctx, userStore)
    })

    router.POST("/login", func(ctx *gin.Context) {
        login(ctx, userStore)
    })

    router.GET("/logout", AuthRequired(), func(ctx *gin.Context) {
        logout(ctx)
    })

    router.GET("/:short", func(c *gin.Context) {
       redirect(c, urlStore)
    })

    router.GET("/stats", AuthRequired(), func(c *gin.Context) {
        getAllShortenUrl(c, urlStore)
    })

    router.POST("/shortener", AuthRequired(), func(c *gin.Context) {
        shortener(c, urlStore)
    })

    port := common.GetEnv("SERVER_PORT")

    router.Run(":" + port)
}