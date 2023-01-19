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

type CurlVerifySIKPParams struct {
	NIK string
}

type CurlVerifySIKPResponse struct {
	Status       int
	Message      string
	MessageLocal string
}

type CurlVerifySIKPMapping struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	XMLName  xml.Name `xml:"Body"`
	Response Response `xml:"get_SIKP_Calon_satuanResponse"`
}

type Response struct {
	XMLName xml.Name `xml:"get_SIKP_Calon_satuanResponse"`
	Result  Result   `xml:"get_SIKP_Calon_satuanResult"`
}

type Result struct {
	XMLName xml.Name `xml:"get_SIKP_Calon_satuanResult"`
	Error   bool     `xml:"error"`
	Code    string   `xml:"code"`
	Message string   `xml:"message"`
	Data    Data     `xml:"data"`
}

type Data struct {
	XMLName    xml.Name `xml:"data"`
	KodeBank   string   `xml:"kode_bank"`
	UploadDate string   `xml:"tgl_upload"`
}

func CurlVerifySIKP(params CurlVerifySIKPParams) (response CurlVerifySIKPResponse, data CurlVerifySIKPMapping) {
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

	log.Println("--ENDPOINT API SIKP VERIFY--")
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
	get := `<get_SIKP_Calon_satuan xmlns="http://tempuri.org/"><id>` + nik + `</id></get_SIKP_Calon_satuan>`
	body := `<soap:Envelope ` + xmlnSchemeXsi + xmlnSchemeXsd + xmlnSoap + `><soap:Body>` + get + `</soap:Body></soap:Envelope>`

	fmt.Println(body)
	r, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(body))) // URL-encoded payload

	//ERROR REQUEST
	if err != nil {
		log.Println("Helper -- ERROR REQUEST CURL SIKP VERIFY ---")
		log.Println("Helper -- ERROR REQUEST CURL SIKP VERIFY : ", err.Error())

		joinString := []string{"Status Code : ", "-", " | Error : ", err.Error()}
		erorrMessage := strings.Join(joinString, "")

		response := CurlVerifySIKPResponse{
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
		log.Println("Helper -- ERROR HIT API CURL SIKP VERIFY ---")
		log.Println("Helper -- ERROR HIT API CURL SIKP VERIFY : ", err.Error())

		joinString := []string{"Status Code : ", "-", " | Error : ", err.Error()}
		erorrMessage := strings.Join(joinString, "")

		response := CurlVerifySIKPResponse{
			Status:       500,
			Message:      Constants.ErorrGeneralMessage,
			MessageLocal: erorrMessage,
		}

		return response, data
	}

	defer curlResponse.Body.Close()

	log.Println("Helper -- RESPONSE STATUS CURL SIKP VERIFY : ", curlResponse.Status)
	log.Println("Helper -- RESPONSE HEADERS CURL SIKP VERIFY : ", curlResponse.Header)
	log.Println("Helper -- REQUEST URL CURL SIKP VERIFY : ", curlResponse.Request.URL)
	log.Println("Helper -- REQUEST CONTENT LENGTH CURL SIKP VERIFY : ", curlResponse.Request.ContentLength)

	//ERROR STATUS CODE
	if curlResponse.StatusCode != 200 {
		// log.Println("Helper --- ERROR STATUS CODE CURL SIKP VERIFY ---")
		// log.Println("Helper --- ERROR STATUS CODE CURL SIKP VERIFY : ", curlResponse.StatusCode)

		joinString := []string{"Status Code : ", "-", " | Error : ", err.Error()}
		erorrMessage := strings.Join(joinString, "")

		response := CurlVerifySIKPResponse{
			Status:       500,
			Message:      Constants.ErorrGeneralMessage,
			MessageLocal: erorrMessage,
		}

		return response, data
	}

	xml.NewDecoder(curlResponse.Body).Decode(&data)

	// log.Println("Helper --- DATA DECODE CURL SIKP VERIFY ---")
	// log.Println(data)

	response = CurlVerifySIKPResponse{
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
