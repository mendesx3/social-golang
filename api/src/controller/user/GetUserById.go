package user

import (
	"api/src/model"
	repo "api/src/repository/user"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err != nil {
		log.Println("Error InitMySql: ", err.Error())
		return
	}

	user, err := repo.NewRepositoryUser()
	if err != nil {
		log.Println("Error NewRepositoryUser: ", err.Error())
		return
	}
	response, err := user.GetUserById(id)
	if err != nil {
		return
	}

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
	}

	r := model.UserToUserRequest(*response)
	c.JSON(http.StatusOK, gin.H{"data": r})
}
