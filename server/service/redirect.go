package service

import (
	"net/http"
	"server/data_access"

	"github.com/gin-gonic/gin"
)


func redirect(c *gin.Context, s *data_access.URLStore) {
	shortURL := c.Param("short")
	originalURL, err := s.GetOriginalURL(shortURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, originalURL)
}