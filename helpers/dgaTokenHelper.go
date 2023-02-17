package helpers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	Config "pinang-mikro-3party/config"
	Constants "pinang-mikro-3party/constants"

	"github.com/gin-gonic/gin"
)

type DGATokenParams struct {
	Username string
	Password string
}

type DGATokenResponse struct {
	Status       int
	Message      string
	MessageLocal string
}

type DGATokenMapping struct {
	Code    int    `json:"code"`
	Expire  string `json:"expire"`
	Token   string `json:"token"`
	Message string `json:"message"`
}

func DGARequestToken(params DGATokenParams) (response DGATokenResponse, data DGATokenMapping) {
	t := time.Now()
	formatDate := t.Format("20060102")
	logJoin := []string{"logs", "/", "log", "-", formatDate, ".log"}
	logFile := strings.Join(logJoin, "")
	_, err := os.Stat(logFile)

	file, _ := os.OpenFile(logFile, os.O_RDWR|os.O_APPEND, 0775)
	if os.IsNotExist(err) {
		file, _ = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	}

	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	log.SetOutput(gin.DefaultWriter)

	paramsCurl := url.Values{}
	paramsCurl.Set("username", params.Username)
	paramsCurl.Set("password", params.Password)

	log.Println("Helper -- PARAMS CURL DGA REQEST TOKEN-")
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
	pathEndpoint := Constants.EndpointRequestToken
	queryStringJoin := []string{firstEndpoint, pathEndpoint}
	endpoint := strings.Join(queryStringJoin, "")

	log.Println("--ENDPOINT CURL DGA REQEST TOKEN--")
	log.Println(endpoint)

	transfCfg := Config.TransportConfig

	client := &http.Client{
		Transport: transfCfg,
		Timeout:   time.Duration(25 * time.Second),
	}

	r, err := http.NewRequest("POST", endpoint, strings.NewReader(paramsCurl.Encode()))

	// ERROR REQUEST
	if err != nil {
		log.Println("Helper --- ERROR REQUEST CURL DGA REQEST TOKEN---")
		log.Println("Helper --- ERROR REQUEST CURL DGA REQEST TOKEN: ", err.Error())

		joinString := []string{"Status Code : ", "-", " | Error : ", err.Error()}
		errorMessage := strings.Join(joinString, "")

		response := DGATokenResponse{
			Status:       500,
			Message:      Constants.ErorrGeneralMessage,
			MessageLocal: errorMessage,
		}

		return response, data
	}

	r.Header.Set(Constants.CURLHeaderContentType, Constants.CURLHeaderContentTypeValue)
	r.Header.Set(Constants.CURLHeaderCacheControl, Constants.CURLHeaderCacheControlValue)

	curlResponse, err := client.Do(r)

	// ERROR CURL
	if err != nil {
		log.Println("Helper -- ERROR HIT API DGA REQUEST TOKEN ---")
		log.Println("Helper -- ERROR HIT API DGA REQUEST TOKEN : ", err.Error())

		joinString := []string{"Status Code :", "-", "| Error : ", err.Error()}
		errorMessage := strings.Join(joinString, "")

		response := DGATokenResponse{
			Status:       500,
			Message:      Constants.ErorrGeneralMessage,
			MessageLocal: errorMessage,
		}

		return response, data
	}

	defer curlResponse.Body.Close()

	log.Println("Helper -- RESPONSE STATUS CURL DGA REQUEST TOKEN : ", curlResponse.Status)
	log.Println("Helper -- RESPONSE HEADERS CURL DGA REQUEST TOKEN : ", curlResponse.Request.Header)
	log.Println("Helper -- REQUEST URL CURL DGA REQUEST TOKEN : ", curlResponse.Request.URL)
	log.Println("Helper -- REQUEST CONTENT LENGTH CURL DGA REQUEST TOKEN : ", curlResponse.Request.ContentLength)

	json.NewDecoder(curlResponse.Body).Decode(&data)
	log.Println("RESULT DATA :", data)

	//ERROR STATUS CODE
	if curlResponse.StatusCode != 200 {
		// log.Println("Helper --- ERROR STATUS CODE CURL SIKP VERIFY ---")
		// log.Println("Helper --- ERROR STATUS CODE CURL SIKP VERIFY : ", curlResponse.StatusCode)

		joinString := []string{"Status Code : ", curlResponse.Status, " | Error : ", data.Message}
		erorrMessage := strings.Join(joinString, "")

		response := DGATokenResponse{
			Status:       500,
			Message:      Constants.ErorrGeneralMessage,
			MessageLocal: erorrMessage,
		}

		return response, data
	}

	log.Println("HELPER =========== DATA DECODE CURL DGA REQUEST TOKEN  ===================")

	response = DGATokenResponse{
		Status:       200,
		Message:      "get dga request token berhasil",
		MessageLocal: "Success get dga request token",
	}

	return response, data
}
