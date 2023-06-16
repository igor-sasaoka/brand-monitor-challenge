package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userRepository "github.com/igor-sasaoka/brand-monitor-front/repository/user"
)

func Auth() gin.HandlerFunc {
    return func(c *gin.Context) {
        s, err := c.Cookie("session")
        if err != nil {
            c.Redirect(http.StatusFound, "/login")
            return
        }
        _, err = userRepository.GetUserBySession(s)
        if err != nil {
            c.Redirect(http.StatusFound, "/login")
            return
        }
    }
}
