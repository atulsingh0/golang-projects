package main

import (
	"fmt"
	"newsfeeder/httpd/handler"
	"newsfeeder/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("welcome to news feeder")

	r := gin.Default()

	feed := newsfeed.New()

	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))
	r.Run()

	// feed := newsfeed.New()
	// fmt.Println(feed)

	// feed.Add(newsfeed.Item{"Hello", "How are you?"})
	// fmt.Println(feed)

}
