package helpers

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
	"sync"

	Config "3party/config"
	Constants "3party/constants"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	// CharacterSet consists of 62 characters [0-9][A-Z][a-z].
	Base         = 62
	CharacterSet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func PubLoggingCloud(requestLog string, responseLog string, dataLog string) {
	log.Println("REQUEST LOG : ", requestLog)
	log.Println("RESPONSE LOG : ", responseLog)
	log.Println("DATA LOG : ", dataLog)

	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup

	dLogStatus := make(chan int, 2)
	dLogMessage := make(chan string, 2)

	log.Println("Publish Channel : ", Constants.ChannelLogging3PartyAPI)

	wg.Add(1)
	go func() {
		logStatus, logMessage := LoggingCloudPubSub(&wg, dataLog, Constants.ChannelLogging3PartyAPI)
		dLogStatus <- logStatus
		dLogMessage <- logMessage

		// log.Println("Service ** RESULT SERVICE PROCESS LOGGING CLOUD **")
		// log.Println(logStatus)
		// log.Println(logMessage)
	}()

	var statusLog = <-dLogStatus
	var messageLog = <-dLogMessage
	log.Println("Service ** RESULT SERVICE OPENAPI LOGGING CLOUD **")
	log.Println(statusLog)
	log.Println(messageLog)
}

func GcmEncode(text string) (result string) {
	secretKey, _ := Config.GCMSecretKey()
	nonce, _ := Config.GCMNonce()

	textByte := []byte(text)

	secretKeyByte := []byte(secretKey)
	block, _ := aes.NewCipher(secretKeyByte)

	nonceByte := []byte(nonce)

	log.Println("PLAIN TEXT : ", text)
	log.Println("SECRET KEY : ", secretKey)
	log.Println("NONCE : ", nonce)
	log.Println("PLAIN TEXT BYTE : ", textByte)
	log.Println("SECRET KEY BYTE : ", secretKeyByte)
	log.Println("NONCE BYTE : ", nonceByte)

	aes, _ := cipher.NewGCM(block)

	cipherText := aes.Seal(nil, nonceByte, textByte, nil)
	clearText := append(nonceByte, cipherText...)

	result = base64.StdEncoding.EncodeToString(clearText)

	return result
}

func GcmDecode(text string) (result string) {
	secretKey, _ := Config.GCMSecretKey()
	nonce, _ := Config.GCMNonce()

	cipherText, _ := base64.StdEncoding.DecodeString(text)

	// remove 0 - 11 index (nonce)
	lenChipperText := len(cipherText)
	clearText := cipherText[12:lenChipperText]

	secretKeyByte := []byte(secretKey)
	block, _ := aes.NewCipher(secretKeyByte)

	nonceByte := []byte(nonce)

	log.Println("CIPHER TEXT : ", text)
	log.Println("CIPHER BYTE : ", cipherText)
	log.Println("SECRET KEY : ", secretKey)
	log.Println("SECRET KEY BYTE : ", secretKeyByte)
	log.Println("NONCE : ", nonce)
	log.Println("NONCE BYTE : ", nonceByte)

	aes, _ := cipher.NewGCM(block)

	plainText, _ := aes.Open(nil, nonceByte, clearText, nil)

	result = string(plainText)

	return result
}

func StringInt(text string) (result int, status int) {
	result, _ = strconv.Atoi(text)

	//fmt.Println("---STRING TO INT---")
	//fmt.Println(result)
	//return

	if result != 0 {
		return result, 200
	}

	return 0, 400
}

func FloatString(text float64) (result string, status int) {
	result = fmt.Sprintf("%.2f", text)

	//fmt.Println("---FLOAT TO STRING---")
	//fmt.Println(result)
	//return

	if result != "" {
		return result, 200
	}

	return "", 400
}

func StringFloat64(text string) (result float64, status int) {
	result, _ = strconv.ParseFloat(text, 64)

	//fmt.Println("---STRING TO INT---")
	//fmt.Println(result)
	//return

	if result != 0 {
		return result, 200
	}

	return 0, 400
}

func IntString(text int) (result string, status int) {
	result = strconv.Itoa(text)

	//fmt.Println("---STRING TO INT---")
	//fmt.Println(result)
	//return

	if result != "" {
		return result, 200
	}

	return "", 400
}

func BoolString(text bool) (result string, status int) {
	result = strconv.FormatBool(text)

	//fmt.Println("---STRING TO INT---")
	//fmt.Println(result)
	//return

	if result != "" {
		return result, 200
	}

	return "", 400
}

func FileSize(text string) (result int, status int) {
	l := len(text)

	// count how many trailing '=' there are (if any)
	eq := 0
	if l >= 2 {
		if text[l-1] == '=' {
			eq++
		}
		if text[l-2] == '=' {
			eq++
		}

		l -= eq
	}

	result = (l*3 - eq) / 4

	return result, 200
}

func NormalizeFormatCurrency(text int) (result string, status int) {
	em := message.NewPrinter(language.English)
	result = em.Sprintf("%f", text)

	// result = fmt.Printf("%#v\n", enNumber)

	return result, 200
}

func NormalizeBirtDate(text string) (result string, status int) {
	result = strings.Replace(text, "T00:00:00Z", "", 1)

	return result, 200
}

func NormalizeDate(text string) (result string, status int) {
	result = strings.Replace(text, "T00:00:00Z", "", 1)

	return result, 200
}

func NormalizeMobileNumber(text string) (result string, status int) {
	if text == "" {
		return "", 400
	}

	result = strings.Replace(text, "+62", "", -1)

	return result, 200
}

func ConvertRomawi(content int) (result string, status int) {
	if content == 0 {
		return result, 400
	}

	switch content {
	case 1:
		result = "I"
	case 2:
		result = "II"
	case 3:
		result = "III"
	case 4:
		result = "IV"
	case 5:
		result = "V"
	case 6:
		result = "VI"
	case 7:
		result = "VII"
	case 8:
		result = "VIII"
	case 9:
		result = "IX"
	case 10:
		result = "X"
	case 11:
		result = "XI"
	case 12:
		result = "XII"
	}

	return result, 200
}

func RemoveRedis(redisKey string, redisDB string) (result int) {
	// open redis connection
	ctx := context.Background()
	redisClient := Config.InitRedisConnection(redisDB)

	redisClient.Del(ctx, redisKey)

	return 200
}

func NormalizeExpiredLimitDateFormat(text string) (result string, status int) {
	strSplit := strings.Split(text, "/")
	if len(strSplit) != 3 {
		return "", 400
	}
	result = fmt.Sprintf("%s-%s-%s", strSplit[2], strSplit[1], strSplit[0])

	return result, 200
}

func NormalizeBase64String(base64 string, mimeType string) (result string, status int) {
	result = "data:"
	switch mimeType {
	case "image/jpeg":
		result += "image/jpeg"
	case "image/png":
		result += "image/png"
	}

	result += ";base64," + base64

	return result, 200
}

func NormalizeDateFormat(text string) (result string, status int) {
	strSplit := strings.Split(text, "/")
	if len(strSplit) != 3 {
		return "", 400
	}
	result = fmt.Sprintf("%s-%s-%s", strSplit[2], strSplit[1], strSplit[0])

	return result, 200
}

func DecodeStringBase64(text string) (result string, status int) {
	status = 200

	decodeString, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		status = 400
		log.Println(err)
	}

	result = string(decodeString)
	return result, status
}
