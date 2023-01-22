package helpers

import (
	Config "3party/config"
	Constants "3party/constants"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type CurlPlafondSIKPParams struct {
	NIK string
}

type CurlPlafondSIKPResponse struct {
	Status       int
	Message      string
	MessageLocal string
}

type CurlPlafondSIKPMapping struct {
	XMLName xml.Name        `xml:"Envelope"`
	Body    BodyPlafondSIKP `xml:"Body"`
}

type BodyPlafondSIKP struct {
	XMLName  xml.Name            `xml:"Body"`
	Response ResponsePlafondSIKP `xml:"get_SIKP_PlafonResponse"`
}

type ResponsePlafondSIKP struct {
	XMLName xml.Name          `xml:"get_SIKP_PlafonResponse"`
	Result  ResultPlafondSIKP `xml:"get_SIKP_PlafonResult"`
}

type ResultPlafondSIKP struct {
	XMLName xml.Name        `xml:"get_SIKP_PlafonResult"`
	Error   bool            `xml:"error"`
	Code    string          `xml:"code"`
	Message string          `xml:"message"`
	Data    DataPlafondSIKP `xml:"data"`
}

type DataPlafondSIKP struct {
	XMLName     xml.Name           `xml:"data"`
	DataPlafond []DataPlafondsSIKP `xml:"data_plafon"`
}

type DataPlafondsSIKP struct {
	XMLName           xml.Name `xml:"data_plafon"`
	Nik               string   `xml:"nik"`
	Skema             string   `xml:"skema"`
	TotalLimitDefault string   `xml:"total_limit_default"`
	TotalLimit        string   `xml:"total_limit"`
	LimitAktifDefault string   `xml:"limit_aktif_default"`
	LimitAktif        string   `xml:"limit_aktif"`
	KodeBank          string   `xml:"kode_bank"`
}

func CurlPlafondSIKP(params CurlPlafondSIKPParams) (response CurlPlafondSIKPResponse, data CurlPlafondSIKPMapping) {
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

	firstEndpoint, _ := Config.SIKPWebService()
	queryStringJoin := []string{firstEndpoint, "/SIKP_Service.asmx"}
	endpoint := strings.Join(queryStringJoin, "")

	log.Println("--ENDPOINT API PLAFOND SIKP--")
	log.Println(endpoint)

	transCfg := Config.TransportConfig

	client := &http.Client{
		Timeout:   time.Duration(25 * time.Second),
		Transport: transCfg,
	}

	xmlnSchemeXsi := `xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" `
	xmlnSchemeXsd := `xmlns:xsd="http://www.w3.org/2001/XMLSchema" `
	xmlnSoap := `xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"`
	nik := params.NIK
	get := `<get_SIKP_Plafon xmlns="http://tempuri.org/"><id>` + nik + `</id></get_SIKP_Plafon>`
	body := `<soap:Envelope ` + xmlnSchemeXsi + xmlnSchemeXsd + xmlnSoap + `><soap:Body>` + get + `</soap:Body></soap:Envelope>`

	fmt.Println(body)
	r, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(body))) // URL-encoded payload

	//ERROR REQUEST
	if err != nil {
		log.Println("Helper -- ERROR REQUEST CURL PLAFOND SIKP ---")
		log.Println("Helper -- ERROR REQUEST CURL PLAFOND SIKP : ", err.Error())

		joinString := []string{"Status Code : ", "-", " | Error : ", err.Error()}
		erorrMessage := strings.Join(joinString, "")

		response := CurlPlafondSIKPResponse{
			Status:       500,
			Message:      Constants.ErorrGeneralMessage,
			MessageLocal: erorrMessage,
		}

		return response, data
	}

	r.Header.Set(Constants.CURLHeaderContentType, Constants.CURLHeaderContentTypeValueTextXML)
	r.Header.Set(Constants.CURLHeaderCacheControl, Constants.CURLHeaderCacheControlValue)
	// r.SetBasicAuth(usernameSSO, passwordSSO)
	curlResponse, err := client.Do(r)

	//ERROR CURL
	if err != nil {
		log.Println("Helper -- ERROR HIT API CURL PLAFOND SIKP ---")
		log.Println("Helper -- ERROR HIT API CURL PLAFOND SIKP : ", err.Error())

		joinString := []string{"Status Code : ", "-", " | Error : ", err.Error()}
		erorrMessage := strings.Join(joinString, "")

		response := CurlPlafondSIKPResponse{
			Status:       500,
			Message:      Constants.ErorrGeneralMessage,
			MessageLocal: erorrMessage,
		}

		return response, data
	}

	defer curlResponse.Body.Close()

	log.Println("Helper -- RESPONSE STATUS CURL PLAFOND SIKP : ", curlResponse.Status)
	log.Println("Helper -- RESPONSE HEADERS CURL PLAFOND SIKP : ", curlResponse.Header)
	log.Println("Helper -- REQUEST URL CURL PLAFOND SIKP : ", curlResponse.Request.URL)
	log.Println("Helper -- REQUEST CONTENT LENGTH CURL PLAFOND SIKP : ", curlResponse.Request.ContentLength)

	//ERROR STATUS CODE
	if curlResponse.StatusCode != 200 {
		// log.Println("Helper --- ERROR STATUS CODE CURL PLAFOND SIKP ---")
		// log.Println("Helper --- ERROR STATUS CODE CURL PLAFOND SIKP : ", curlResponse.StatusCode)

		joinString := []string{"Status Code : ", "-", " | Error : ", err.Error()}
		erorrMessage := strings.Join(joinString, "")

		response := CurlPlafondSIKPResponse{
			Status:       500,
			Message:      Constants.ErorrGeneralMessage,
			MessageLocal: erorrMessage,
		}

		return response, data
	}

	xml.NewDecoder(curlResponse.Body).Decode(&data)

	// log.Println("Helper --- DATA DECODE CURL PLAFOND SIKP ---")
	// log.Println(data)

	response = CurlPlafondSIKPResponse{
		Status:       200,
		Message:      "Response Success",
		MessageLocal: "Response 200 Success",
	}

	//DUMMY FILE
	// data.Body.Response.Result.Code = "12"
	// data.Body.Response.Result.Message = "Data ditemukan"
	// data.Body.Response.Result.Data.KodeBank = "494"

	//RETURN
	return response, data
}
