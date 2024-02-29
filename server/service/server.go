package service

import (
	"net/http"
	"server/common"
	"server/data_access"

	"github.com/gin-gonic/gin"
)


func SetupServer() {
    router := gin.Default()
    store := data_access.NewURLStore()

    router.GET("/:short", func(c *gin.Context) {
        shortURL := c.Param("short")
        originalURL, err := store.GetOriginalURL(shortURL)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
            return
        }
        c.Redirect(http.StatusMovedPermanently, originalURL)
    })

    router.GET("/stats", func(c *gin.Context) {
        urls, err := store.GetAllShortenedURLs()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve URLs"})
            return
        }
        c.JSON(http.StatusOK, urls)
    })

    router.POST("/shortener", func(c *gin.Context) {
        var request struct {
            URL string `json:"url" binding:"required"`
        }
        if err := c.ShouldBindJSON(&request); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        err := store.Insert(request.URL, request.URL)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not shorten URL"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"shortURL": request.URL})
    })

    port := common.GetEnv("SERVER_PORT")

    router.Run(":" + port)
}