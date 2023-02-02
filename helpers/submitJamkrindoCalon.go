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

type CurlSubmitJamkrindoCalonParams struct {
	AlamatDebitur        string
	AlamatUsaha          string
	CabangRekanan        string
	Cif                  string
	Coverage             string
	FlagTransfer         string
	IDCalonDebiturKur    string
	JangkaWaktu          string
	JenisAgunan          string
	JenisIdentitas       string
	JenisKelamin         string
	JenisKredit          string
	JenisKur             string
	JenisLinkage         string
	JenisPengikatan      string
	JmlTKerja            string
	KodeBank             string
	KodePekerjaan        string
	KodePos              string
	KodeSektor           string
	KodeUker             string
	LembagaLinkage       string
	ModalUsaha           string
	NamaDebitur          string
	NilaiAgunan          string
	NoHp                 string
	NoIdentitas          string
	NoIjinUsaha          string
	NoPk                 string
	NoRekening           string
	NoTelepon            string
	NomorAplikasi        string
	Npwp                 string
	PlafonKredit         string
	StatusAplikasi       string
	StatusKolektibilitas string
	StatusLunas          string
	SukuBunga            string
	TanggalAkhir         string
	TanggalAwal          string
	TanggalLahir         string
	TanggalMulaiUsaha    string
	TanggalPk            string
	TanggalRekam         string
	UsahaProduktif       string
}

type CurlSubmitJamkrindoCalonResponse struct {
	Status       int
	Message      string
	MessageLocal string
}

type CurlSubmitJamkrindoCalonMapping struct {
	XMLName xml.Name                 `xml:"Envelope"`
	Body    BodySubmitJamkrindoCalon `xml:"Body"`
}

type BodySubmitJamkrindoCalon struct {
	XMLName  xml.Name                     `xml:"Body"`
	Fault    FaultSubmitJamkrindoCalon    `xml:"Fault"`
	Response ResponseSubmitJamkrindoCalon `xml:"submit_Jamkrindo_CalonResponse"`
}

type FaultSubmitJamkrindoCalon struct {
	XMLName     xml.Name `xml:"Fault"`
	FaultCode   string   `xml:"faultcode"`
	FaultString string   `xml:"faultstring"`
}

type ResponseSubmitJamkrindoCalon struct {
	XMLName xml.Name                   `xml:"submit_Jamkrindo_CalonResponse"`
	Result  ResultSubmitJamkrindoCalon `xml:"submit_Jamkrindo_CalonResult"`
}

type ResultSubmitJamkrindoCalon struct {
	XMLName xml.Name                 `xml:"submit_Jamkrindo_CalonResult"`
	Error   bool                     `xml:"error"`
	Code    string                   `xml:"code"`
	Message string                   `xml:"message"`
	Data    DataSubmitJamkrindoCalon `xml:"data"`
}

type DataSubmitJamkrindoCalon struct {
	XMLName       xml.Name                       `xml:"data"`
	DataJamkrindo []SubmitJamkrindoCalonResponse `xml:"submit_Jamkrindo_Calonz"`
}

type JamkrindoData struct {
	XMLName              xml.Name `xml:"submit_Jamkrindo_Calon"`
	AlamatDebitur        string   `xml:"ALAMAT_DEBITUR"`
	AlamatUsaha          string   `xml:"ALAMAT_USAHA"`
	CabangRekanan        string   `xml:"CABANG_REKANAN"`
	Cif                  string   `xml:"CIF"`
	Coverage             string   `xml:"COVERAGE"`
	FlagTransfer         string   `xml:"FLAG_TRANSFER"`
	IDCalonDebiturKur    string   `xml:"ID_CALON_DEBITUR_KUR"`
	JangkaWaktu          string   `xml:"JANGKA_WAKTU"`
	JenisAgunan          string   `xml:"JENIS_AGUNAN"`
	JenisIdentitas       string   `xml:"JENIS_IDENTITAS"`
	JenisKelamin         string   `xml:"JENIS_KELAMIN"`
	JenisKredit          string   `xml:"JENIS_KREDIT"`
	JenisKur             string   `xml:"JENIS_KUR"`
	JenisLinkage         string   `xml:"JENIS_LINKAGE"`
	JenisPengikatan      string   `xml:"JENIS_PENGIKATAN"`
	JmlTKerja            string   `xml:"JML_T_KERJA"`
	KodeBank             string   `xml:"KODE_BANK"`
	KodePekerjaan        string   `xml:"KODE_PEKERJAAN"`
	KodePos              string   `xml:"KODE_POS"`
	KodeSektor           string   `xml:"KODE_SEKTOR"`
	KodeUker             string   `xml:"KODE_UKER"`
	LembagaLinkage       string   `xml:"LEMBAGA_LINKAGE"`
	ModalUsaha           string   `xml:"MODAL_USAHA"`
	NamaDebitur          string   `xml:"NAMA_DEBITUR"`
	NilaiAgunan          string   `xml:"NILAI_AGUNAN"`
	NoHp                 string   `xml:"NO_HP"`
	NoIdentitas          string   `xml:"NO_IDENTITAS"`
	NoIjinUsaha          string   `xml:"NO_IJIN_USAHA"`
	NoPk                 string   `xml:"NO_PK"`
	NoRekening           string   `xml:"NO_REKENING"`
	NoTelepon            string   `xml:"NO_TELEPON"`
	NomorAplikasi        string   `xml:"NOMOR_APLIKASI"`
	Npwp                 string   `xml:"NPWP"`
	PlafonKredit         string   `xml:"PLAFON_KREDIT"`
	StatusAplikasi       string   `xml:"STATUS_APLIKASI"`
	StatusKolektibilitas string   `xml:"STATUS_KOLEKTIBILITAS"`
	StatusLunas          string   `xml:"STATUS_LUNAS"`
	SukuBunga            string   `xml:"SUKU_BUNGA"`
	TanggalAkhir         string   `xml:"TANGGAL_AKHIR"`
	TanggalAwal          string   `xml:"TANGGAL_AWAL"`
	TanggalLahir         string   `xml:"TANGGAL_LAHIR"`
	TanggalMulaiUsaha    string   `xml:"TANGGAL_MULAI_USAHA"`
	TanggalPk            string   `xml:"TANGGAL_PK"`
	TanggalRekam         string   `xml:"TANGGAL_REKAM"`
	UsahaProduktif       string   `xml:"USAHA_PRODUKTIF"`
}

type SubmitJamkrindoCalonResponse struct {
	XMLName        xml.Name `xml:"submit_Jamkrindo_Calon"`
	NoIjinPrinsip  string   `xml:"no_ijin_prinsip"`
	NoSertifikat   string   `xml:"no_sertifikat"`
	NomorUrut      string   `xml:"nomor_urut"`
	TglIjinPrinsip string   `xml:"tgl_ijin_prinsip"`
	TglSertifikat  string   `xml:"tgl_sertifikat"`
}

func CurlSubmitJamkrindoCalon(params CurlSubmitJamkrindoCalonParams) (response CurlSubmitJamkrindoCalonResponse, data CurlSubmitJamkrindoCalonMapping) {
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

	log.Println("--ENDPOINT API Curl Submit Jamkrindo Calon--")
	log.Println(endpoint)

	transCfg := Config.TransportConfig

	client := &http.Client{
		Timeout:   time.Duration(25 * time.Second),
		Transport: transCfg,
	}

	xmlnSchemeXsi := `xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" `
	xmlnSchemeXsd := `xmlns:xsd="http://www.w3.org/2001/XMLSchema" `
	xmlnSoap := `xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"`
	AlamatDebitur := params.AlamatDebitur
	AlamatUsaha := params.AlamatUsaha
	CabangRekanan := params.CabangRekanan
	Cif := params.Cif
	Coverage := params.Coverage
	FlagTransfer := params.FlagTransfer
	IDCalonDebiturKur := params.IDCalonDebiturKur
	JangkaWaktu := params.JangkaWaktu
	JenisAgunan := params.JenisAgunan
	JenisIdentitas := params.JenisIdentitas
	JenisKelamin := params.JenisKelamin
	JenisKredit := params.JenisKredit
	JenisKur := params.JenisKur
	JenisLinkage := params.JenisLinkage
	JenisPengikatan := params.JenisPengikatan
	JmlTKerja := params.JmlTKerja
	KodeBank := params.KodeBank
	KodePekerjaan := params.KodePekerjaan
	KodePos := params.KodePos
	KodeSektor := params.KodeSektor
	KodeUker := params.KodeUker
	LembagaLinkage := params.LembagaLinkage
	ModalUsaha := params.ModalUsaha
	NamaDebitur := params.NamaDebitur
	NilaiAgunan := params.NilaiAgunan
	NoHp := params.NoHp
	NoIdentitas := params.NoIdentitas
	NoIjinUsaha := params.NoIjinUsaha
	NoPk := params.NoPk
	NoRekening := params.NoRekening
	NoTelepon := params.NoTelepon
	NomorAplikasi := params.NomorAplikasi
	Npwp := params.Npwp
	PlafonKredit := params.PlafonKredit
	StatusAplikasi := params.StatusAplikasi
	StatusKolektibilitas := params.StatusKolektibilitas
	StatusLunas := params.StatusLunas
	SukuBunga := params.SukuBunga
	TanggalAkhir := params.TanggalAkhir
	TanggalAwal := params.TanggalAwal
	TanggalLahir := params.TanggalLahir
	TanggalMulaiUsaha := params.TanggalMulaiUsaha
	TanggalPk := params.TanggalPk
	TanggalRekam := params.TanggalRekam
	UsahaProduktif := params.UsahaProduktif
	submit := `
	<submit_Jamkrindo_Calon xmlns="http://tempuri.org/">
      <k_calon_jamkrindo>
        <ALAMAT_DEBITUR> ` + AlamatDebitur + `</ALAMAT_DEBITUR>
        <ALAMAT_USAHA> ` + AlamatUsaha + `</ALAMAT_USAHA>
        <CABANG_REKANAN> ` + CabangRekanan + `</CABANG_REKANAN>
        <CIF> ` + Cif + `</CIF>
        <COVERAGE> ` + Coverage + `</COVERAGE>
        <FLAG_TRANSFER> ` + FlagTransfer + `</FLAG_TRANSFER>
        <ID_CALON_DEBITUR_KUR> ` + IDCalonDebiturKur + `</ID_CALON_DEBITUR_KUR>
        <JANGKA_WAKTU> ` + JangkaWaktu + `</JANGKA_WAKTU>
        <JENIS_AGUNAN> ` + JenisAgunan + `</JENIS_AGUNAN>
        <JENIS_IDENTITAS> ` + JenisIdentitas + `</JENIS_IDENTITAS>
        <JENIS_KELAMIN> ` + JenisKelamin + `</JENIS_KELAMIN>
        <JENIS_KREDIT> ` + JenisKredit + `</JENIS_KREDIT>
        <JENIS_KUR> ` + JenisKur + `</JENIS_KUR>
        <JENIS_LINKAGE> ` + JenisLinkage + `</JENIS_LINKAGE>
        <JENIS_PENGIKATAN> ` + JenisPengikatan + `</JENIS_PENGIKATAN>
        <JML_T_KERJA> ` + JmlTKerja + `</JML_T_KERJA>
        <KODE_BANK> ` + KodeBank + `</KODE_BANK>
        <KODE_PEKERJAAN> ` + KodePekerjaan + `</KODE_PEKERJAAN>
        <KODE_POS> ` + KodePos + `</KODE_POS>
        <KODE_SEKTOR> ` + KodeSektor + `</KODE_SEKTOR>
        <KODE_UKER> ` + KodeUker + `</KODE_UKER>
        <LEMBAGA_LINKAGE> ` + LembagaLinkage + `</LEMBAGA_LINKAGE>
        <MODAL_USAHA> ` + ModalUsaha + `</MODAL_USAHA>
        <NAMA_DEBITUR> ` + NamaDebitur + `</NAMA_DEBITUR>
        <NILAI_AGUNAN> ` + NilaiAgunan + `</NILAI_AGUNAN>
        <NO_HP> ` + NoHp + `</NO_HP>
        <NO_IDENTITAS> ` + NoIdentitas + `</NO_IDENTITAS>
        <NO_IJIN_USAHA> ` + NoIjinUsaha + `</NO_IJIN_USAHA>
        <NO_PK> ` + NoPk + `</NO_PK>
        <NO_REKENING> ` + NoRekening + `</NO_REKENING>
        <NO_TELEPON> ` + NoTelepon + `</NO_TELEPON>
        <NOMOR_APLIKASI> ` + NomorAplikasi + `</NOMOR_APLIKASI>
        <NPWP> ` + Npwp + `</NPWP>
        <PLAFON_KREDIT> ` + PlafonKredit + `</PLAFON_KREDIT>
        <STATUS_APLIKASI> ` + StatusAplikasi + `</STATUS_APLIKASI>
        <STATUS_KOLEKTIBILITAS> ` + StatusKolektibilitas + `</STATUS_KOLEKTIBILITAS>
        <STATUS_LUNAS> ` + StatusLunas + `</STATUS_LUNAS>
        <SUKU_BUNGA> ` + SukuBunga + `</SUKU_BUNGA>
        <TANGGAL_AKHIR> ` + TanggalAkhir + `</TANGGAL_AKHIR>
        <TANGGAL_AWAL> ` + TanggalAwal + `</TANGGAL_AWAL>
        <TANGGAL_LAHIR> ` + TanggalLahir + `</TANGGAL_LAHIR>
        <TANGGAL_MULAI_USAHA> ` + TanggalMulaiUsaha + `</TANGGAL_MULAI_USAHA>
        <TANGGAL_PK> ` + TanggalPk + `</TANGGAL_PK>
        <TANGGAL_REKAM> ` + TanggalRekam + `</TANGGAL_REKAM>
        <USAHA_PRODUKTIF> ` + UsahaProduktif + `</USAHA_PRODUKTIF>
      </k_calon_jamkrindo>
    </submit_Jamkrindo_Calon>
	`
	body := `<soap:Envelope ` + xmlnSchemeXsi + xmlnSchemeXsd + xmlnSoap + `><soap:Body>` + submit + `</soap:Body></soap:Envelope>`

	fmt.Println(body)
	r, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(body))) // URL-encoded payload

	//ERROR REQUEST
	if err != nil {
		log.Println("Helper -- ERROR REQUEST CURL Curl Submit Jamkrindo Calon ---")
		log.Println("Helper -- ERROR REQUEST CURL Curl Submit Jamkrindo Calon : ", err.Error())

		joinString := []string{"Status Code : ", "-", " | Error : ", err.Error()}
		erorrMessage := strings.Join(joinString, "")

		response := CurlSubmitJamkrindoCalonResponse{
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
		log.Println("Helper -- ERROR HIT API CURL Curl Submit Jamkrindo Calon ---")
		log.Println("Helper -- ERROR HIT API CURL Curl Submit Jamkrindo Calon : ", err.Error())

		joinString := []string{"Status Code : ", "-", " | Error : ", err.Error()}
		erorrMessage := strings.Join(joinString, "")

		response := CurlSubmitJamkrindoCalonResponse{
			Status:       500,
			Message:      Constants.ErorrGeneralMessage,
			MessageLocal: erorrMessage,
		}

		return response, data
	}

	defer curlResponse.Body.Close()

	log.Println("Helper -- RESPONSE STATUS CURL Curl Submit Jamkrindo Calon : ", curlResponse.Status)
	log.Println("Helper -- RESPONSE HEADERS CURL Curl Submit Jamkrindo Calon : ", curlResponse.Header)
	log.Println("Helper -- REQUEST URL CURL Curl Submit Jamkrindo Calon : ", curlResponse.Request.URL)
	log.Println("Helper -- REQUEST CONTENT LENGTH CURL Curl Submit Jamkrindo Calon : ", curlResponse.Request.ContentLength)

	xml.NewDecoder(curlResponse.Body).Decode(&data)
	log.Println("RESULT DATA :", data)

	//ERROR STATUS CODE
	if curlResponse.StatusCode != 200 {
		// log.Println("Helper --- ERROR STATUS CODE CURL Curl Submit Jamkrindo Calon ---")
		// log.Println("Helper --- ERROR STATUS CODE CURL Curl Submit Jamkrindo Calon : ", curlResponse.StatusCode)

		joinString := []string{"Status Code : ", "Web Service Error", " | Error : ", data.Body.Fault.FaultString}
		erorrMessage := strings.Join(joinString, "")

		response := CurlSubmitJamkrindoCalonResponse{
			Status:       500,
			Message:      Constants.ErorrGeneralMessage,
			MessageLocal: erorrMessage,
		}

		return response, data
	}

	// log.Println("Helper --- DATA DECODE CURL Curl Submit Jamkrindo Calon ---")
	// log.Println(data)

	response = CurlSubmitJamkrindoCalonResponse{
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
