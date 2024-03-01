package service

import (
	"net/http"
	"server/common"
	"server/data_access"
	internal_encoding "server/encoding"

	"github.com/gin-contrib/sessions"
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

	session := sessions.Default(c)
	userID := session.Get("user")
	if userID == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Error"})
		return
	}

	count, _ := s.GetShortenedURLCount()

	hash := internal_encoding.Base62Encode(uint64(count))

	err := s.Insert(hash, request.URL, userID.(string))
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not shorten URL"})
		return
	}

	shortenUrl := common.IdToUrl(hash)

	c.JSON(http.StatusOK, gin.H{"shortURL": shortenUrl})
}
