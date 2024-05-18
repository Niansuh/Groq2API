package main

import (
	"groqai2api/global"
	"groqai2api/initialize"
)

func main() {
	// Initial configuration
	initialize.InitConfig()
	// Initialize cache
	initialize.InitCache()
	// Initialize agent
	initialize.InitProxy()
	// Initialize account
	initialize.InitAuth()
	// Initialize routing
	Router := initialize.InitRouter()
	if err := Router.Run(global.Host + ":" + global.Port); err != nil {
		panic(err)
	}
}
