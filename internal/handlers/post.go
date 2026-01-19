
package handlers

import "github.com/gin-gonic/gin"

func CreatePost(c *gin.Context) {
    c.JSON(201, gin.H{"status": "post created"})
}
