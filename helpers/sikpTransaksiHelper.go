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

type CurlSubmitSIKPTransaksiParams struct {
	NomorRekening  string
	TglTransaksi   string
	TglPelaporan   string
	Limit          string
	Outstanding    string
	AngsuranPokok  string
	Kolektibilitas string
}

type CurlSubmitSIKPTransaksiResponse struct {
	Status       int
	Message      string
	MessageLocal string
}

type CurlSubmitSIKPTransaksiMapping struct {
	XMLName xml.Name                `xml:"Envelope"`
	Body    BodySubmitSIKPTransaksi `xml:"Body"`
}

type BodySubmitSIKPTransaksi struct {
	XMLName  xml.Name                    `xml:"Body"`
	Fault    FaultSubmitSIKPTransaksi    `xml:"Fault"`
	Response ResponseSubmitSIKPTransaksi `xml:"submit_SIKP_TransaksiResponse"`
}

type FaultSubmitSIKPTransaksi struct {
	XMLName     xml.Name `xml:"Fault"`
	FaultCode   string   `xml:"faultcode"`
	FaultString string   `xml:"faultstring"`
}

type ResponseSubmitSIKPTransaksi struct {
	XMLName xml.Name                  `xml:"submit_SIKP_TransaksiResponse"`
	Result  ResultSubmitSIKPTransaksi `xml:"submit_SIKP_TransaksiResult"`
}

type ResultSubmitSIKPTransaksi struct {
	XMLName xml.Name                `xml:"submit_SIKP_TransaksiResult"`
	Error   bool                    `xml:"error"`
	Code    string                  `xml:"code"`
	Message string                  `xml:"message"`
	Data    DataSubmitSIKPTransaksi `xml:"data"`
}

type DataSubmitSIKPTransaksi struct {
	XMLName       xml.Name                      `xml:"data"`
	DataJamkrindo []SubmitSIKPTransaksiResponse `xml:"submit_SIKP_Transaksi"`
}

type SubmitSIKPTransaksiResponse struct {
	XMLName xml.Name `xml:"submit_SIKP_Transaksi"`
	Error   bool     `xml:"error"`
	Code    string   `xml:"code"`
	Message string   `xml:"message"`
}

func CurlSubmitSIKPTransaksi(params CurlSubmitSIKPTransaksiParams) (response CurlSubmitSIKPTransaksiResponse, data CurlSubmitSIKPTransaksiMapping) {
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

	log.Println("--ENDPOINT API Curl Submit SIKP Transaksi--")
	log.Println(endpoint)

	transCfg := Config.TransportConfig

	client := &http.Client{
		Timeout:   time.Duration(25 * time.Second),
		Transport: transCfg,
	}

	xmlnSchemeXsi := `xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" `
	xmlnSchemeXsd := `xmlns:xsd="http://www.w3.org/2001/XMLSchema" `
	xmlnSoap := `xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"`

	NomorRekening := params.NomorRekening
	TglTransaksi := params.TglTransaksi
	TglPelaporan := params.TglPelaporan
	Limit := params.Limit
	Outstanding := params.Outstanding
	AngsuranPokok := params.AngsuranPokok
	Kolektibilitas := params.Kolektibilitas

	submit := `
	<submit_SIKP_Transaksi xmlns="http://tempuri.org/">
	<k_Transaksi>
	  <NOMOR_REKENING>` + NomorRekening + `</NOMOR_REKENING>
	  <TGL_TRANSAKSI>` + TglTransaksi + `</TGL_TRANSAKSI>
	  <TGL_PELAPORAN>` + TglPelaporan + `</TGL_PELAPORAN>
	  <LIMIT>` + Limit + `</LIMIT>
	  <OUTSTANDING>` + Outstanding + `</OUTSTANDING>
	  <ANGSURAN_POKOK>` + AngsuranPokok + `</ANGSURAN_POKOK>
	  <KOLEKTIBILITAS>` + Kolektibilitas + `</KOLEKTIBILITAS>
	</k_Transaksi>
  </submit_SIKP_Transaksi>
	`
	body := `<soap:Envelope ` + xmlnSchemeXsi + xmlnSchemeXsd + xmlnSoap + `><soap:Body>` + submit + `</soap:Body></soap:Envelope>`

	fmt.Println(body)
	r, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(body))) // URL-encoded payload

	//ERROR REQUEST
	if err != nil {
		log.Println("Helper -- ERROR REQUEST CURL Curl Submit SIKP Transaksi ---")
		log.Println("Helper -- ERROR REQUEST CURL Curl Submit SIKP Transaksi : ", err.Error())

		joinString := []string{"Status Code : ", "-", " | Error : ", err.Error()}
		erorrMessage := strings.Join(joinString, "")

		response := CurlSubmitSIKPTransaksiResponse{
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
		log.Println("Helper -- ERROR HIT API CURL Curl Submit SIKP Transaksi ---")
		log.Println("Helper -- ERROR HIT API CURL Curl Submit SIKP Transaksi : ", err.Error())

		joinString := []string{"Status Code : ", "-", " | Error : ", err.Error()}
		erorrMessage := strings.Join(joinString, "")

		response := CurlSubmitSIKPTransaksiResponse{
			Status:       500,
			Message:      Constants.ErorrGeneralMessage,
			MessageLocal: erorrMessage,
		}

		return response, data
	}

	defer curlResponse.Body.Close()

	log.Println("Helper -- RESPONSE STATUS CURL Curl Submit SIKP Transaksi : ", curlResponse.Status)
	log.Println("Helper -- RESPONSE HEADERS CURL Curl Submit SIKP Transaksi : ", curlResponse.Header)
	log.Println("Helper -- REQUEST URL CURL Curl Submit SIKP Transaksi : ", curlResponse.Request.URL)
	log.Println("Helper -- REQUEST CONTENT LENGTH CURL Curl Submit SIKP Transaksi : ", curlResponse.Request.ContentLength)

	xml.NewDecoder(curlResponse.Body).Decode(&data)
	log.Println("RESULT DATA :", data)

	//ERROR STATUS CODE
	if curlResponse.StatusCode != 200 {
		// log.Println("Helper --- ERROR STATUS CODE CURL Curl Submit SIKP Transaksi ---")
		// log.Println("Helper --- ERROR STATUS CODE CURL Curl Submit SIKP Transaksi : ", curlResponse.StatusCode)

		joinString := []string{"Status Code : ", "Web Service Error", " | Error : ", data.Body.Fault.FaultString}
		erorrMessage := strings.Join(joinString, "")

		response := CurlSubmitSIKPTransaksiResponse{
			Status:       500,
			Message:      Constants.ErorrGeneralMessage,
			MessageLocal: erorrMessage,
		}

		return response, data
	}

	// log.Println("Helper --- DATA DECODE CURL Curl Submit SIKP Transaksi ---")
	// log.Println(data)

	response = CurlSubmitSIKPTransaksiResponse{
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
