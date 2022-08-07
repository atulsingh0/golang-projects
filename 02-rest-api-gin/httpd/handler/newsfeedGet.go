package handler

import (
	"net/http"

	"newsfeeder/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func NewsfeedGet(feed *newsfeed.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {

		restuls := feed.GetAll()
		c.JSON(http.StatusOK, restuls)
	}
}
