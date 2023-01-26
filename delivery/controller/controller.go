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

	"3party/constants"
	pb "3party/delivery/proto/3party"
	Helpers "3party/helpers"

	Config "3party/config"
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
						Endpoint:     constants.EndpointSIKPCheckPlafond,
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
						Endpoint:     constants.EndpointSIKPCheckPlafond,
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
						Endpoint:     constants.EndpointSIKPCheckPlafond,
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
						Endpoint:     constants.EndpointSIKPCheckPlafond,
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
				Endpoint:     constants.EndpointSIKPCheckPlafond,
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

		data := pb.DataCheckPlafondSIKP{
			KtpNumber:          vData.Nik,
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
			Endpoint:     constants.EndpointSIKPCheckPlafond,
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
