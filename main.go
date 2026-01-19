
package main

import (
    "github.com/gin-gonic/gin"
    "reddit-go/internal/handlers"
)

func main() {
    r := gin.Default()

    r.POST("/posts", handlers.CreatePost)
    r.GET("/feed", handlers.GetFeed)
    r.POST("/vote", handlers.Vote)
    r.POST("/comments", handlers.AddComment)

    r.Run(":8080")
}
