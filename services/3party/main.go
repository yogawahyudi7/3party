package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	Controller "pinang-mikro-3party/delivery/controller"
	pb "pinang-mikro-3party/delivery/proto/3party"

	Config "pinang-mikro-3party/config"
)

func init() {
	// local active
	// err := godotenv.Load("services/3party/service-dev.env")
	err := godotenv.Load("services/3party/service-staging.env")
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
	t := time.Now()
	formatDate := t.Format("20060102")
	logJoin := []string{"logs", "/", "log", "-", formatDate, ".log"}
	logFile := strings.Join(logJoin, "")
	_, err := os.Stat(logFile)

	//check exist file log
	f, _ := os.OpenFile(logFile, os.O_RDWR|os.O_APPEND, 0755)
	if os.IsNotExist(err) {
		f, _ = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)

	}
	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	log.SetOutput(gin.DefaultWriter)

	port, _ := Config.PortService()

	log.Println("--- START SERVICE V1 --- ", "PORT", port)
	// fmt.Println("--- START SERVICE V1 --- ", "PORT", port)

	lis, err := net.Listen("tcp", port)
	// log.Println("Failed to listen : ", err)
	if err != nil {
		log.Println("Failed to listen : ", err.Error())
		// fmt.Println("Failed to listen : ", err.Error())
	}

	grpcServer := grpc.NewServer(
		grpc.MaxRecvMsgSize(1024*1024*20),
		grpc.MaxSendMsgSize(1024*1024*20),
	)

	pb.RegisterThirdPartyServiceServer(grpcServer, &Controller.Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Println("Failed to server : ", err)
		// fmt.Println("Failed to server : ", err)
	}

}
