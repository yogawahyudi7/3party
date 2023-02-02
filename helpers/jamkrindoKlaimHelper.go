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

type CurlJamkrindoKlaimParams struct {
	CabangRekanan      string
	JenisAgunan        string
	JenisKredit        string
	JenisKur           string
	JenisPengikatan    string
	JumlahKerugian     string
	KodeBank           string
	KodeLbu            string
	KodeUker           string
	NamaDebitur        string
	NilaiPengikatan    string
	NilaiPenjaminan    string
	NilaiPersen        string
	NilaiTuntutanKlaim string
	NoRekening         string
	NoSp2              string
	NoSpk              string
	NomorPk            string
	NomorSsertifikat   string
	PenggunaanKredit   string
	Periode            string
	Plafond            string
	SebabKlaim         string
	TglJatuhTempo      string
	TglMulai           string
	TglSp2             string
	TglSpk             string
	TglSsertifikat     string
	TglStatus          string
	Tindakan1          string
	Tindakan2          string
	Tindakan3          string
	Tindakan4          string
	Tindakan5          string
	TunggakanBunga     string
	TunggakanDenda     string
	TunggakanPokok     string
}

type CurlJamkrindoKlaimResponse struct {
	Status       int
	Message      string
	MessageLocal string
}

type CurlJamkrindoKlaimMapping struct {
	XMLName xml.Name           `xml:"Envelope"`
	Body    BodyJamkrindoKlaim `xml:"Body"`
}

type BodyJamkrindoKlaim struct {
	XMLName  xml.Name               `xml:"Body"`
	Fault    FaultJamkrindoKlaim    `xml:"Fault"`
	Response ResponseJamkrindoKlaim `xml:"Jamkrindo_KlaimResponse"`
}

type FaultJamkrindoKlaim struct {
	XMLName     xml.Name `xml:"Fault"`
	FaultCode   string   `xml:"faultcode"`
	FaultString string   `xml:"faultstring"`
}

type ResponseJamkrindoKlaim struct {
	XMLName xml.Name             `xml:"Jamkrindo_KlaimResponse"`
	Result  ResultJamkrindoKlaim `xml:"Jamkrindo_KlaimResult"`
}

type ResultJamkrindoKlaim struct {
	XMLName xml.Name           `xml:"Jamkrindo_KlaimResult"`
	Error   bool               `xml:"error"`
	Code    string             `xml:"code"`
	Message string             `xml:"message"`
	Data    DataJamkrindoKlaim `xml:"data"`
}

type DataJamkrindoKlaim struct {
	XMLName            xml.Name                 `xml:"data"`
	DataJamkrindoKlaim []JamkrindoKlaimResponse `xml:"jamkrindo_klaim"`
}

type JamkrindoKlaimResponse struct {
	XMLName         xml.Name `xml:"jamkrindo_klaim"`
	FlagStatus      string   `xml:"flag_status"`
	NoRespond       string   `xml:"no_respond"`
	TglRespondKlaim string   `xml:"tgl_respond_klaim"`
}

func CurlJamkrindoKlaim(params CurlJamkrindoKlaimParams) (response CurlJamkrindoKlaimResponse, data CurlJamkrindoKlaimMapping) {
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

	log.Println("--ENDPOINT API Curl Jamkrindo Klaim --")
	log.Println(endpoint)

	transCfg := Config.TransportConfig

	client := &http.Client{
		Timeout:   time.Duration(25 * time.Second),
		Transport: transCfg,
	}

	xmlnSchemeXsi := `xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" `
	xmlnSchemeXsd := `xmlns:xsd="http://www.w3.org/2001/XMLSchema" `
	xmlnSoap := `xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"`

	CabangRekanan := params.CabangRekanan
	JenisAgunan := params.JenisAgunan
	JenisKredit := params.JenisKredit
	JenisKur := params.JenisKur
	JenisPengikatan := params.JenisPengikatan
	JumlahKerugian := params.JumlahKerugian
	KodeBank := params.KodeBank
	KodeLbu := params.KodeLbu
	KodeUker := params.KodeUker
	NamaDebitur := params.NamaDebitur
	NilaiPengikatan := params.NilaiPengikatan
	NilaiPenjaminan := params.NilaiPenjaminan
	NilaiPersen := params.NilaiPersen
	NilaiTuntutanKlaim := params.NilaiTuntutanKlaim
	NoRekening := params.NoRekening
	NoSp2 := params.NoSp2
	NoSpk := params.NoSpk
	NomorPk := params.NomorPk
	NomorSsertifikat := params.NomorSsertifikat
	PenggunaanKredit := params.PenggunaanKredit
	Periode := params.Periode
	Plafond := params.Plafond
	SebabKlaim := params.SebabKlaim
	TglJatuhTempo := params.TglJatuhTempo
	TglMulai := params.TglMulai
	TglSp2 := params.TglSp2
	TglSpk := params.TglSpk
	TglSsertifikat := params.TglSsertifikat
	TglStatus := params.TglStatus
	Tindakan1 := params.Tindakan1
	Tindakan2 := params.Tindakan2
	Tindakan3 := params.Tindakan3
	Tindakan4 := params.Tindakan4
	Tindakan5 := params.Tindakan5
	TunggakanBunga := params.TunggakanBunga
	TunggakanDenda := params.TunggakanDenda
	TunggakanPokok := params.TunggakanPokok

	submit := `
	<Jamkrindo_Klaim xmlns="http://tempuri.org/">
      <klaim_jamkrindo>
        <CABANG_REKANAN> ` + CabangRekanan + `</CABANG_REKANAN>
        <JENIS_AGUNAN> ` + JenisAgunan + `</JENIS_AGUNAN>
        <JENIS_KREDIT> ` + JenisKredit + `</JENIS_KREDIT>
        <JENIS_KUR> ` + JenisKur + `</JENIS_KUR>
        <JENIS_PENGIKATAN> ` + JenisPengikatan + `</JENIS_PENGIKATAN>
        <JUMLAH_KERUGIAN> ` + JumlahKerugian + `</JUMLAH_KERUGIAN>
        <KODE_BANK> ` + KodeBank + `</KODE_BANK>
        <KODE_LBU> ` + KodeLbu + `</KODE_LBU>
        <KODE_UKER> ` + KodeUker + `</KODE_UKER>
        <NAMA_DEBITUR> ` + NamaDebitur + `</NAMA_DEBITUR>
        <NILAI_PENGIKATAN> ` + NilaiPengikatan + `</NILAI_PENGIKATAN>
        <NILAI_PENJAMINAN> ` + NilaiPenjaminan + `</NILAI_PENJAMINAN>
        <NILAI_PERSEN> ` + NilaiPersen + `</NILAI_PERSEN>
        <NILAI_TUNTUTAN_KLAIM> ` + NilaiTuntutanKlaim + `</NILAI_TUNTUTAN_KLAIM>
        <NO_REKENING> ` + NoRekening + `</NO_REKENING>
        <NO_SP2> ` + NoSp2 + `</NO_SP2>
        <NO_SPK> ` + NoSpk + `</NO_SPK>
        <NOMOR_PK> ` + NomorPk + `</NOMOR_PK>
        <NOMOR_SSERTIFIKAT> ` + NomorSsertifikat + `</NOMOR_SSERTIFIKAT>
        <PENGGUNAAN_KREDIT> ` + PenggunaanKredit + `</PENGGUNAAN_KREDIT>
        <PERIODE> ` + Periode + `</PERIODE>
        <PLAFOND> ` + Plafond + `</PLAFOND>
        <SEBAB_KLAIM> ` + SebabKlaim + `</SEBAB_KLAIM>
        <TGL_JATUH_TEMPO> ` + TglJatuhTempo + `</TGL_JATUH_TEMPO>
        <TGL_MULAI> ` + TglMulai + `</TGL_MULAI>
        <TGL_SP2> ` + TglSp2 + `</TGL_SP2>
        <TGL_SPK> ` + TglSpk + `</TGL_SPK>
        <TGL_SSERTIFIKAT> ` + TglSsertifikat + `</TGL_SSERTIFIKAT>
        <TGL_STATUS> ` + TglStatus + `</TGL_STATUS>
        <TINDAKAN1> ` + Tindakan1 + `</TINDAKAN1>
        <TINDAKAN2> ` + Tindakan2 + `</TINDAKAN2>
        <TINDAKAN3> ` + Tindakan3 + `</TINDAKAN3>
        <TINDAKAN4> ` + Tindakan4 + `</TINDAKAN4>
        <TINDAKAN5> ` + Tindakan5 + `</TINDAKAN5>
        <TUNGGAKAN_BUNGA> ` + TunggakanBunga + `</TUNGGAKAN_BUNGA>
        <TUNGGAKAN_DENDA> ` + TunggakanDenda + `</TUNGGAKAN_DENDA>
        <TUNGGAKAN_POKOK> ` + TunggakanPokok + `</TUNGGAKAN_POKOK>
      </klaim_jamkrindo>
    </Jamkrindo_Klaim>
	`
	body := `<soap:Envelope ` + xmlnSchemeXsi + xmlnSchemeXsd + xmlnSoap + `><soap:Body>` + submit + `</soap:Body></soap:Envelope>`

	fmt.Println(body)
	r, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(body))) // URL-encoded payload

	//ERROR REQUEST
	if err != nil {
		log.Println("Helper -- ERROR REQUEST CURL Curl Jamkrindo Klaim  ---")
		log.Println("Helper -- ERROR REQUEST CURL Curl Jamkrindo Klaim  : ", err.Error())

		joinString := []string{"Status Code : ", "-", " | Error : ", err.Error()}
		erorrMessage := strings.Join(joinString, "")

		response := CurlJamkrindoKlaimResponse{
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
		log.Println("Helper -- ERROR HIT API CURL Curl Jamkrindo Klaim  ---")
		log.Println("Helper -- ERROR HIT API CURL Curl Jamkrindo Klaim  : ", err.Error())

		joinString := []string{"Status Code : ", "-", " | Error : ", err.Error()}
		erorrMessage := strings.Join(joinString, "")

		response := CurlJamkrindoKlaimResponse{
			Status:       500,
			Message:      Constants.ErorrGeneralMessage,
			MessageLocal: erorrMessage,
		}

		return response, data
	}

	defer curlResponse.Body.Close()

	log.Println("Helper -- RESPONSE STATUS CURL Curl Jamkrindo Klaim  : ", curlResponse.Status)
	log.Println("Helper -- RESPONSE HEADERS CURL Curl Jamkrindo Klaim  : ", curlResponse.Header)
	log.Println("Helper -- REQUEST URL CURL Curl Jamkrindo Klaim  : ", curlResponse.Request.URL)
	log.Println("Helper -- REQUEST CONTENT LENGTH CURL Curl Jamkrindo Klaim  : ", curlResponse.Request.ContentLength)

	xml.NewDecoder(curlResponse.Body).Decode(&data)
	log.Println("RESULT DATA :", data)

	//ERROR STATUS CODE
	if curlResponse.StatusCode != 200 {
		// log.Println("Helper --- ERROR STATUS CODE CURL Curl Jamkrindo Klaim  ---")
		// log.Println("Helper --- ERROR STATUS CODE CURL Curl Jamkrindo Klaim  : ", curlResponse.StatusCode)

		joinString := []string{"Status Code : ", "Web Service Error", " | Error : ", data.Body.Fault.FaultString}
		erorrMessage := strings.Join(joinString, "")

		response := CurlJamkrindoKlaimResponse{
			Status:       500,
			Message:      Constants.ErorrGeneralMessage,
			MessageLocal: erorrMessage,
		}

		return response, data
	}

	// log.Println("Helper --- DATA DECODE CURL Curl Jamkrindo Klaim  ---")
	// log.Println(data)

	response = CurlJamkrindoKlaimResponse{
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
