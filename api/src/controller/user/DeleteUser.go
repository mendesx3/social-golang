package user

import (
	repo "api/src/repository/user"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func DeleteUserBy(c *gin.Context) {
	var id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("Error converting id to int: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		return
	}

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}
	}
	userRepository, err := repo.NewRepositoryUser()
	if err != nil {
		log.Println("Error NewRepositoryUser: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	err = userRepository.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
