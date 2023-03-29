package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTime(c *gin.Context) {
	currentTime := time.Now()
	c.IndentedJSON(http.StatusOK, currentTime)
}
