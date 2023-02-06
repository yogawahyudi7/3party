package helpers

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	Config "pinang-mikro-3party/config"
	Constants "pinang-mikro-3party/constants"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type CurlSubmitSIKPAkadParams struct {
	Nik             string
	Rekeninglama    string
	Rekeningbaru    string
	Statusakad      string
	Statusrekening  string
	Nomorakad       string
	Tglakad         string
	Tgljatuhtempo   string
	Nilaiakad       string
	Kodepenjamin    string
	Nomorpenjaminan string
	Nilaidijamin    string
	Skema           string
	Sektor          string
	Negaratujuan    string
}

type CurlSubmitSIKPAkadResponse struct {
	Status       int
	Message      string
	MessageLocal string
}

type CurlSubmitSIKPAkadMapping struct {
	XMLName xml.Name           `xml:"Envelope"`
	Body    BodySubmitSIKPAkad `xml:"Body"`
}

type BodySubmitSIKPAkad struct {
	XMLName  xml.Name               `xml:"Body"`
	Fault    FaultSubmitSIKPAkad    `xml:"Fault"`
	Response ResponseSubmitSIKPAkad `xml:"submit_SIKP_AkadResponse"`
}

type FaultSubmitSIKPAkad struct {
	XMLName     xml.Name `xml:"Fault"`
	FaultCode   string   `xml:"faultcode"`
	FaultString string   `xml:"faultstring"`
}

type ResponseSubmitSIKPAkad struct {
	XMLName xml.Name             `xml:"submit_SIKP_AkadResponse"`
	Result  ResultSubmitSIKPAkad `xml:"submit_SIKP_AkadResult"`
}

type ResultSubmitSIKPAkad struct {
	XMLName xml.Name `xml:"submit_SIKP_AkadResult"`
	Error   bool     `xml:"error"`
	Code    string   `xml:"code"`
	Message string   `xml:"message"`
}

func CurlSubmitSIKPAkad(params CurlSubmitSIKPAkadParams) (response CurlSubmitSIKPAkadResponse, data CurlSubmitSIKPAkadMapping) {
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

	log.Println("--ENDPOINT API Curl Submit AKAD SIKP--")
	log.Println(endpoint)

	transCfg := Config.TransportConfig

	client := &http.Client{
		Timeout:   time.Duration(25 * time.Second),
		Transport: transCfg,
	}

	xmlnSchemeXsi := `xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" `
	xmlnSchemeXsd := `xmlns:xsd="http://www.w3.org/2001/XMLSchema" `
	xmlnSoap := `xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"`

	Nik := params.Nik
	Rekeninglama := params.Rekeninglama
	Rekeningbaru := params.Rekeningbaru
	Statusakad := params.Statusakad
	Statusrekening := params.Statusrekening
	Nomorakad := params.Nomorakad
	Tglakad := params.Tglakad
	Tgljatuhtempo := params.Tgljatuhtempo
	Nilaiakad := params.Nilaiakad
	Kodepenjamin := params.Kodepenjamin
	Nomorpenjaminan := params.Nomorpenjaminan
	Nilaidijamin := params.Nilaidijamin
	Skema := params.Skema
	Sektor := params.Sektor
	Negaratujuan := params.Negaratujuan

	submit := `
	  <submit_SIKP_Akad xmlns="http://tempuri.org/">
      <k_akad>
        <NIK>` + Nik + `</NIK>
        <REKENING_LAMA>` + Rekeninglama + `</REKENING_LAMA>
        <REKENING_BARU>` + Rekeningbaru + `</REKENING_BARU>
        <STATUS_AKAD>` + Statusakad + `</STATUS_AKAD>
        <STATUS_REKENING>` + Statusrekening + `</STATUS_REKENING>
        <NOMOR_AKAD>` + Nomorakad + `</NOMOR_AKAD>
        <TGL_AKAD>` + Tglakad + `</TGL_AKAD>
        <TGL_JATUH_TEMPO>` + Tgljatuhtempo + `</TGL_JATUH_TEMPO>
        <NILAI_AKAD>` + Nilaiakad + `</NILAI_AKAD>
        <KODE_PENJAMIN>` + Kodepenjamin + `</KODE_PENJAMIN>
        <NOMOR_PENJAMINAN>` + Nomorpenjaminan + `</NOMOR_PENJAMINAN>
        <NILAI_DIJAMIN>` + Nilaidijamin + `</NILAI_DIJAMIN>
        <SKEMA>` + Skema + `</SKEMA>
        <SEKTOR>` + Sektor + `</SEKTOR>
        <NEGARA_TUJUAN>` + Negaratujuan + `</NEGARA_TUJUAN>
      </k_akad>
    </submit_SIKP_Akad>
	`
	body := `<soap:Envelope ` + xmlnSchemeXsi + xmlnSchemeXsd + xmlnSoap + `><soap:Body>` + submit + `</soap:Body></soap:Envelope>`

	fmt.Println(body)
	r, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(body))) // URL-encoded payload

	//ERROR REQUEST
	if err != nil {
		log.Println("Helper -- ERROR REQUEST CURL Curl Submit AKAD SIKP ---")
		log.Println("Helper -- ERROR REQUEST CURL Curl Submit AKAD SIKP : ", err.Error())

		joinString := []string{"Status Code : ", "-", " | Error : ", err.Error()}
		erorrMessage := strings.Join(joinString, "")

		response := CurlSubmitSIKPAkadResponse{
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
		log.Println("Helper -- ERROR HIT API CURL Curl Submit AKAD SIKP ---")
		log.Println("Helper -- ERROR HIT API CURL Curl Submit AKAD SIKP : ", err.Error())

		joinString := []string{"Status Code : ", "-", " | Error : ", err.Error()}
		erorrMessage := strings.Join(joinString, "")

		response := CurlSubmitSIKPAkadResponse{
			Status:       500,
			Message:      Constants.ErorrGeneralMessage,
			MessageLocal: erorrMessage,
		}

		return response, data
	}

	defer curlResponse.Body.Close()

	log.Println("Helper -- RESPONSE STATUS CURL Curl Submit AKAD SIKP : ", curlResponse.Status)
	log.Println("Helper -- RESPONSE HEADERS CURL Curl Submit AKAD SIKP : ", curlResponse.Header)
	log.Println("Helper -- REQUEST URL CURL Curl Submit AKAD SIKP : ", curlResponse.Request.URL)
	log.Println("Helper -- REQUEST CONTENT LENGTH CURL Curl Submit AKAD SIKP : ", curlResponse.Request.ContentLength)

	xml.NewDecoder(curlResponse.Body).Decode(&data)
	log.Println("RESULT DATA :", data)

	//ERROR STATUS CODE
	if curlResponse.StatusCode != 200 {
		// log.Println("Helper --- ERROR STATUS CODE CURL Curl Submit AKAD SIKP ---")
		// log.Println("Helper --- ERROR STATUS CODE CURL Curl Submit AKAD SIKP : ", curlResponse.StatusCode)

		joinString := []string{"Status Code : ", "Web Service Error", " | Error : ", data.Body.Fault.FaultString}
		erorrMessage := strings.Join(joinString, "")

		response := CurlSubmitSIKPAkadResponse{
			Status:       500,
			Message:      Constants.ErorrGeneralMessage,
			MessageLocal: erorrMessage,
		}

		return response, data
	}

	// log.Println("Helper --- DATA DECODE CURL Curl Submit AKAD SIKP ---")
	// log.Println(data)

	response = CurlSubmitSIKPAkadResponse{
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
