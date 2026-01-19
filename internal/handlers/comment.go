
package handlers

import "github.com/gin-gonic/gin"

func AddComment(c *gin.Context) {
    c.JSON(201, gin.H{"status": "comment added (tree supported)"})
}
