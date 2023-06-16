package main

import (
	"github.com/igor-sasaoka/brand-monitor-job/config"
	"github.com/igor-sasaoka/brand-monitor-job/server"
	"github.com/igor-sasaoka/brand-monitor-job/service"
)

//Will start an endpoint that receives and register an email.
//When registered, a welcome message will be sent to te registered email.
//Check request.sh to see curl example.
func main(){
    config.Init()    
    service.SpawnListeners()
    server.Init()
}
