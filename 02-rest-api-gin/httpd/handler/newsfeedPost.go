package handler

import (
	"log"
	"net/http"
	"newsfeeder/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

type newsfeedPostReq struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func NewsfeedPost(feed *newsfeed.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {

		reqBody := newsfeedPostReq{}
		c.Bind(&reqBody)

		log.Println(reqBody)
		item := newsfeed.Item{
			Title: reqBody.Title,
			Post:  reqBody.Post,
		}
		feed.Add(item)
		c.Status(http.StatusNoContent)
	}
}
