package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"3party/constants"
	pb "3party/delivery/proto/3party"

	Config "3party/config"
	Helpers "3party/helpers"
)

type server struct {
	pb.ThirdPartyServiceServer
}

// var validate *validator.Validate

func init() {
	// local active
	err := godotenv.Load("services/3party/.env")
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

func (s *server) Testing(ctx context.Context, request *pb.TestingRequest) (*pb.TestingResponse, error) {
	t := time.Now()
	formatDate := t.Format("20060102")
	logJoin := []string{"logs", "/", "services", "/", "3party", "/", "log", "-", formatDate, ".log"}
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

	log.Println("Testing : ", request.GetText())

	return &pb.TestingResponse{Result: "Welcome To 3PARTY Pinang Performa"}, nil
}

func (s *server) VerificationSIKP(ctx context.Context, request *pb.VerificationSIKPRequest) (*pb.VerificationSIKPReponse, error) {
	t := time.Now()
	formatDate := t.Format("20060102")
	logJoin := []string{"logs", "/", "services", "/", "3party", "/", "log", "-", formatDate, ".log"}
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

	response := &pb.VerificationSIKPReponse{
		Status:                    0,
		Message:                   "-",
		MessageLocal:              "-",
		EmbedDataVerificationSIKP: nil,
	}

	typeLog := ""
	//REQUEST PARAMETER
	userId := request.GetUserId()
	ktpNumberEncode := request.GetKtpNumber()

	//REQUEST FOR LOGGING
	requestData := map[string]interface{}{
		"userId":    userId,
		"ktpNumber": ktpNumberEncode,
	}
	dataRequest, _ := json.Marshal(requestData)
	jsonDataRequest := string(dataRequest)

	//DECODE PARAMETER
	ktpNumber := ""
	{
		//NIK
		typeLog = "ktpNumber-decode"
		result, statusCode := Helpers.DecodeStringBase64(ktpNumberEncode)
		if statusCode != 200 {
			response = &pb.VerificationSIKPReponse{
				Status:                    400,
				Message:                   "Maaf, Encode Parameter ktpNumber tidak valid.",
				MessageLocal:              "Encode Parameter ktpNumber tidak valid.",
				EmbedDataVerificationSIKP: nil,
			}

			//PROCESS TO LOGGING CLOUD
			{
				// DATA RESPONSE
				strStatusCode, _ := Helpers.IntString(400)
				responseData := map[string]interface{}{
					"statusCode":   strStatusCode,
					"responseData": response,
				}

				dataResponse, _ := json.Marshal(responseData)
				jsonDataResponse := string(dataResponse)

				dataLog := Config.LoggingCloudPubSub{
					Status:       "400",
					TypeLog:      typeLog,
					Endpoint:     constants.EndpointSIKPVerification,
					UserId:       userId,
					ActionDate:   time.Now().Format(constants.FullLayoutTime),
					Description:  constants.Desc3PartyLogging,
					DataRequest:  string(dataRequest),
					DataResponse: string(dataResponse),
				}

				logData, _ := json.Marshal(dataLog)
				jsonDataLog := string(logData)

				Helpers.PubLoggingCloud(jsonDataRequest, jsonDataResponse, jsonDataLog)

			}
			return response, nil
		}
		ktpNumber = result
	}

	//VALIDATION
	{
		//NIK
		typeLog = "ktpNumber-validation"

		if ktpNumber == "" {
			response = &pb.VerificationSIKPReponse{
				Status:                    400,
				Message:                   "Maaf, Parameter ktpNumber belum diisi.",
				MessageLocal:              "Parameter ktpNumber belum diisi.",
				EmbedDataVerificationSIKP: nil,
			}

			//PROCESS TO LOGGING CLOUD
			{
				// DATA RESPONSE
				strStatusCode, _ := Helpers.IntString(400)
				responseData := map[string]interface{}{
					"statusCode":   strStatusCode,
					"responseData": response,
				}

				dataResponse, _ := json.Marshal(responseData)
				jsonDataResponse := string(dataResponse)

				dataLog := Config.LoggingCloudPubSub{
					Status:       "400",
					TypeLog:      typeLog,
					Endpoint:     constants.EndpointSIKPVerification,
					UserId:       userId,
					ActionDate:   time.Now().Format(constants.FullLayoutTime),
					Description:  constants.Desc3PartyLogging,
					DataRequest:  string(dataRequest),
					DataResponse: string(dataResponse),
				}

				logData, _ := json.Marshal(dataLog)
				jsonDataLog := string(logData)

				Helpers.PubLoggingCloud(jsonDataRequest, jsonDataResponse, jsonDataLog)

			}
			return response, nil
		}
	}

	//CURL FERIVICATION SIKP
	typeLog = "curl-verification-sikp"

	params := Helpers.CurlVerifySIKPParams{
		NIK: ktpNumber,
	}

	var responseResult Helpers.CurlVerifySIKPResponse
	var dataResult Helpers.CurlVerifySIKPMapping
	responseResult, dataResult = Helpers.CurlVerifySIKP(params)

	log.Println("Service ** RESULT QUERY SIKP VERIFY **")
	log.Println(responseResult)
	log.Println(dataResult)

	//response success
	if responseResult.Status != 200 {
		response = &pb.VerificationSIKPReponse{
			Status:                    int64(responseResult.Status),
			Message:                   responseResult.Message,
			MessageLocal:              responseResult.MessageLocal,
			EmbedDataVerificationSIKP: &pb.EmbedDataVerificationSIKP{},
		}

		//PROCESS TO LOGGING CLOUD
		{
			// DATA RESPONSE
			strStatusCode, _ := Helpers.IntString(400)
			responseData := map[string]interface{}{
				"statusCode":   strStatusCode,
				"responseData": response,
			}

			dataResponse, _ := json.Marshal(responseData)
			jsonDataResponse := string(dataResponse)

			dataLog := Config.LoggingCloudPubSub{
				Status:       "400",
				TypeLog:      typeLog,
				Endpoint:     constants.EndpointSIKPVerification,
				UserId:       userId,
				ActionDate:   time.Now().Format(constants.FullLayoutTime),
				Description:  constants.Desc3PartyLogging,
				DataRequest:  string(dataRequest),
				DataResponse: string(dataResponse),
			}

			logData, _ := json.Marshal(dataLog)
			jsonDataLog := string(logData)

			Helpers.PubLoggingCloud(jsonDataRequest, jsonDataResponse, jsonDataLog)

		}
		return response, nil

	}

	//response success
	typeLog = "success-verification-sikp"

	intStatusCode, _ := Helpers.StringInt(dataResult.Body.Response.Result.Code)

	response = &pb.VerificationSIKPReponse{
		Status:       200,
		Message:      responseResult.Message,
		MessageLocal: responseResult.MessageLocal,
		EmbedDataVerificationSIKP: &pb.EmbedDataVerificationSIKP{
			StatusCode:        int64(intStatusCode),
			StatusDescription: dataResult.Body.Response.Result.Message,
			DataSIKP: &pb.DataSIKP{
				BankCode:   dataResult.Body.Response.Result.Data.KodeBank,
				UploadDate: dataResult.Body.Response.Result.Data.UploadDate,
			},
		},
	}

	//PROCESS TO LOGGING CLOUD
	{
		// DATA RESPONSE
		strStatusCode, _ := Helpers.IntString(400)
		responseData := map[string]interface{}{
			"statusCode":   strStatusCode,
			"responseData": response,
		}

		dataResponse, _ := json.Marshal(responseData)
		jsonDataResponse := string(dataResponse)

		dataLog := Config.LoggingCloudPubSub{
			Status:       "200",
			TypeLog:      typeLog,
			Endpoint:     constants.EndpointSIKPVerification,
			UserId:       userId,
			ActionDate:   time.Now().Format(constants.FullLayoutTime),
			Description:  constants.Desc3PartyLogging,
			DataRequest:  string(dataRequest),
			DataResponse: string(dataResponse),
		}

		logData, _ := json.Marshal(dataLog)
		jsonDataLog := string(logData)

		Helpers.PubLoggingCloud(jsonDataRequest, jsonDataResponse, jsonDataLog)

	}
	return response, nil
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

	pb.RegisterThirdPartyServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Println("Failed to server : ", err)
		// fmt.Println("Failed to server : ", err)
	}

}
