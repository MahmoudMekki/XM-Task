package server

import (
	"github.com/MahmoudMekki/XM-Task/router"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	engine := gin.Default()
	routerInterface := router.NewRouter(engine)
	engine = routerInterface.SetRouter()
	return engine
}
