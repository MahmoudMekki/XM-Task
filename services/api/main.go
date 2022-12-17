package main

import (
	"context"
	"fmt"
	"github.com/MahmoudMekki/XM-Task/config"
	"github.com/MahmoudMekki/XM-Task/database"
	"github.com/MahmoudMekki/XM-Task/migration"
	"github.com/MahmoudMekki/XM-Task/router"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	err := database.CreateDBConnection()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	migration.RunMigration()
}
func main() {
	gin.SetMode(config.GetEnvVar("GIN_MODE"))
	engine := gin.Default()
	routerInterface := router.NewRouter(engine)
	engine = routerInterface.SetRouter()
	srv := &http.Server{Addr: config.GetEnvVar("WEB_SERVER_PORT"), Handler: engine}
	log.Info().Msg(fmt.Sprintf("web server is running on localhost%s", config.GetEnvVar("WEB_SERVER_PORT")))
	go func() {
		log.Err(srv.ListenAndServe())
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	<-sigChan
	log.Info().Msg("Received a terminate signal, Gracefully shutdown the server")
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(tc)
}
