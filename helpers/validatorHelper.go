package helpers

import (
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
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

var validate *validator.Validate

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
