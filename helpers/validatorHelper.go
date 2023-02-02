package helpers

import (
	"fmt"
	"log"
	"regexp"

	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"

	// "github.com/go-playground/validator/v10"
	"gopkg.in/go-playground/validator.v9"
	id_translations "gopkg.in/go-playground/validator.v9/translations/id"
)

var (
	validate  *validator.Validate
	uni       *ut.UniversalTranslator
	validates *validator.Validate
)

type ValidateStruct struct {
	Name string `validate:"name"`
}

type ValidateTextStruct struct {
	Text string `validate:"text"`
}

type ValidateDateStruct struct {
	DateString string `validate:"dateString"`
}

type ValidateNumericStruct struct {
	Integer string `validate:"integer"`
}

type ValidateAplhaNumericStruct struct {
	Text string `validate:"text"`
}

// var validate *validator.Validate

// custom validator
func ValidateName(fl validator.FieldLevel) bool {
	regexExp := regexp.MustCompile(`^[a-z A-Z,.\-]+$`)
	checking := regexExp.MatchString(fl.Field().String())

	log.Println("String Name : ", fl.Field().String())
	log.Println("Result Name : ", checking)

	return checking
}

func ValidateText(fl validator.FieldLevel) bool {
	regexExp := regexp.MustCompile(`^[a-z A-Z]+$`)
	checking := regexExp.MatchString(fl.Field().String())

	log.Println("String : ", fl.Field().String())
	log.Println("Result : ", checking)

	return checking
}

func ValidateAlphaNumeric(fl validator.FieldLevel) bool {
	regexExp := regexp.MustCompile(`^[A-Z0-9]+$`)
	checking := regexExp.MatchString(fl.Field().String())

	log.Println("String : ", fl.Field().String())
	log.Println("Result : ", checking)

	return checking
}

func ValidateDate(fl validator.FieldLevel) bool {
	regexExp := regexp.MustCompile(`^[0-9][0-9][0-9][0-9][-][0-9][0-9][-][0-9][0-9]$`)
	checking := regexExp.MatchString(fl.Field().String())

	log.Println("String Date : ", fl.Field().String())
	log.Println("Result Date : ", checking)

	return checking
}

func ValidateRtRw(fl validator.FieldLevel) bool {
	regexExp := regexp.MustCompile(`^[0-9][0-9][0-9]$`)
	checking := regexExp.MatchString(fl.Field().String())

	log.Println("String : ", fl.Field().String())
	log.Println("Result : ", checking)

	return checking
}

func ValidateNumeric(fl validator.FieldLevel) bool {
	regexExp := regexp.MustCompile(`^[0-9]+$`)
	checking := regexExp.MatchString(fl.Field().String())

	log.Println("String Numeric : ", fl.Field().String())
	log.Println("Result Numeric : ", checking)

	return checking
}

func ValidateAlphaNumericDisbursementCode(fl validator.FieldLevel) bool {
	regexExp := regexp.MustCompile(`^[A-Z0-9]+$`)
	checking := regexExp.MatchString(fl.Field().String())

	log.Println("String : ", fl.Field().String())
	log.Println("Result : ", checking)

	return checking
}

func ValidateFullText(fl validator.FieldLevel) bool {
	regexExp := regexp.MustCompile(`^[a-z A-Z 0-9,.]+$`)
	checking := regexExp.MatchString(fl.Field().String())

	log.Println("String : ", fl.Field().String())
	log.Println("Result : ", checking)

	return checking
}

func ValidatorUserId(userId string) (status int, errorMessage string, errorMessageLocal string) {
	status = 200

	validate = validator.New()

	if userId == "" {
		status = 400
		errorMessage = "Maaf, Parameter user id belum diisi"
		errorMessageLocal = "Parameter userId belum diisi"

		return status, errorMessage, errorMessageLocal

	}

	log.Println("User Id :", userId)
	//helper validator
	validate.RegisterValidation("integer", ValidateNumeric)
	str := ValidateNumericStruct{Integer: userId}

	errs := validate.Struct(str)
	if errs != nil {
		log.Println("Error Id : ", errs.Error())
		status = 400
		errorMessage = "Maaf, Parameter user id hanya boleh menggunakan angka"
		errorMessageLocal = "Format Parameter userId tidak sesuai"

		return status, errorMessage, errorMessageLocal

	}

	//validasi max min
	errs2 := validate.Var(userId, "min=6,max=6")

	if errs2 != nil {
		log.Println("Error Password : ", errs2.Error())

		status = 400
		errorMessage = "Maaf, ID User harus 6 digit. Format : 123XXX"
		errorMessageLocal = "Parameter userId tidak sesuai"

		return status, errorMessage, errorMessageLocal
	}

	return status, errorMessage, errorMessageLocal
}

func ValidatorKtpNumber(text string) (status int, errorMessage string, errorMessageLocal string) {
	status = 200

	validate = validator.New()

	if text == "" {
		status = 400
		errorMessage = "Nomor KTP tidak boleh kosong"
		errorMessageLocal = "Parameter ktpNumber kosong"

		return status, errorMessage, errorMessageLocal

	}

	//validasi numeric

	//helper validator numeric regex
	validate.RegisterValidation("integer", ValidateNumeric)
	str := ValidateNumericStruct{Integer: text}

	errsValidInt := validate.Struct(str)
	if errsValidInt != nil {

		log.Println("ERR : ", errsValidInt.Error())

		status = 400
		errorMessage = "Nomor KTP hanya boleh menggunakan angka"
		errorMessageLocal = "Format Parameter ktpNumber tidak sesuai"

		return status, errorMessage, errorMessageLocal
	}

	intText, _ := StringInt(text)

	//validasi gagal int
	if intText == 0 {
		status = 400
		errorMessage = "Nomor KTP hanya boleh menggunakan angka"
		errorMessageLocal = "Format Parameter ktpNumber tidak sesuai"

		return status, errorMessage, errorMessageLocal
	}

	errsValid := validate.Var(intText, "numeric")
	if errsValid != nil {
		status = 400
		errorMessage = "Nomor KTP hanya boleh menggunakan angka"
		errorMessageLocal = "Format Parameter ktpNumber tidak sesuai"

		return status, errorMessage, errorMessageLocal

	}

	//validasi max min
	errsValid2 := validate.Var(text, "min=16,max=16")
	if errsValid2 != nil {
		status = 400
		errorMessage = "Nomor KTP tidak 16 digit"
		errorMessageLocal = "Format Parameter ktpNumber tidak sesuai"

		return status, errorMessage, errorMessageLocal

	}

	return status, errorMessage, errorMessageLocal
}

func CustomeValidateRequest(model interface{}) (status int64, errMsg, errMsgLocal string) {
	status = 200

	validate := validator.New()
	validate.RegisterValidation("name", IsName)
	validate.RegisterValidation("text", IsText)
	validate.RegisterValidation("float", IsFloat)
	validate.RegisterValidation("dateString", IsDateTime)
	validate.RegisterValidation("numnonzero", IsNumberMoreThanZero)
	validate.RegisterValidation("alnumspecial", IsTextAlphaNumSpecial)
	validate.RegisterValidation("address", IsAddress)

	err := validate.Struct(model)
	id := id.New()
	uni = ut.New(id, id)
	trans, _ := uni.GetTranslator("id")

	id_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		for _, errFiled := range err.(validator.ValidationErrors) {
			fmt.Println("========= INI ERROR FILED : ", errFiled.Field())
			fmt.Println("========= INI ERROR FILED : ", errFiled.Value())
			switch errFiled.Tag() {
			case "required":
				AddCustomeErrorMessage(errFiled.Tag(), "Belum di isi", validate, trans)

				status = 400
				errMsg = errFiled.Translate(trans)
				errMsgLocal = errFiled.Translate(trans)

				return status, errMsg, errMsgLocal

			case "numeric":
				AddCustomeErrorMessage(errFiled.Tag(), "Harus Berupa Numeric", validate, trans)

				status = 400
				errMsg = errFiled.Translate(trans)
				errMsgLocal = errFiled.Translate(trans)

				return status, errMsg, errMsgLocal

			case "number":
				AddCustomeErrorMessage(errFiled.Tag(), "Harus Berupa Angka dan tidak kecil dari 0.", validate, trans)

				status = 400
				errMsg = errFiled.Translate(trans)
				errMsgLocal = "Parameter " + errFiled.Field() + " Tidak Sesuai"

				return status, errMsg, errMsgLocal

			case "name":
				AddCustomeErrorMessage(errFiled.Tag(), "Harus Berupa huruf", validate, trans)

				status = 400
				errMsg = errFiled.Translate(trans)
				errMsgLocal = errFiled.Translate(trans)

				return status, errMsg, errMsgLocal

			case "float":
				AddCustomeErrorMessage(errFiled.Tag(), "Harus berupa Angka atau Angka Decimal", validate, trans)

				status = 400
				errMsg = errFiled.Translate(trans)
				errMsgLocal = "Parameter " + errFiled.Field() + " Tidak Sesuai"

				return status, errMsg, errMsgLocal

			case "text":
				AddCustomeErrorMessage(errFiled.Tag(), "Harus berupa huruf", validate, trans)

				status = 400
				errMsg = errFiled.Translate(trans)
				errMsgLocal = errFiled.Translate(trans)

				return status, errMsg, errMsgLocal

			case "address":
				AddCustomeErrorMessage(errFiled.Tag(), "Tidak boleh memuat karakter spesial", validate, trans)

				status = 400
				errMsg = errFiled.Translate(trans)
				errMsgLocal = errFiled.Translate(trans)

				return status, errMsg, errMsgLocal

			case "alnumspecial":
				AddCustomeErrorMessage(errFiled.Tag(), "Harus berupa huruf, karakter spesial (-), dan angka. Contoh: (abc-123)", validate, trans)

				status = 400
				errMsg = errFiled.Translate(trans)
				errMsgLocal = errFiled.Translate(trans)

				return status, errMsg, errMsgLocal

			case "dateString":
				AddCustomeErrorMessage(errFiled.Tag(), "Format tanggal tidak sesuai. (Format : Y-m-d)", validate, trans)

				status = 400
				errMsg = errFiled.Translate(trans)
				errMsgLocal = errFiled.Translate(trans)

				return status, errMsg, errMsgLocal

			case "numnonzero":
				AddCustomeErrorMessage(errFiled.Tag(), "Harus Berupa Angka dan Tidak Sama dengan 0", validate, trans)

				status = 400
				errMsg = errFiled.Translate(trans)
				errMsgLocal = errFiled.Translate(trans)

				return status, errMsg, errMsgLocal

			case "ne": // ne = not equal
				if errFiled.Field() == "UserId" || errFiled.Field() == "PrakarsaId" || errFiled.Field() == "PartnershipMemberId" || errFiled.Field() == "PipelineId" {
					AddCustomeErrorMessage(errFiled.Tag(), "Tidak Ditemukan", validate, trans)

					status = 400
					errMsg = errFiled.Translate(trans)
					errMsgLocal = errFiled.Translate(trans)

					return status, errMsg, errMsgLocal
				}

			}
		}

	}
	return status, errMsg, errMsgLocal
}

func IsFloat(fl validator.FieldLevel) bool {
	regexExp := regexp.MustCompile(`^(\d|[1-9]+\d*|\.\d+|0\.\d+|[1-9]+\d*\.\d+)$`)
	checking := regexExp.MatchString(fl.Field().String())

	log.Println("String float : ", fl.Field().String())
	log.Println("Result float : ", checking)

	return checking
}

func IsDateTime(fl validator.FieldLevel) bool {
	regexExp := regexp.MustCompile(`^((19|20)\d\d)[-](0?[1-9]|1[012])[-](0?[1-9]|[12][0-9]|3[01])$`)
	return regexExp.MatchString(fl.Field().String())
}

func IsTextAlphaNumSpecial(fl validator.FieldLevel) bool {
	regexExp := regexp.MustCompile(`^[a-zA-Z]+\-[0-9]+$`)
	return regexExp.MatchString(fl.Field().String())
}

func IsText(fl validator.FieldLevel) bool {
	regexExp := regexp.MustCompile(`^[a-z A-Z \- ]*$`)
	return regexExp.MatchString(fl.Field().String())
}

func IsName(fl validator.FieldLevel) bool {
	regexExp := regexp.MustCompile(`^[a-z A-Z]*$`)
	return regexExp.MatchString(fl.Field().String())
}

func IsAddress(fl validator.FieldLevel) bool {
	regexExp := regexp.MustCompile(`^[a-z A-Z 0-9\. ]*$`)
	checking := regexExp.MatchString(fl.Field().String())

	log.Println("String Value :", fl.Field().String())
	log.Println("String Result :", checking)
	return checking
}

func IsNumberMoreThanZero(fl validator.FieldLevel) bool {
	reg := regexp.MustCompile("^[1-90-9]*$")
	return reg.MatchString(fl.Field().String())
}

func AliasString(str string) string {
	switch str {
	case "Year":
		return "Tahun"
	case "Month":
		return "Bulan"
	case "SkppDate":
		return "Tanggal SKPP"
	case "UserId":
		return "User"
	case "id":
		return "ID"
	case "PipelineId":
		return "Pipeline ID"
	case "PartnershipCode":
		return "Kode Partnership"
	case "BranchCode":
		return "Kode Cabang"
	case "PrakarsaId":
		return "Prakarsa"
	case "partnershipMemberId":
		return "Member ID Partner"
	case "EstimationLoan":
		return "Estimasi Pinjaman"
	case "RecordCount":
		return "Limit Data"
	case "StartIndex":
		return "Offset Data"

	default:
		return str
	}
}

func AddCustomeErrorMessage(Tag string, messageError string, validate *validator.Validate, trans ut.Translator) {
	validate.RegisterTranslation(Tag, trans, func(ut ut.Translator) error {
		return ut.Add(Tag, "{0} "+messageError, true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		var t string
		switch fe.Field() {
		case "UserId", "PrakarsaId", "PartnershipMemberId", "PipelineId", "EstimationLoan", "RecordCount", "StartIndex":
			t, _ = ut.T(Tag, "Maaf Parameter "+AliasString(fe.Field()))
		default:
			t, _ = ut.T(Tag, fe.Field())
		}

		return t
	})
}
