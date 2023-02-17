package controller

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	Constants "pinang-mikro-3party/constants"
	pb "pinang-mikro-3party/delivery/proto/3party"
	Helpers "pinang-mikro-3party/helpers"

	Config "pinang-mikro-3party/config"
	Models "pinang-mikro-3party/delivery/Models"
)

type Server struct {
	pb.ThirdPartyServiceServer
}

func (s *Server) Testing(ctx context.Context, request *pb.TestingRequest) (*pb.TestingResponse, error) {
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

func (s *Server) VerificationSIKP(ctx context.Context, request *pb.VerificationSIKPRequest) (*pb.VerificationSIKPReponse, error) {
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
	userIdEncode := request.GetUserId()
	ktpNumberEncode := request.GetKtpNumber()

	//REQUEST FOR LOGGING
	requestData := map[string]interface{}{
		"userId":    userIdEncode,
		"ktpNumber": ktpNumberEncode,
	}
	dataRequest, _ := json.Marshal(requestData)
	jsonDataRequest := string(dataRequest)

	//DECODE PARAMETER
	userId := ""
	ktpNumber := ""
	{
		//USERID
		{
			typeLog = "userId-decode"
			result, statusCode := Helpers.DecodeStringBase64(userIdEncode)
			if statusCode != 200 {
				response = &pb.VerificationSIKPReponse{
					Status:                    400,
					Message:                   "Maaf, Encode Parameter userId tidak valid.",
					MessageLocal:              "Encode Parameter userId tidak valid.",
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
						Endpoint:     Constants.EndpointSIKPVerification,
						UserId:       userId,
						ActionDate:   time.Now().Format(Constants.FullLayoutTime),
						Description:  Constants.Desc3PartyLogging,
						DataRequest:  string(dataRequest),
						DataResponse: string(dataResponse),
					}

					logData, _ := json.Marshal(dataLog)
					jsonDataLog := string(logData)

					Helpers.PubLoggingCloud(jsonDataRequest, jsonDataResponse, jsonDataLog)

				}
				return response, nil
			}
			userId = result
		}

		//NIK
		{
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
						Endpoint:     Constants.EndpointSIKPVerification,
						UserId:       userId,
						ActionDate:   time.Now().Format(Constants.FullLayoutTime),
						Description:  Constants.Desc3PartyLogging,
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
	}

	//VALIDATION PARAMETER
	{
		// USER ID
		{
			typeLog = "userId-validation"

			ktpValidateStatus, ktpValidateMessage, ktpValidateMessageLocal := Helpers.ValidatorUserId(userId)
			if ktpValidateStatus != 200 {
				response = &pb.VerificationSIKPReponse{
					Status:                    int64(ktpValidateStatus),
					Message:                   ktpValidateMessage,
					MessageLocal:              ktpValidateMessageLocal,
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
						Endpoint:     Constants.EndpointSIKPVerification,
						UserId:       userId,
						ActionDate:   time.Now().Format(Constants.FullLayoutTime),
						Description:  Constants.Desc3PartyLogging,
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

		// KTP NUMBER
		{
			typeLog = "ktpNumber-validation"

			ktpValidateStatus, ktpValidateMessage, ktpValidateMessageLocal := Helpers.ValidatorKtpNumber(ktpNumber)
			if ktpValidateStatus != 200 {
				response = &pb.VerificationSIKPReponse{
					Status:                    int64(ktpValidateStatus),
					Message:                   ktpValidateMessage,
					MessageLocal:              ktpValidateMessageLocal,
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
						Endpoint:     Constants.EndpointSIKPVerification,
						UserId:       userId,
						ActionDate:   time.Now().Format(Constants.FullLayoutTime),
						Description:  Constants.Desc3PartyLogging,
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
	}

	//CURL FERIVICATION SIKP
	typeLog = "curl-verification-sikp"

	params := Helpers.CurlVerificationSIKPParams{
		NIK: ktpNumber,
	}

	// log.Println("DECODE USER ID :", userId)
	// log.Println("DECODE KTP NUMBER :", params, ktpNumber)

	var responseResult Helpers.CurlVerificationSIKPResponse
	var dataResult Helpers.CurlVerificationSIKPMapping
	responseResult, dataResult = Helpers.CurlVerificationSIKP(params)

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
				Endpoint:     Constants.EndpointSIKPVerification,
				UserId:       userId,
				ActionDate:   time.Now().Format(Constants.FullLayoutTime),
				Description:  Constants.Desc3PartyLogging,
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
			DataVerificationSIKP: &pb.DataVerificationSIKP{
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
			Endpoint:     Constants.EndpointSIKPVerification,
			UserId:       userId,
			ActionDate:   time.Now().Format(Constants.FullLayoutTime),
			Description:  Constants.Desc3PartyLogging,
			DataRequest:  string(dataRequest),
			DataResponse: string(dataResponse),
		}

		logData, _ := json.Marshal(dataLog)
		jsonDataLog := string(logData)

		Helpers.PubLoggingCloud(jsonDataRequest, jsonDataResponse, jsonDataLog)

	}
	return response, nil
}

func (s *Server) CheckPlafondSIKP(ctx context.Context, request *pb.CheckPlafondSIKPRequest) (*pb.CheckPlafondSIKPReponse, error) {
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

	response := &pb.CheckPlafondSIKPReponse{
		Status:                    0,
		Message:                   "-",
		MessageLocal:              "-",
		EmbedDataCheckPlafondSIKP: nil,
	}

	typeLog := ""
	//REQUEST PARAMETER
	userIdEncode := request.GetUserId()
	ktpNumberEncode := request.GetKtpNumber()

	//REQUEST FOR LOGGING
	requestData := map[string]interface{}{
		"userId":    userIdEncode,
		"ktpNumber": ktpNumberEncode,
	}
	dataRequest, _ := json.Marshal(requestData)
	jsonDataRequest := string(dataRequest)

	//DECODE PARAMETER
	userId := ""
	ktpNumber := ""
	{
		//USERID
		{
			typeLog = "ktpNumber-decode"
			result, statusCode := Helpers.DecodeStringBase64(userIdEncode)
			if statusCode != 200 {
				response = &pb.CheckPlafondSIKPReponse{
					Status:                    400,
					Message:                   "Maaf, Encode Parameter userId tidak valid.",
					MessageLocal:              "Encode Parameter userId tidak valid.",
					EmbedDataCheckPlafondSIKP: nil,
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
						Endpoint:     Constants.EndpointSIKPCheckPlafond,
						UserId:       userId,
						ActionDate:   time.Now().Format(Constants.FullLayoutTime),
						Description:  Constants.Desc3PartyLogging,
						DataRequest:  string(dataRequest),
						DataResponse: string(dataResponse),
					}

					logData, _ := json.Marshal(dataLog)
					jsonDataLog := string(logData)

					Helpers.PubLoggingCloud(jsonDataRequest, jsonDataResponse, jsonDataLog)

				}
				return response, nil
			}
			userId = result
		}

		//NIK
		{
			typeLog = "ktpNumber-decode"
			result, statusCode := Helpers.DecodeStringBase64(ktpNumberEncode)
			if statusCode != 200 {
				response = &pb.CheckPlafondSIKPReponse{
					Status:                    400,
					Message:                   "Maaf, Encode Parameter ktpNumber tidak valid.",
					MessageLocal:              "Encode Parameter ktpNumber tidak valid.",
					EmbedDataCheckPlafondSIKP: nil,
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
						Endpoint:     Constants.EndpointSIKPCheckPlafond,
						UserId:       userId,
						ActionDate:   time.Now().Format(Constants.FullLayoutTime),
						Description:  Constants.Desc3PartyLogging,
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
	}

	//VALIDATION PARAMETER
	{
		// USER ID
		{
			typeLog = "userId-validation"

			ktpValidateStatus, ktpValidateMessage, ktpValidateMessageLocal := Helpers.ValidatorUserId(userId)
			if ktpValidateStatus != 200 {
				response = &pb.CheckPlafondSIKPReponse{
					Status:                    int64(ktpValidateStatus),
					Message:                   ktpValidateMessage,
					MessageLocal:              ktpValidateMessageLocal,
					EmbedDataCheckPlafondSIKP: nil,
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
						Endpoint:     Constants.EndpointSIKPCheckPlafond,
						UserId:       userId,
						ActionDate:   time.Now().Format(Constants.FullLayoutTime),
						Description:  Constants.Desc3PartyLogging,
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

		// KTP NUMBER
		{
			typeLog = "ktpNumber-validation"

			ktpValidateStatus, ktpValidateMessage, ktpValidateMessageLocal := Helpers.ValidatorKtpNumber(ktpNumber)
			if ktpValidateStatus != 200 {
				response = &pb.CheckPlafondSIKPReponse{
					Status:                    int64(ktpValidateStatus),
					Message:                   ktpValidateMessage,
					MessageLocal:              ktpValidateMessageLocal,
					EmbedDataCheckPlafondSIKP: nil,
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
						Endpoint:     Constants.EndpointSIKPCheckPlafond,
						UserId:       userId,
						ActionDate:   time.Now().Format(Constants.FullLayoutTime),
						Description:  Constants.Desc3PartyLogging,
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
	}

	//CURL FERIVICATION SIKP
	typeLog = "curl-verification-sikp"

	params := Helpers.CurlPlafondSIKPParams{
		NIK: ktpNumber,
	}

	var responseResult Helpers.CurlPlafondSIKPResponse
	var dataResult Helpers.CurlPlafondSIKPMapping
	responseResult, dataResult = Helpers.CurlPlafondSIKP(params)

	log.Println("Service ** RESULT QUERY PLAFOND SIKP **")
	log.Println(responseResult)
	log.Println(dataResult)

	//response success
	if responseResult.Status != 200 {
		response = &pb.CheckPlafondSIKPReponse{
			Status:                    int64(responseResult.Status),
			Message:                   responseResult.Message,
			MessageLocal:              responseResult.MessageLocal,
			EmbedDataCheckPlafondSIKP: nil,
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
				Endpoint:     Constants.EndpointSIKPCheckPlafond,
				UserId:       userId,
				ActionDate:   time.Now().Format(Constants.FullLayoutTime),
				Description:  Constants.Desc3PartyLogging,
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
	typeLog = "success-check-plafond-sikp"

	intStatusCode, _ := Helpers.StringInt(dataResult.Body.Response.Result.Code)

	var resultData []*pb.DataCheckPlafondSIKP
	for _, vData := range dataResult.Body.Response.Result.Data.DataPlafond {

		scheme, _ := Helpers.StringInt(vData.Skema)
		totalLimitDefault, _ := Helpers.StringInt(vData.TotalLimitDefault)
		totalLimit, _ := Helpers.StringInt(vData.TotalLimit)
		limitActiveDefault, _ := Helpers.StringInt(vData.LimitAktifDefault)
		limitActive, _ := Helpers.StringInt(vData.LimitAktif)
		bankCode, _ := Helpers.StringInt(vData.KodeBank)

		//DECODE RESPONSE
		ktpNumberEncode, _ := Helpers.EncodeStringBase64(vData.Nik)

		data := pb.DataCheckPlafondSIKP{
			KtpNumber:          ktpNumberEncode,
			Scheme:             int64(scheme),
			TotalLimitDefault:  int64(totalLimitDefault),
			TotalLimit:         int64(totalLimit),
			LimitActiveDefault: int64(limitActiveDefault),
			LimitActive:        int64(limitActive),
			BankCode:           int64(bankCode),
		}

		resultData = append(resultData, &data)

	}

	response = &pb.CheckPlafondSIKPReponse{
		Status:       200,
		Message:      responseResult.Message,
		MessageLocal: responseResult.MessageLocal,
		EmbedDataCheckPlafondSIKP: &pb.EmbedDataCheckPlafondSIKP{
			StatusCode:           int64(intStatusCode),
			StatusDescription:    dataResult.Body.Response.Result.Message,
			DataCheckPlafondSIKP: resultData,
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
			Endpoint:     Constants.EndpointSIKPCheckPlafond,
			UserId:       userId,
			ActionDate:   time.Now().Format(Constants.FullLayoutTime),
			Description:  Constants.Desc3PartyLogging,
			DataRequest:  string(dataRequest),
			DataResponse: string(dataResponse),
		}

		logData, _ := json.Marshal(dataLog)
		jsonDataLog := string(logData)

		Helpers.PubLoggingCloud(jsonDataRequest, jsonDataResponse, jsonDataLog)

	}
	return response, nil
}

func (s *Server) VerificationSICD(ctx context.Context, request *pb.VerificationSICDRequest) (*pb.VerificationSICDReponse, error) {
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

	response := &pb.VerificationSICDReponse{
		Status:                    0,
		Message:                   "-",
		MessageLocal:              "-",
		EmbedDataVerificationSICD: nil,
	}

	typeLog := ""
	//REQUEST PARAMETER
	userIdEncode := request.GetUserId()
	ktpNumberEncode := request.GetKtpNumber()

	//REQUEST FOR LOGGING
	requestData := map[string]interface{}{
		"userId":    userIdEncode,
		"ktpNumber": ktpNumberEncode,
	}
	dataRequest, _ := json.Marshal(requestData)
	jsonDataRequest := string(dataRequest)

	//DECODE PARAMETER
	userId := ""
	ktpNumber := ""
	{
		//USERID
		{
			typeLog = "userId-decode"
			result, statusCode := Helpers.DecodeStringBase64(userIdEncode)
			if statusCode != 200 {
				response = &pb.VerificationSICDReponse{
					Status:                    400,
					Message:                   "Maaf, Encode Parameter userId tidak valid.",
					MessageLocal:              "Encode Parameter userId tidak valid.",
					EmbedDataVerificationSICD: nil,
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
						Endpoint:     Constants.EndpointSIKPVerification,
						UserId:       userId,
						ActionDate:   time.Now().Format(Constants.FullLayoutTime),
						Description:  Constants.Desc3PartyLogging,
						DataRequest:  string(dataRequest),
						DataResponse: string(dataResponse),
					}

					logData, _ := json.Marshal(dataLog)
					jsonDataLog := string(logData)

					Helpers.PubLoggingCloud(jsonDataRequest, jsonDataResponse, jsonDataLog)

				}
				return response, nil
			}
			userId = result
		}

		//NIK
		{
			typeLog = "ktpNumber-decode"
			result, statusCode := Helpers.DecodeStringBase64(ktpNumberEncode)
			if statusCode != 200 {
				response = &pb.VerificationSICDReponse{
					Status:                    400,
					Message:                   "Maaf, Encode Parameter ktpNumber tidak valid.",
					MessageLocal:              "Encode Parameter ktpNumber tidak valid.",
					EmbedDataVerificationSICD: nil,
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
						Endpoint:     Constants.EndpointSIKPVerification,
						UserId:       userId,
						ActionDate:   time.Now().Format(Constants.FullLayoutTime),
						Description:  Constants.Desc3PartyLogging,
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
	}

	//VALIDATION PARAMETER
	{
		// USER ID
		{
			typeLog = "userId-validation"

			ktpValidateStatus, ktpValidateMessage, ktpValidateMessageLocal := Helpers.ValidatorUserId(userId)
			if ktpValidateStatus != 200 {
				response = &pb.VerificationSICDReponse{
					Status:                    int64(ktpValidateStatus),
					Message:                   ktpValidateMessage,
					MessageLocal:              ktpValidateMessageLocal,
					EmbedDataVerificationSICD: nil,
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
						Endpoint:     Constants.EndpointSIKPVerification,
						UserId:       userId,
						ActionDate:   time.Now().Format(Constants.FullLayoutTime),
						Description:  Constants.Desc3PartyLogging,
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

		// KTP NUMBER
		{
			typeLog = "ktpNumber-validation"

			ktpValidateStatus, ktpValidateMessage, ktpValidateMessageLocal := Helpers.ValidatorKtpNumber(ktpNumber)
			if ktpValidateStatus != 200 {
				response = &pb.VerificationSICDReponse{
					Status:                    int64(ktpValidateStatus),
					Message:                   ktpValidateMessage,
					MessageLocal:              ktpValidateMessageLocal,
					EmbedDataVerificationSICD: nil,
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
						Endpoint:     Constants.EndpointSIKPVerification,
						UserId:       userId,
						ActionDate:   time.Now().Format(Constants.FullLayoutTime),
						Description:  Constants.Desc3PartyLogging,
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
	}

	//GET PARAM DGA BACKEND USERNAME
	usernameParams := Models.DGAParametersParams{
		Desc: Constants.BackendUsernameDGAParams,
	}

	usernameResponse, usernameData := Models.FindDGAParamters(usernameParams)

	if usernameResponse.Status != 200 {
		response = &pb.VerificationSICDReponse{
			Status:                    int64(usernameResponse.Status),
			Message:                   usernameResponse.Message,
			MessageLocal:              usernameResponse.MessageLocal,
			EmbedDataVerificationSICD: nil,
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
				Endpoint:     Constants.EndpointSIKPVerification,
				UserId:       userId,
				ActionDate:   time.Now().Format(Constants.FullLayoutTime),
				Description:  Constants.Desc3PartyLogging,
				DataRequest:  string(dataRequest),
				DataResponse: string(dataResponse),
			}

			logData, _ := json.Marshal(dataLog)
			jsonDataLog := string(logData)

			Helpers.PubLoggingCloud(jsonDataRequest, jsonDataResponse, jsonDataLog)

		}
		return response, nil

	}

	//GET PARAM DGA BACKEND PASSWORD
	passwordParams := Models.DGAParametersParams{
		Desc: Constants.BackendPasswordDGAParams,
	}

	passwordParamsResponse, passwordData := Models.FindDGAParamters(passwordParams)

	if passwordParamsResponse.Status != 200 {
		response = &pb.VerificationSICDReponse{
			Status:                    int64(passwordParamsResponse.Status),
			Message:                   passwordParamsResponse.Message,
			MessageLocal:              passwordParamsResponse.MessageLocal,
			EmbedDataVerificationSICD: nil,
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
				Endpoint:     Constants.EndpointSIKPVerification,
				UserId:       userId,
				ActionDate:   time.Now().Format(Constants.FullLayoutTime),
				Description:  Constants.Desc3PartyLogging,
				DataRequest:  string(dataRequest),
				DataResponse: string(dataResponse),
			}

			logData, _ := json.Marshal(dataLog)
			jsonDataLog := string(logData)

			Helpers.PubLoggingCloud(jsonDataRequest, jsonDataResponse, jsonDataLog)

		}
		return response, nil

	}

	backendUsername := usernameData.Value
	backendPassword := passwordData.Value

	//CURL DGA REQUEST TOKEN
	typeLog = "curl-dga-request-token"

	dgaRequestTokenParams := Helpers.DGATokenParams{
		Username: backendUsername,
		Password: backendPassword,
	}

	// log.Println("DECODE USER ID :", userId)
	// log.Println("DECODE KTP NUMBER :", params, ktpNumber)

	dgaRequestTokenResponse, dgaRequestTokenData := Helpers.DGARequestToken(dgaRequestTokenParams)

	log.Println("Service ** RESULT QUERY DGA REQUEST TOKEN **")
	log.Println(dgaRequestTokenResponse)
	log.Println(dgaRequestTokenData)

	//response success
	if dgaRequestTokenResponse.Status != 200 {
		response = &pb.VerificationSICDReponse{
			Status:                    int64(dgaRequestTokenResponse.Status),
			Message:                   dgaRequestTokenResponse.Message,
			MessageLocal:              dgaRequestTokenResponse.MessageLocal,
			EmbedDataVerificationSICD: nil,
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
				Endpoint:     Constants.EndpointSIKPVerification,
				UserId:       userId,
				ActionDate:   time.Now().Format(Constants.FullLayoutTime),
				Description:  Constants.Desc3PartyLogging,
				DataRequest:  string(dataRequest),
				DataResponse: string(dataResponse),
			}

			logData, _ := json.Marshal(dataLog)
			jsonDataLog := string(logData)

			Helpers.PubLoggingCloud(jsonDataRequest, jsonDataResponse, jsonDataLog)

		}
		return response, nil

	}

	//DGA TOKEN
	dgaToken := dgaRequestTokenData.Token

	//CURL SICD LOGS
	typeLog = "curl-sicd-logs"

	sicdLogsParams := Helpers.CurlSICDLogsParams{
		NIK:   ktpNumber,
		Token: dgaToken,
	}

	// log.Println("DECODE USER ID :", userId)
	// log.Println("DECODE KTP NUMBER :", params, ktpNumber)

	sicdLogsResponse, sicdLogsData := Helpers.CurlSICDLogs(sicdLogsParams)

	log.Println("Service ** RESULT QUERY DGA REQUEST TOKEN **")
	log.Println(sicdLogsResponse)
	log.Println(sicdLogsData)

	//response success
	if sicdLogsResponse.Status != 200 {
		response = &pb.VerificationSICDReponse{
			Status:                    int64(sicdLogsResponse.Status),
			Message:                   sicdLogsResponse.Message,
			MessageLocal:              sicdLogsResponse.MessageLocal,
			EmbedDataVerificationSICD: nil,
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
				Endpoint:     Constants.EndpointSIKPVerification,
				UserId:       userId,
				ActionDate:   time.Now().Format(Constants.FullLayoutTime),
				Description:  Constants.Desc3PartyLogging,
				DataRequest:  string(dataRequest),
				DataResponse: string(dataResponse),
			}

			logData, _ := json.Marshal(dataLog)
			jsonDataLog := string(logData)

			Helpers.PubLoggingCloud(jsonDataRequest, jsonDataResponse, jsonDataLog)

		}
		return response, nil

	}

	dataSicdLogs := sicdLogsData.Data
	log.Println("LEN DATA SICDLOGS :", len(dataSicdLogs))

	if len(dataSicdLogs) > 0 {
		//CHECK LAST CREATE SUBMISSON LOAN
		lastSubmission := sicdLogsData.Data[0].CreatedAt
		lastSubmissionString, _ := Helpers.DateFormat(lastSubmission, "-")
		log.Println("================", lastSubmissionString)
		lastSubmissionDate, _ := time.Parse(Constants.LayoutYMD, lastSubmissionString)
		log.Println("================", lastSubmissionDate)
		add1Month := lastSubmissionDate.AddDate(0, 1, 0)
		log.Println("================", add1Month)

		after1Month := t.After(add1Month)
		log.Println("================", after1Month)

		yesterdayDate := t.AddDate(0, 0, -1).String()
		yesterdayDateParse, _ := time.Parse(Constants.LayoutYMD, yesterdayDate[:10])
		yesterdayDateFormat := yesterdayDateParse.Format(Constants.LayoutYMD)
		yesterdayDateFormat = strings.ReplaceAll(yesterdayDateFormat, "-", "")
		log.Println("DATE :", yesterdayDate)
		log.Println("DATE :", yesterdayDate[:10])
		log.Println("DATE WITH FORMAT :", yesterdayDateFormat)
		// return response, nil
		//FIND REKENING PINANG
		rekeningPinangParams := Models.RekeningPinangParams{
			// ID: "99111554",
			KtpNumber: "3171051211950001",
		}

		rekeningPinangResponse, rekeningPinangData := Models.FindRekeningPinang(rekeningPinangParams)

		log.Println("Service ** RESULT QUERY REKENING PINANG FIND **")
		log.Println(rekeningPinangResponse)
		log.Println(rekeningPinangData)

		//response success
		if rekeningPinangResponse.Status != 200 {
			response = &pb.VerificationSICDReponse{
				Status:                    int64(rekeningPinangResponse.Status),
				Message:                   rekeningPinangResponse.Message,
				MessageLocal:              rekeningPinangResponse.MessageLocal,
				EmbedDataVerificationSICD: nil,
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
					Endpoint:     Constants.EndpointSIKPVerification,
					UserId:       userId,
					ActionDate:   time.Now().Format(Constants.FullLayoutTime),
					Description:  Constants.Desc3PartyLogging,
					DataRequest:  string(dataRequest),
					DataResponse: string(dataResponse),
				}

				logData, _ := json.Marshal(dataLog)
				jsonDataLog := string(logData)

				Helpers.PubLoggingCloud(jsonDataRequest, jsonDataResponse, jsonDataLog)

			}
			return response, nil

		}

		log.Println("---------------------------------")
		log.Println("AcctNo :", rekeningPinangData[0].Acctno)
		log.Println("CIF :", rekeningPinangData[0].CIF)
		log.Println("Type :", rekeningPinangData[0].Type)
		log.Println("NamaDebitur :", rekeningPinangData[0].NamaDebitur)
		log.Println("Status :", rekeningPinangData[0].Status)
		log.Println("Cbal :", rekeningPinangData[0].Cbal)
		log.Println("Bikole :", rekeningPinangData[0].Bikole)
		log.Println("KodeCabang :", rekeningPinangData[0].KodeCabang)
		log.Println("NoIdentitas :", rekeningPinangData[0].NoIdentitas)
		log.Println("TglLahir :", rekeningPinangData[0].TglLahir)
		log.Println("TglPembukaan :", rekeningPinangData[0].TglPembukaan)
		log.Println("TglMature :", rekeningPinangData[0].TglMature)

		return response, nil

		//////////////////////////
		if !after1Month {

			statusCode := 2
			statusDescription := "Pinjaman tidak dapat diproses, nasabah masih dalam tenggat waktu penolakan pinjaman sebelumnya selama 30 hari"

			dataVerificationSICD := []*pb.DataVerificationSICD{}

			for _, vData := range sicdLogsData.Data {
				data := &pb.DataVerificationSICD{
					Id:         int64(vData.Id),
					PrakarsaId: int64(vData.PrakarsaId),
					KtpNumber:  vData.KtpNumber,
					Reason:     vData.Reason,
					Status:     int64(vData.Status),
					CreatedAt:  vData.CreatedAt,
					UpdatedAt:  vData.UpdatedAt,
				}

				dataVerificationSICD = append(dataVerificationSICD, data)
			}

			resultData := &pb.EmbedDataVerificationSICD{
				StatusCode:           int64(statusCode),
				StatusDescription:    statusDescription,
				DataVerificationSICD: dataVerificationSICD,
			}

			response = &pb.VerificationSICDReponse{
				Status:                    int64(sicdLogsResponse.Status),
				Message:                   sicdLogsResponse.Message,
				MessageLocal:              sicdLogsResponse.MessageLocal,
				EmbedDataVerificationSICD: resultData,
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
					Endpoint:     Constants.EndpointSIKPVerification,
					UserId:       userId,
					ActionDate:   time.Now().Format(Constants.FullLayoutTime),
					Description:  Constants.Desc3PartyLogging,
					DataRequest:  string(dataRequest),
					DataResponse: string(dataResponse),
				}

				logData, _ := json.Marshal(dataLog)
				jsonDataLog := string(logData)

				Helpers.PubLoggingCloud(jsonDataRequest, jsonDataResponse, jsonDataLog)

			}
			return response, nil

		} else {
			// DataTable dtResult = dbu.GetDataPinang(NIK);

		}

	} else {

	}
	// log.Println("-------------------------------- :", sicdLogsData.Data[0].CreatedAt)

	return response, nil
	//response success
	typeLog = "success-verification-sicd"

	statusCode := 2
	statusDescription := "Pinjaman tidak dapat diproses, nasabah masih dalam tenggat waktu penolakan pinjaman sebelumnya selama 30 hari"

	dataVerificationSICD := []*pb.DataVerificationSICD{}
	for _, vData := range sicdLogsData.Data {

		data := &pb.DataVerificationSICD{
			Id:         int64(vData.Id),
			PrakarsaId: int64(vData.PrakarsaId),
			KtpNumber:  vData.KtpNumber,
			Reason:     vData.Reason,
			Status:     int64(vData.Status),
			CreatedAt:  vData.CreatedAt,
			UpdatedAt:  vData.UpdatedAt,
		}

		dataVerificationSICD = append(dataVerificationSICD, data)
	}

	resultData := &pb.EmbedDataVerificationSICD{
		StatusCode:           int64(statusCode),
		StatusDescription:    statusDescription,
		DataVerificationSICD: dataVerificationSICD,
	}

	response = &pb.VerificationSICDReponse{
		Status:                    int64(sicdLogsResponse.Status),
		Message:                   sicdLogsResponse.Message,
		MessageLocal:              sicdLogsResponse.MessageLocal,
		EmbedDataVerificationSICD: resultData,
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
			Endpoint:     Constants.EndpointSIKPVerification,
			UserId:       userId,
			ActionDate:   time.Now().Format(Constants.FullLayoutTime),
			Description:  Constants.Desc3PartyLogging,
			DataRequest:  string(dataRequest),
			DataResponse: string(dataResponse),
		}

		logData, _ := json.Marshal(dataLog)
		jsonDataLog := string(logData)

		Helpers.PubLoggingCloud(jsonDataRequest, jsonDataResponse, jsonDataLog)

	}
	return response, nil
}
