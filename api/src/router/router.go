package router

import (
	"api/controller/user"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	u := r.Group("/user")
	u.POST("/", user.PostUser)
	u.PUT("/:id", user.PutUserBy)
	u.GET("/", user.GetAllUser)
	u.GET("/:id", user.GetUserBy)
	u.DELETE("/:id", user.DeleteUserBy)
}
