package helpers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	Config "pinang-mikro-3party/config"
	Constants "pinang-mikro-3party/constants"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type CurlSICDLogsParams struct {
	NIK   string
	Token string
}

type CurlSICDLogsResponse struct {
	Status       int
	Message      string
	MessageLocal string
}

type CurlSICDLogsMapping struct {
	Status       int                     `json:"status"`
	Message      string                  `json:"message"`
	MessageLocal string                  `json:"desc"`
	Data         []DataEmbedCurlSICDLogs `json:"data"`
}

type DataEmbedCurlSICDLogs struct {
	Id         int    `json:"id"`
	PrakarsaId int    `json:"prakarsaId"`
	KtpNumber  string `json:"ktpNumber"`
	Reason     string `json:"reason"`
	Status     int    `json:"status"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

func CurlSICDLogs(params CurlSICDLogsParams) (response CurlSICDLogsResponse, data CurlSICDLogsMapping) {
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

	var paramsCurl = url.Values{}
	paramsCurl.Set("ktpNumber", params.NIK)

	log.Println("Helper -- PARAMS CURL SICD LOGS-")
	log.Println(paramsCurl)

	//JSON DATA REQUEST
	dataRequest, err := json.Marshal(paramsCurl)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Helper -- ** DATA REQUEST**")
	log.Println(string(dataRequest))

	firstEndpoint, _ := Config.APIBEUrl()
	pathEndpoint := Constants.EndpointSICDLogs
	queryStringJoin := []string{firstEndpoint, pathEndpoint}
	endpoint := strings.Join(queryStringJoin, "")

	log.Println("--ENDPOINT CURL SICD LOGS--")
	log.Println(endpoint)

	transCfg := Config.TransportConfig

	client := &http.Client{
		Timeout:   time.Duration(25 * time.Second),
		Transport: transCfg,
	}

	r, err := http.NewRequest("POST", endpoint, strings.NewReader(paramsCurl.Encode())) // URL-encoded payload

	//ERROR REQUEST
	if err != nil {
		log.Println("Helper -- ERROR REQUEST CURL SICD LOGS ---")
		log.Println("Helper -- ERROR REQUEST CURL SICD LOGS : ", err.Error())

		joinString := []string{"Status Code : ", "-", " | Error : ", err.Error()}
		erorrMessage := strings.Join(joinString, "")

		response := CurlSICDLogsResponse{
			Status:       500,
			Message:      Constants.ErorrGeneralMessage,
			MessageLocal: erorrMessage,
		}

		return response, data
	}

	//GET DIGIAGRI TOKEN

	dgaToken := "Bearer " + params.Token
	r.Header.Set(Constants.CURLHeaderAuthorization, dgaToken)
	r.Header.Set(Constants.CURLHeaderContentType, Constants.CURLHeaderContentTypeValue)
	r.Header.Set(Constants.CURLHeaderCacheControl, Constants.CURLHeaderCacheControlValue)

	curlResponse, err := client.Do(r)

	//ERROR CURL
	if err != nil {
		log.Println("Helper --- ERROR HIT API CURL SIKP VERIFY ---")
		log.Println("Helper -- ERROR HIT API CURL SIKP VERIFY : ", err.Error())

		joinString := []string{"Status Code : ", "-", " | Error : ", err.Error()}
		erorrMessage := strings.Join(joinString, "")

		response := CurlSICDLogsResponse{
			Status:       500,
			Message:      Constants.ErorrGeneralMessage,
			MessageLocal: erorrMessage,
		}

		return response, data
	}

	defer curlResponse.Body.Close()

	log.Println("Helper -- RESPONSE STATUS CURL SIKP VERIFY : ", curlResponse.Status)
	log.Println("Helper -- RESPONSE STATUS CURL SIKP VERIFY : ", curlResponse.StatusCode)
	log.Println("Helper -- RESPONSE HEADERS CURL SIKP VERIFY : ", curlResponse.Header)
	log.Println("Helper -- REQUEST URL CURL SIKP VERIFY : ", curlResponse.Request.URL)
	log.Println("Helper -- REQUEST CONTENT LENGTH CURL SIKP VERIFY : ", curlResponse.Request.ContentLength)

	json.NewDecoder(curlResponse.Body).Decode(&data)
	log.Println("RESULT DATA :", data)

	//ERROR STATUS CODE
	if curlResponse.StatusCode != 200 {
		// log.Println("Helper --- ERROR STATUS CODE CURL SIKP VERIFY ---")
		// log.Println("Helper --- ERROR STATUS CODE CURL SIKP VERIFY : ", curlResponse.StatusCode)

		joinString := []string{"Status Code : ", curlResponse.Status, " | Error : ", data.MessageLocal}
		erorrMessage := strings.Join(joinString, "")

		response := CurlSICDLogsResponse{
			Status:       500,
			Message:      Constants.ErorrGeneralMessage,
			MessageLocal: erorrMessage,
		}

		return response, data
	}

	// log.Println("Helper --- DATA DECODE CURL SIKP VERIFY ---")
	// log.Println(data)

	response = CurlSICDLogsResponse{
		Status:       200,
		Message:      "Response Success",
		MessageLocal: "Response 200 Success",
	}

	//RETURN
	return response, data
}
