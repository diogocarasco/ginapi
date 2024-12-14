package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Root path request OK!"})
}
