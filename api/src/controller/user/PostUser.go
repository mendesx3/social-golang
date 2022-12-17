package user

import (
	"api/src/model/request"
	"api/src/repository/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		fmt.Println(userRequest)
		return
	}
	u := user.Create(userRequest)
	c.JSON(http.StatusOK, u)
}
