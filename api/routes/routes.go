package routes

import (
	"github.com/diogocarasco/ginapi/api/handlers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", handlers.HomeHandler)

	return router
}
