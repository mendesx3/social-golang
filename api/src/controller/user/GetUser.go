package user

import (
	"api/src/repository/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserBy(c *gin.Context) {
	log.Println("GetUserBy")
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "GetUserBy"})
}

func GetAllUser(c *gin.Context) {
	u := user.GetAll()
	c.JSON(http.StatusOK, u)
}
