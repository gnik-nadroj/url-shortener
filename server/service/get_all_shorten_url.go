package service

import (
	"net/http"
	"server/data_access"
	"server/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func getAllShortenUrl(c *gin.Context, s *data_access.URLStore) {
	urls, err := s.GetAllShortenedURLs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve URLs"})
		return
	}

	session := sessions.Default(c)
	userID := session.Get("user")
	if userID == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Error"})
		return
	}

	userUrls :=	make([]model.Url, 0, len(urls))
	id := userID.(string)

	for _, url := range urls {
		if id == url.UserID {
			userUrls = append(userUrls, url)
		}
	}
	c.JSON(http.StatusOK, userUrls)
}