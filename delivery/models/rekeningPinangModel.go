package models

import (
	"errors"
	"io"
	"log"
	"os"
	Config "pinang-mikro-3party/config"
	"strings"
	"time"

	Constants "pinang-mikro-3party/constants"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RekeningPinang struct {
	ID                     string  `gorm:"column:ID"`
	CIF                    string  `gorm:"column:CIF"`
	NoRekening             string  `gorm:"column:NO_REKENING"`
	NamaRekening           string  `gorm:"column:NAMA_REKENING"`
	CloseFlag              string  `gorm:"column:CLOSE_FLAG"`
	CloseDate              string  `gorm:"column:CLOSE_DATE"`
	Outstanding            float64 `gorm:"column:OUTSTANDING"`
	Plafond                float64 `gorm:"column:PLAFOND"`
	MainKolek              string  `gorm:"column:MAIN_KOLEK"`
	SubKolek               string  `gorm:"column:SUB_KOLEK"`
	SukuBunga              float64 `gorm:"column:SUKU_BUNGA"`
	DayPastDue             int     `gorm:"column:DAY_PAST_DUE"`
	NilaiPokokDitagih      float64 `gorm:"column:NILAI_POKOK_DITAGIH"`
	NilaiPokokSudahDibayar float64 `gorm:"column:NILAI_POKOK_SUDAH_DIBAYAR"`
	NilaiPokokBelumDibayar float64 `gorm:"column:NILAI_POKOK_BELUM_DIBAYAR"`
	NilaiBungaDitagih      float64 `gorm:"column:NILAI_BUNGA_DITAGIH"`
	NilaiBungaSudahDibayar float64 `gorm:"column:NILAI_BUNGA_SUDAH_DIBAYAR"`
	NilaiBungaBelumDibayar float64 `gorm:"column:NILAI_BUNGA_BELUM_DIBAYAR"`
	Penalti                float64 `gorm:"column:PENALTI"`
	GolonganDebitur        string  `gorm:"column:GOLONGAN_DEBITUR"`
	HubunganDenganBank     string  `gorm:"column:HUBUNGAN_DENGAN_BANK"`
	StatusDebitur          string  `gorm:"column:STATUS_DEBITUR"`
	KategoriDebitur        string  `gorm:"column:KATEGORI_DEBITUR"`
	KategoriPortofolio     string  `gorm:"column:KATEGORI_PORTOFOLIO"`
	LembagaPemeringkat     string  `gorm:"column:LEMBAGA_PEMERINGKAT"`
	PeringkatDebitur       string  `gorm:"column:PERINGKAT_DEBITUR"`
	TanggalPemeringkatan   string  `gorm:"column:TANGGAL_PEMERINGKATAN"`
	JenisKredit            string  `gorm:"column:JENIS_KREDIT"`
	SifatKredit            string  `gorm:"column:SIFAT_KREDIT"`
	JenisPenggunaan        string  `gorm:"column:JENIS_PENGGUNAAN"`
	OrientasiPenggunaan    string  `gorm:"column:ORIENTASI_PENGGUNAAN"`
	Currency               string  `gorm:"column:CURRENCY"`
	KategoriKredit         string  `gorm:"column:KATEGORI_KREDIT"`
	NoFasilitas            string  `gorm:"column:NO_FASILITAS"`
	TipePinjaman           string  `gorm:"column:TIPE_PINJAMAN"`
	KodeCabang             string  `gorm:"column:KODE_CABANG"`
	SegementasiKredit      string  `gorm:"column:SEGMENTASI_KREDIT"`
	Prakarsa               string  `gorm:"column:PRAKARSA"`
	SektorEkonomi          string  `gorm:"column:SEKTOR_EKONOMI"`
	Tenor                  int     `gorm:"column:TENOR"`
	TanggalMulai           string  `gorm:"column:TANGGAL_MULAI"`
	TanggalJatuhWaktu      string  `gorm:"column:TANGGAL_JATUH_WAKTU"`
	NamaPerusahaan         string  `gorm:"column:NAMA_PERUSAHAAN"`
	KodeEmployd            string  `gorm:"column:KODE_EMPLOYD"`
	AlamatEmail            string  `gorm:"column:ALAMAT_EMAIL"`
	NoHp                   int     `gorm:"column:NO_HP"`
	NoId                   string  `gorm:"column:NO_ID"`
	Agama                  string  `gorm:"column:AGAMA"`
	Gender                 string  `gorm:"column:GENDER"`
	TempatLahir            string  `gorm:"column:TEMPAT_LAHIR"`
	TanggalLahir           string  `gorm:"column:TANGGAL_LAHIR"`
	PendidikanTerakhir     string  `gorm:"column:PENDIDIKAN_TERAKHIR"`
	Alamat                 string  `gorm:"column:ALAMAT"`
	KodeKota               string  `gorm:"column:KODE_KOTA"`
	KodePos                int     `gorm:"column:KODE_POS"`
	NamaKerabat            string  `gorm:"column:NAMA_KERABAT"`
	TelpKerabat            int     `gorm:"column:TELP_KERABAT"`
	FlagWo                 string  `gorm:"column:FLAG_WO"`
	PYMHD                  string  `gorm:"column:PYMHD"`
	NomorAkad              int     `gorm:"column:NOMORAKAD"`
}

func (RekeningPinang) TableName() string {
	return "rekening_pinang"
}

type RekeningPinangParams struct {
	ID        string
	KtpNumber string
}

type RekeningPinangResponse struct {
	Status       int
	Message      string
	MessageLocal string
}

type RekeningPinangMapping struct {
	Acctno       string
	CIF          string
	Type         string
	NamaDebitur  string
	Status       string
	Cbal         string
	Bikole       string
	KodeCabang   string
	NoIdentitas  string
	TglLahir     string
	TglPembukaan string
	TglMature    string
}

func FindRekeningPinang(params RekeningPinangParams) (response RekeningPinangResponse, data []RekeningPinangMapping) {
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

	response = RekeningPinangResponse{
		Status:       400,
		Message:      "Server sedang dalam perbaikan. Silahkan coba beberapa saat lagi.",
		MessageLocal: "Error default",
	}

	/* open connection */
	db, err := Config.InitAgroDWH()
	if err != nil {
		response = RekeningPinangResponse{
			Status:       400,
			Message:      "Server sedang dalam perbaikan. Silahkan coba beberapa saat lagi.",
			MessageLocal: err.Error(),
		}

		return response, data
	}
	/* end open connection */

	/* build query */
	var result []RekeningPinangMapping

	query := db.Debug()

	yesterday := t.AddDate(0, 0, -1).String()
	yesterdayParse, _ := time.Parse(Constants.LayoutYMD, yesterday[:10])
	yesterdayDate := yesterdayParse.Format(Constants.LayoutYMD)
	yesterdayDateFormat := strings.ReplaceAll(yesterdayDate, "-", "")

	query = query.Raw("SELECT [CIF], [NO_REKENING] as acctno, 'PINANG' as type, [NAMA_REKENING] as nama_debitur "+
		",case when([CLOSE_FLAG] = 'Y') then '2' else '1' end as status "+
		",cast(cast([OUTSTANDING] as decimal(19, 2)) as varchar) as cbal, [MAIN_KOLEK] as bikole, [KODE_CABANG] as kode_cabang "+
		",[NO_ID] as no_identitas, substring([TANGGAL_LAHIR], 9,2) + '/' + substring([TANGGAL_LAHIR], 6,2) + '/'+ substring([TANGGAL_LAHIR], 1, 4) as tgl_lahir, substring([TANGGAL_MULAI], 9,2) + '/' + substring([TANGGAL_MULAI], 6,2) + '/' + substring([TANGGAL_MULAI], 1, 4) as tgl_pembukaan, substring([CLOSE_DATE], 9,2) + '/' + substring([CLOSE_DATE], 6,2) + '/' + substring([CLOSE_DATE], 1, 4) as tgl_mature "+
		"FROM [DWH_PINANG].[dbo].[rekening_pinang] "+
		"where no_id = ? "+
		" union all "+
		"SELECT C.[CIFNUM] as [CIF], M.ACCOUNTNUMBER as acctno, M.LOANTYPE as type, M.SHORTNAME as nama_debitur, M.STATUS as status, "+
		"cast(cast(M.CBAL as decimal(19, 2)) as varchar) as cbal, B.[KOLEK] as bikole, substring(M.ACCOUNTNUMBER, 1, 4) as kode_cabang, NO_ID as no_identitas, C.TANGGAL_LAHIR as tgl_lahir, substring(M.OPENDAT6, 1,2) + '/' + substring(M.OPENDAT6, 3,2) + '/20' + substring(M.OPENDAT6,5,2) as tgl_pembukaan, substring(M.MATURITYDATE6, 1,2) + '/' + substring(M.MATURITYDATE6, 3,2) + '/20' + substring(M.MATURITYDATE6,5,2) as tgl_mature "+
		"FROM[DWH_ABCS_M_CFMAST].[dbo].[ABCS_M_CFMAST_"+yesterdayDateFormat+"] C "+
		"JOIN[DWH_ABCS_M_LNMAST].[dbo].[ABCS_M_LNMAST_"+yesterdayDateFormat+"] M on C.CIFNUM = M.CIFNUM "+
		"JOIN[DWH_ABCS_M_LNBICD].[dbo].[ABCS_M_LNBICD_"+yesterdayDateFormat+"] B on M.ACCOUNTNUMBER = B.ACCOUNTNUMBER "+
		"WHERE C.NO_ID = ? ", params.KtpNumber, params.KtpNumber).Scan(&result)

	// query = query.Raw("SELECT [NO_REKENING] as acctno , CIF FROM dwh_pinang.dbo.rekening_pinang WHERE no_id = ?", params.KtpNumber).Scan(&result)

	/* error handling */
	if query.Error != nil { /*error query*/
		if query.Statement.Error != nil {
			response = RekeningPinangResponse{
				Status:       404,
				Message:      "Rekening Pinang Tidak Ditemukan",
				MessageLocal: query.Statement.Error.Error(),
			}

			log.Println("Model -- ERROR QUERY REKENING PINANG FIND --")
			log.Println("Model -- ERROR QUERY REKENING PINANG FIND : ", query.Statement.Error.Error())
		}

		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			response = RekeningPinangResponse{
				Status:       404,
				Message:      "Rekening Pinang Tidak Ditemukan",
				MessageLocal: "Error record not found",
			}

			log.Println("Model -- DATA NOT FOUND REKENING PINANG FIND --")
		}

		return response, data
	}

	if len(result) < 1 { /*data empty*/
		response = RekeningPinangResponse{
			Status:       404,
			Message:      "Rekening Pinang Tidak Ditemukan",
			MessageLocal: "Data Rekening Pinang tidak ditemukan",
		}

		log.Println("Model -- DATA EMPTY REKENING PINANG FIND --")

		return response, data
	}
	/* end error handling */

	//MAPPING DATA REKENING PINANG
	for _, vData := range result {

		allData := RekeningPinangMapping{
			Acctno:       vData.Acctno,
			CIF:          vData.CIF,
			Type:         vData.Type,
			NamaDebitur:  vData.NamaDebitur,
			Status:       vData.Status,
			Cbal:         vData.Cbal,
			Bikole:       vData.Bikole,
			KodeCabang:   vData.KodeCabang,
			NoIdentitas:  vData.NoIdentitas,
			TglLahir:     vData.TglLahir,
			TglPembukaan: vData.TglPembukaan,
			TglMature:    vData.TglMature,
		}

		data = append(data, allData)
	}

	//END MAPPING DATA RKENING PINANG
	log.Println("Model -- RESPONSE MODEL REKENING PINANG FIND --")
	// log.Println(data)
	// return

	response = RekeningPinangResponse{
		Status:       200,
		Message:      "Data ditemukan",
		MessageLocal: "Data alid",
	}

	return response, data

}
