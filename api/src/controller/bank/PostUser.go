package bank

import (
	"api/src/model/request"
	"fmt"

	"github.com/gin-gonic/gin"
)

func PostUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		fmt.Println(userRequest)
		return
	}
	fmt.Println(userRequest)
}
