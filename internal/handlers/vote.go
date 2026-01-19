
package handlers

import "github.com/gin-gonic/gin"

func Vote(c *gin.Context) {
    c.JSON(200, gin.H{"status": "vote registered"})
}
