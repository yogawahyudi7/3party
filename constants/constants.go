package constants

const (
	UserBasic = "b3BueDIwMjI=" // opnx2022
	PassBasic = "YWJDS29waWtv" // abCKopiko

	ChannelLogging3PartyAPI = "3party-log"

	CURLHeaderContentType             = "Content-Type"
	CURLHeaderContentTypeValue        = "application/x-www-form-urlencoded"
	CURLHeaderContentTypeValueTextXML = "text/xml; charset=utf-8"
	CURLHeaderCacheControl            = "Cache-Control"
	CURLHeaderCacheControlValue       = "no-cache"

	Timezone           = "Asia/Jakarta"
	FullLayoutTime     = "2006-01-02 15:04:05.000"
	FullLayoutDateTime = "2006-01-02 15:04:05"
	LayoutYMD          = "2006-01-02"
	LayoutDMY          = "02-01-2006"
	LayoutYMDHIS       = "20060102150405"
	LayoutDMYSlash     = "02/01/2006"
	LayoutFull         = "2006-01-02T15:04:05.000Z"

	TextLogDataRequest  = "Logging ** DATA REQUEST**"
	TextLogDataResponse = "Logging ** DATA RESPONSE**"

	//LOGGING DESC
	Desc3PartyLogging = "3party-logging"

	//ENDPOINT
	EndpointSIKPVerification     = "/v/1/sikp/verification"
	EndpointSIKPCheckPlafond     = "/v/1/sikp/check-plafond"
	EndpointSubmitJamkrindoCalon = "v/1/sikp/submit-jamkrindo-calon"
	EndpointJamkrindoKlaim       = "/v/1/sikp/jamkrindo-klaim"
	EndpointSubmitSIKPTransaksi  = "/v/1/sikp/submit-transaksi"

	DataFoundMessage       = "Data Ditemukan."
	DataNotFoundMessage    = "Data Tidak Ditemukan."
	SuccessProccessMessage = "Success Process."
	FailedProccessMessage  = "Failed Process."
	MessageOutOfService    = "Maaf, Untuk saat ini kamu belum bisa mengirimkan datamu karena sudah melewati masa operational kami diatas jam 23.00 WIB. Kamu dapat mengirimkan kembali pukul 06.00 WIB. Terima kasih."
	ErorrGeneralMessage    = "Maaf, Server sedang dalam perbaikan. Silahkan coba beberapa saat lagi."
)
