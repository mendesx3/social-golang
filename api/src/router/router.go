package router

import (
	"api/src/controller/bank"
	"api/src/controller/user"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	API := r.Group("/api/v1")
	initUserRouter(API)
	initAccountBankRouter(API)
}

func initUserRouter(API *gin.RouterGroup) {
	u := API.Group("/user")
	u.POST("/", user.PostUser)
	u.PUT("/:id", user.PutUserBy)
	u.GET("/", user.GetAllUser)
	u.GET("/:id", user.GetUserBy)
	u.DELETE("/:id", user.DeleteUserBy)
}

func initAccountBankRouter(API *gin.RouterGroup) {
	b := API.Group("/bank")
	b.POST("/", bank.PostUser)
	b.PUT("/:id", bank.PutUserBy)
	b.GET("/", bank.GetAllUser)
	b.GET("/:id", bank.GetUserBy)
	b.DELETE("/:id", bank.DeleteUserBy)
}
