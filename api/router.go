package api

import (
	"github.com/gin-gonic/gin"
	"groqai2api/initialize"
	"net/http"
)

var router *gin.Engine

func init() {
	// Initial configuration
	initialize.InitConfig()
	// Initialize cache
	initialize.InitCache()
	// Initialize agent
	initialize.InitProxy()
	// Initialize account
	initialize.InitAuth()
	// Initialize gin
	router = initialize.InitRouter()
}

func Listen(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
