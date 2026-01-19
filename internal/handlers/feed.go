
package handlers

import "github.com/gin-gonic/gin"

func GetFeed(c *gin.Context) {
    c.JSON(200, gin.H{"status": "feed with time-decay ranking"})
}
