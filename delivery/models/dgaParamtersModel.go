package models

import (
	"errors"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	Config "pinang-mikro-3party/config"
)

type DGAParameters struct {
	// gorm.Model
	Value string `gorm:"column:value"`
	Desc  string `gorm:"column:desc"`
}

func (DGAParameters) TableName() string {
	return "param"
}

type DGAParametersParams struct {
	Value string
	Desc  string
}

type DGAParametersResponse struct {
	Status       int
	Message      string
	MessageLocal string
}

type DGAParametersMapping struct {
	Value string
	Desc  string
}

func FindDGAParamters(params DGAParametersParams) (response DGAParametersResponse, data DGAParametersMapping) {
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

	response = DGAParametersResponse{
		Status:       400,
		Message:      "Server sedang dalam perbaikan. Silahkan coba beberapa saat lagi.",
		MessageLocal: "Error default",
	}

	/* open connection */
	db, err := Config.InitDigiagriDb()
	if err != nil {
		response = DGAParametersResponse{
			Status:       400,
			Message:      "Server sedang dalam perbaikan. Silahkan coba beberapa saat lagi.",
			MessageLocal: err.Error(),
		}

		return response, data
	}
	/* end open connection */

	/* build query */
	var result DGAParameters

	query := db.Debug()

	query = query.Where("[desc] = ?", params.Desc)

	query = query.Find(&result)

	/* error handling */
	if query.Error != nil { /*error query*/
		if query.Statement.Error != nil {
			response = DGAParametersResponse{
				Status:       404,
				Message:      "ACL Tidak Ditemukan",
				MessageLocal: query.Statement.Error.Error(),
			}

			log.Println("Model -- ERROR QUERY DGA PARAMETERS FIND --")
			log.Println("Model -- ERROR QUERY DGA PARAMETERS FIND : ", query.Statement.Error.Error())
		}

		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			response = DGAParametersResponse{
				Status:       404,
				Message:      "ACL Tidak Ditemukan",
				MessageLocal: "Error record not found",
			}

			log.Println("Model -- DATA NOT FOUND DGA PARAMETERS FIND --")
		}

		return response, data
	}

	if result.Value == "" { /*data empty*/
		response = DGAParametersResponse{
			Status:       404,
			Message:      "ACL Tidak Ditemukan",
			MessageLocal: "Data ACL tidak ditemukan",
		}

		log.Println("Model -- DATA EMPTY DGA PARAMETERS FIND --")

		return response, data
	}
	/* end error handling */

	//MAPPING DATA DGA PARAMETERS

	// data = DGAParametersMapping{
	// 	Value: result.Value,
	// 	Desc:  result.Desc,
	// }

	data.Value = result.Value
	data.Desc = result.Desc

	//END MAPPING DATA DGA PARAMETERS

	log.Println("Model -- RESPONSE MODEL DGA PARAMETERS FIND --")
	// log.Println(data)
	// return

	response = DGAParametersResponse{
		Status:       200,
		Message:      "Data ditemukan",
		MessageLocal: "Data Valid",
	}

	return response, data

}
