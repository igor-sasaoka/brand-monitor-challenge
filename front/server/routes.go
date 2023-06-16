package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igor-sasaoka/brand-monitor-front/controller"
	"github.com/igor-sasaoka/brand-monitor-front/middleware"
)

func Routes() *gin.Engine {
    r := gin.Default()

    r.Static("/dist", "./dist")  
    r.LoadHTMLGlob("templates/*")

    r.GET("/login", func(c *gin.Context) {
        c.HTML(http.StatusOK, "login.html", gin.H{})
    })
    r.GET("/register", func(c *gin.Context) {
        c.HTML(http.StatusOK, "register.html", gin.H{})
    })

    r.POST("/register", controller.Register)
    r.POST("/login", controller.Login)

    r.GET("/", func(c *gin.Context) {
        c.Redirect(http.StatusFound, "/app")
    })

    //authenticated routes
    rAuth := r.Group("/app", middleware.Auth())
    rAuth.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{})
    })

    return r
}

