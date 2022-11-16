package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PutUserBy(c *gin.Context) {
	log.Println("PutUserBy")
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "PutUserBy"})
}
