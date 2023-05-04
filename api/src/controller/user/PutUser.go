package user

import (
	"api/src/model"
	"api/src/model/request"
	repo "api/src/repository/user"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func PutUserBy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println("Error converting id to int: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		return
	}

	userRepository, err := repo.NewRepositoryUser()
	if err != nil {
		log.Println("Error NewRepositoryUser: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	_, err = userRepository.GetUserById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}
	}

	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Println("Error PutUser: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}
	userRequest.ID = id

	userUpdate, err := userRepository.Update(*model.UserRequestToUser(userRequest))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNoContent, gin.H{"error": "not up"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
	}

	response := model.UserToUserRequest(*userUpdate)
	c.JSON(http.StatusOK, gin.H{"data": response})
}
