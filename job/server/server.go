package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/igor-sasaoka/brand-monitor-job/controller"
)

func Init() {
    r := gin.Default()
    r.POST("/register", controller.RegisterAccount)

    log.Fatal(r.Run())
}
