package service

import (
	"net/http"
	"server/data_access"

	"github.com/gin-gonic/gin"
)

func getAllShortenUrl(c *gin.Context, s *data_access.URLStore) {
	urls, err := s.GetAllShortenedURLs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve URLs"})
		return
	}
	c.JSON(http.StatusOK, urls)
}