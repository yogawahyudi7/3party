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

func (s *Server) submitJamkrindoCalon(ctx context.Context, request *pb.SubmitJamkrindoCalonRequest) (*pb.SubmitJamkrindoCalonResponse, error) {
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

	response := &pb.SubmitJamkrindoCalonResponse{
		Status:                        0,
		Message:                       "-",
		MessageLocal:                  "-",
		EmbedDatasubmitJamkrindoCalon: nil,
	}

	typeLog := ""
	//REQUEST PARAMETER
	AlamatDebitur := request.GetAlamatDebitur()
	AlamatUsaha := request.GetAlamatUsaha()
	CabangRekanan := request.GetCabangRekanan()
	Cif := request.GetCif()
	Coverage := request.GetCoverage()
	FlagTransfer := request.GetFlagTransfer()
	IDCalonDebiturKur := request.GetIDCalonDebiturKur()
	JangkaWaktu := request.GetJangkaWaktu()
	JenisAgunan := request.GetJenisAgunan()
	JenisIdentitas := request.GetJenisIdentitas()
	JenisKelamin := request.GetJenisKelamin()
	JenisKredit := request.GetJenisKredit()
	JenisKur := request.GetJenisKur()
	JenisLinkage := request.GetJenisLinkage()
	JenisPengikatan := request.GetJenisPengikatan()
	JmlTKerja := request.GetJmlTKerja()
	KodeBank := request.GetKodeBank()
	KodePekerjaan := request.GetKodePekerjaan()
	KodePos := request.GetKodePos()
	KodeSektor := request.GetKodeSektor()
	KodeUker := request.GetKodeUker()
	LembagaLinkage := request.GetLembagaLinkage()
	ModalUsaha := request.GetModalUsaha()
	NamaDebitur := request.GetNamaDebitur()
	NilaiAgunan := request.GetNilaiAgunan()
	NoHp := request.GetNoHp()
	NoIdentitas := request.GetNoIdentitas()
	NoIjinUsaha := request.GetNoIjinUsaha()
	NoPk := request.GetNoPk()
	NoRekening := request.GetNoRekening()
	NoTelepon := request.GetNoTelepon()
	NomorAplikasi := request.GetNomorAplikasi()
	Npwp := request.GetNpwp()
	PlafonKredit := request.GetPlafonKredit()
	StatusAplikasi := request.GetStatusAplikasi()
	StatusKolektibilitas := request.GetStatusKolektibilitas()
	StatusLunas := request.GetStatusLunas()
	SukuBunga := request.GetSukuBunga()
	TanggalAkhir := request.GetTanggalAkhir()
	TanggalAwal := request.GetTanggalAwal()
	TanggalLahir := request.GetTanggalLahir()
	TanggalMulaiUsaha := request.GetTanggalMulaiUsaha()
	TanggalPk := request.GetTanggalPk()
	TanggalRekam := request.GetTanggalRekam()
	UsahaProduktif := request.GetUsahaProduktif()
	//REQUEST FOR LOGGING
	requestData := map[string]interface{}{
		"AlamatDebitur":        AlamatDebitur,
		"AlamatUsaha":          AlamatUsaha,
		"CabangRekanan":        CabangRekanan,
		"Cif":                  Cif,
		"Coverage":             Coverage,
		"FlagTransfer":         FlagTransfer,
		"IDCalonDebiturKur":    IDCalonDebiturKur,
		"JangkaWaktu":          JangkaWaktu,
		"JenisAgunan":          JenisAgunan,
		"JenisIdentitas":       JenisIdentitas,
		"JenisKelamin":         JenisKelamin,
		"JenisKredit":          JenisKredit,
		"JenisKur":             JenisKur,
		"JenisLinkage":         JenisLinkage,
		"JenisPengikatan":      JenisPengikatan,
		"JmlTKerja":            JmlTKerja,
		"KodeBank":             KodeBank,
		"KodePekerjaan":        KodePekerjaan,
		"KodePos":              KodePos,
		"KodeSektor":           KodeSektor,
		"KodeUker":             KodeUker,
		"LembagaLinkage":       LembagaLinkage,
		"ModalUsaha":           ModalUsaha,
		"NamaDebitur":          NamaDebitur,
		"NilaiAgunan":          NilaiAgunan,
		"NoHp":                 NoHp,
		"NoIdentitas":          NoIdentitas,
		"NoIjinUsaha":          NoIjinUsaha,
		"NoPk":                 NoPk,
		"NoRekening":           NoRekening,
		"NoTelepon":            NoTelepon,
		"NomorAplikasi":        NomorAplikasi,
		"Npwp":                 Npwp,
		"PlafonKredit":         PlafonKredit,
		"StatusAplikasi":       StatusAplikasi,
		"StatusKolektibilitas": StatusKolektibilitas,
		"StatusLunas":          StatusLunas,
		"SukuBunga":            SukuBunga,
		"TanggalAkhir":         TanggalAkhir,
		"TanggalAwal":          TanggalAwal,
		"TanggalLahir":         TanggalLahir,
		"TanggalMulaiUsaha":    TanggalMulaiUsaha,
		"TanggalPk":            TanggalPk,
		"TanggalRekam":         TanggalRekam,
		"UsahaProduktif":       UsahaProduktif,
	}
	dataRequest, _ := json.Marshal(requestData)
	jsonDataRequest := string(dataRequest)

	//DECODE PARAMETER

	//CURL SUBMIT JAMKRINDO CALON
	typeLog = "curl-submit-jamkrindo-calon"

	params := Helpers.CurlSubmitJamkrindoCalonParams{
		AlamatDebitur:        AlamatDebitur,
		AlamatUsaha:          AlamatUsaha,
		CabangRekanan:        CabangRekanan,
		Cif:                  Cif,
		Coverage:             Coverage,
		FlagTransfer:         FlagTransfer,
		IDCalonDebiturKur:    IDCalonDebiturKur,
		JangkaWaktu:          JangkaWaktu,
		JenisAgunan:          JenisAgunan,
		JenisIdentitas:       JenisIdentitas,
		JenisKelamin:         JenisKelamin,
		JenisKredit:          JenisKredit,
		JenisKur:             JenisKur,
		JenisLinkage:         JenisLinkage,
		JenisPengikatan:      JenisPengikatan,
		JmlTKerja:            JmlTKerja,
		KodeBank:             KodeBank,
		KodePekerjaan:        KodePekerjaan,
		KodePos:              KodePos,
		KodeSektor:           KodeSektor,
		KodeUker:             KodeUker,
		LembagaLinkage:       LembagaLinkage,
		ModalUsaha:           ModalUsaha,
		NamaDebitur:          NamaDebitur,
		NilaiAgunan:          NilaiAgunan,
		NoHp:                 NoHp,
		NoIdentitas:          NoIdentitas,
		NoIjinUsaha:          NoIjinUsaha,
		NoPk:                 NoPk,
		NoRekening:           NoRekening,
		NoTelepon:            NoTelepon,
		NomorAplikasi:        NomorAplikasi,
		Npwp:                 Npwp,
		PlafonKredit:         PlafonKredit,
		StatusAplikasi:       StatusAplikasi,
		StatusKolektibilitas: StatusKolektibilitas,
		StatusLunas:          StatusLunas,
		SukuBunga:            SukuBunga,
		TanggalAkhir:         TanggalAkhir,
		TanggalAwal:          TanggalAwal,
		TanggalLahir:         TanggalLahir,
		TanggalMulaiUsaha:    TanggalMulaiUsaha,
		TanggalPk:            TanggalPk,
		TanggalRekam:         TanggalRekam,
		UsahaProduktif:       UsahaProduktif,
	}

	var responseResult Helpers.CurlSubmitJamkrindoCalonResponse
	var dataResult Helpers.CurlSubmitJamkrindoCalonMapping
	responseResult, dataResult = Helpers.CurlSubmitJamkrindoCalon(params)

	log.Println("Service ** RESULT QUERY PLAFOND SIKP **")
	log.Println(responseResult)
	log.Println(dataResult)

	//response success
	if responseResult.Status != 200 {
		response = &pb.SubmitJamkrindoCalonResponse{
			Status:                        int64(responseResult.Status),
			Message:                       responseResult.Message,
			MessageLocal:                  responseResult.MessageLocal,
			EmbedDatasubmitJamkrindoCalon: nil,
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
				UserId:       "",
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

	var resultData []*pb.DatasubmitJamkrindoCalon
	for _, vData := range dataResult.Body.Response.Result.Data.DataJamkrindo {

		noIjinPrinsip := vData.NoIjinPrinsip
		noSertifikat := vData.NoSertifikat
		NomorUrut := vData.NomorUrut
		tglIjinPrinsip := vData.TglIjinPrinsip
		tglSertifikat := vData.TglSertifikat

		data := pb.DatasubmitJamkrindoCalon{
			NoIjinPrinsip:  noIjinPrinsip,
			NoSertifikat:   noSertifikat,
			NomorUrut:      NomorUrut,
			TglIjinPrinsip: tglIjinPrinsip,
			TglSertifikat:  tglSertifikat,
		}

		resultData = append(resultData, &data)

	}

	response = &pb.SubmitJamkrindoCalonResponse{
		Status:       200,
		Message:      responseResult.Message,
		MessageLocal: responseResult.MessageLocal,
		EmbedDatasubmitJamkrindoCalon: &pb.EmbedDatasubmitJamkrindoCalon{
			StatusCode:               int64(intStatusCode),
			StatusDescription:        dataResult.Body.Response.Result.Message,
			DatasubmitJamkrindoCalon: resultData,
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
			UserId:       "",
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
