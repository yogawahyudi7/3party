package main

import (
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"

	Config "3party/config"
	Routes "3party/delivery/routes"
)

func init() {
	// local active
	err := godotenv.Load("clients/3party/client-dev.env")
	// err := godotenv.Load("clients/3party/client-staging.env")
	// err := godotenv.Load(".env")
	if err != nil {
		log.Println(err.Error())
		log.Fatal("Error loading .env file")
	}

	appEnv, _ := Config.AppEnv()
	appName, _ := Config.AppName()

	log.Println("APP_ENV : ", appEnv)
	log.Println("APP_NAME : ", appName)
	log.Println("-----------------------------------")
}

func main() {
	//SETUP PORT
	portMain, _ := Config.PortClient()

	serverMain := &http.Server{
		Addr:           portMain,
		Handler:        Routes.RouterMain(),
		ReadTimeout:    360 * time.Second,
		WriteTimeout:   360 * time.Second,
		MaxHeaderBytes: 1 << 60,
	}

	log.Println("RUN SERVER IN ", portMain, " HIT : ", time.Now())

	serverMain.ListenAndServe() //running on port 9008
}
