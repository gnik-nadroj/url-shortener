package service

import (
	"net/http"
	"server/data_access"
	internal_encoding "server/encoding"

	"github.com/gin-gonic/gin"
)

func shortener(c *gin.Context, s *data_access.URLStore) {
	var request struct {
		URL string `json:"url" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	count, _ := s.GetShortenedURLCount()

	shortenUrl := internal_encoding.Base62Encode(uint64(count))
	err := s.Insert(shortenUrl, request.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not shorten URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"shortURL": shortenUrl})
}