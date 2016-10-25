package server

import (
	"../config"
	"github.com/gin-gonic/gin"
)

func Init() {
	config := config.GetConfig()
	gin.SetMode(config.GetString("GIN_MODE"))
	r := NewRouter()
	r.Run(config.GetString("SERVER_PORT"))
}
