package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"

	_ "gopkg.in/goracle.v2"
)

var adega, _ = InitAdega()

var f, _ = os.OpenFile("log_"+time.Now().Format("01-02-2006")+".log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)

func main() {
	// Load env variables
	handleError(godotenv.Load())

	// Logging to a file.
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	defer f.Close()

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":" + os.Getenv("port"))

}

func setupRouter() *gin.Engine {
	/* Setup for router */
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r.PUT("adega/vinho/:vinho", PutVinho)
	r.POST("adega/vinho/", PostVinho)
	r.POST("adega/vinhos/", PostVinhos)
	r.PUT("adega/vinho/:vinho/garrafas/:quantidade", UpdateQtd)
	r.GET("adega/vinho/:vinho", GetVinho)
	r.GET("adega", GetAdega)
	r.DELETE("adega/vinho/:vinho", DelVinho)

	return r
}
