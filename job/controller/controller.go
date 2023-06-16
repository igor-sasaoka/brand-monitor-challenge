package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igor-sasaoka/brand-monitor-job/model"
	"github.com/igor-sasaoka/brand-monitor-job/service"
)


func RegisterAccount(c *gin.Context) {
    var a model.Account
    if err := c.ShouldBindJSON(&a); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := service.GetAccountRegistry().Add(a); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        c.Abort()
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Account created successfully"})
}
