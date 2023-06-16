package main

import (
	"github.com/igor-sasaoka/brand-monitor-front/database"
	"github.com/igor-sasaoka/brand-monitor-front/server"
)

//MVC like implementation that handles user registration and auth session.
//Once authenticated, used will have access to /app page and will be able to send commands
//to an external backend (backend service) and display the response.
func main() {
    database.Init()
    server.Init()
}
