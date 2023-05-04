package user

import (
	"api/src/model"
	"api/src/model/request"
	repo "api/src/repository/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func PostUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Println("Error PostUser: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.UserRequestToUser(userRequest)

	userRepository, err := repo.NewRepositoryUser()
	response, err := userRepository.Create(*user)
	if err != nil {
		log.Println("Error PostUser: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": response})
}
