package config

import (
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

type LoggingCloudPubSub struct {
	Status          string
	TypeLog         string
	Endpoint        string
	UserId          string
	PrakarsaId      string
	PartnershipCode string
	ActionDate      string
	Description     string
	DataRequest     string
	DataResponse    string
}

func AppEnv() (name string, status int) {
	result := os.Getenv("APP_ENV")

	return result, 200
}

func AppName() (name string, status int) {
	result := os.Getenv("APP_NAME")

	return result, 200
}

func PortClient() (name string, status int) {
	result := os.Getenv("PORT_CLIENT")

	return result, 200
}

func PortService() (name string, status int) {
	result := os.Getenv("PORT_SERVICE")

	return result, 200
}

func GCMSecretKey() (result string, status int) {
	result = os.Getenv("GCM_SECRETKEY") // gcm secret

	return result, 200
}

func GCMNonce() (result string, status int) {
	result = os.Getenv("GCM_NONCE") // gcm nonce

	return result, 200
}

func MediaPath() (result string, status int) {
	url := os.Getenv("MEDIA_PATH")

	// fmt.Println("--MEDIA PATH--")
	// fmt.Println(url)

	return url, 200
}

func MediaBucket() (result string, status int) {
	url := os.Getenv("MEDIA_BUCKET")

	// fmt.Println("--MEDIA BUCKET--")
	// fmt.Println(url)

	return url, 200
}

func CredentialGoogleApplication() (result string, status int) {
	credentials := os.Getenv("GOOGLE_APPLICATION_PIKRO_CREDENTIALS")

	// fmt.Println("--GOOGLE APPLICATION CREDENTIALS--")
	// fmt.Println(url)

	return credentials, 200
}

func ProjectIdGoogleApplication() (result string, status int) {
	projectId := os.Getenv("GOOGLE_APPLICATION_PIKRO_PROJECTID")

	// fmt.Println("--GOOGLE APPLICATION PROJECT ID--")
	// fmt.Println(projectId)

	return projectId, 200
}

func StringToInt(text string) (result int, status int) {
	result, _ = strconv.Atoi(text)

	return result, 200
}

func RedisUrl() (name string, status int) {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")

	stringUrl := []string{host, ":", port}
	result := strings.Join(stringUrl, "")

	// fmt.Println("--REDIS URL--")
	// fmt.Println(result)

	return result, 200
}

func RedisPassword() (name string, status int) {
	result := os.Getenv("REDIS_PASSWORD")

	// fmt.Println("--REDIS PASSWORD--")
	// fmt.Println(result)

	return result, 200
}

func RedisDB() (name int, status int) {
	dbRedis := os.Getenv("REDIS_DB")

	result, _ := StringToInt(dbRedis)

	// fmt.Println("--REDIS DB--")
	// fmt.Println(result)

	return result, 200
}

func RedisDBOTP() (result string) {
	result = os.Getenv("REDIS_DB_OTP") // redis db

	return result
}

func RedisDBDLoanDocument() (result string) {
	result = os.Getenv("REDIS_DB_LOAN_DOCUMENT") // redis db loan document

	return result
}

func InitDb() (*gorm.DB, error) {
	// Conection Pikro
	dbHostPikro := os.Getenv("PINANGMIKRO_HOST") // db host pikro
	dbUserPikro := os.Getenv("PINANGMIKRO_USER") // db host pikro
	dbPassPikro := os.Getenv("PINANGMIKRO_PASS") // db host pikro
	dbNamePikro := os.Getenv("PINANGMIKRO_DB")   // db host pikro

	dsnString := []string{"sqlserver://", dbUserPikro, ":", dbPassPikro, "@", dbHostPikro, "?database=", dbNamePikro} // connection string pikro
	dsn := strings.Join(dsnString, "")

	// fmt.Println("--CONNECTION STRING--")
	// fmt.Println(dsn)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{}) // open connection pinang mikro db

	if err != nil {
		return nil, err
	}

	db.Logger.LogMode(logger.Info)

	db.Use(dbresolver.Register(dbresolver.Config{}).SetMaxIdleConns(50).SetMaxOpenConns(100).SetConnMaxLifetime(time.Hour))

	return db, nil
}

func InitPartnershipDb() (*gorm.DB, error) {
	// err := godotenv.Load()
	// if err != nil {
	// 	return nil, err
	// }

	// Conection Partnership
	dbHostPartnership := os.Getenv("PARTNERSHIP_HOST") // db host partnership
	dbUserPartnership := os.Getenv("PARTNERSHIP_USER") // db user partnership
	dbPassPartnership := os.Getenv("PARTNERSHIP_PASS") // db password partnership
	dbNamePartnership := os.Getenv("PARTNERSHIP_DB")   // db name partnership

	dsnStringPartnership := []string{"sqlserver://", dbUserPartnership, ":", dbPassPartnership, "@", dbHostPartnership, "?database=", dbNamePartnership} // connection string partnership
	dsnPartnership := strings.Join(dsnStringPartnership, "")
	// fmt.Println("--CONNECTION STRING--")
	// fmt.Println(dsn)

	db, err := gorm.Open(sqlserver.Open(dsnPartnership), &gorm.Config{}) // open connection pinang mikro partnership db

	if err != nil {
		return nil, err
	}

	db.Logger.LogMode(logger.Info)

	db.Use(dbresolver.Register(dbresolver.Config{}).SetMaxIdleConns(50).SetMaxOpenConns(100).SetConnMaxLifetime(time.Hour))

	return db, nil
}

func InitDigiagriDb() (*gorm.DB, error) {
	// Conection Pikro
	dbHostPikro := os.Getenv("DIGIAGRI_HOST") // db host pikro
	dbUserPikro := os.Getenv("DIGIAGRI_USER") // db host pikro
	dbPassPikro := os.Getenv("DIGIAGRI_PASS") // db host pikro
	dbNamePikro := os.Getenv("DIGIAGRI_DB")   // db host pikro

	dsnString := []string{"sqlserver://", dbUserPikro, ":", dbPassPikro, "@", dbHostPikro, "?database=", dbNamePikro, "&encrypt=disable"} // connection string pikro
	dsn := strings.Join(dsnString, "")

	// fmt.Println("--CONNECTION STRING--")
	// fmt.Println(dsn)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{}) // open connection pinang mikro db

	if err != nil {
		return nil, err
	}

	db.Logger.LogMode(logger.Info)

	db.Use(dbresolver.Register(dbresolver.Config{}).SetMaxIdleConns(50).SetMaxOpenConns(100).SetConnMaxLifetime(time.Hour))

	return db, nil
}

func InitAgroAPIDb() (*gorm.DB, error) {
	// Conection Pikro
	dbHostPikro := os.Getenv("AGROAPI_HOST") // db host pikro
	dbUserPikro := os.Getenv("AGROAPI_USER") // db host pikro
	dbPassPikro := os.Getenv("AGROAPI_PASS") // db host pikro
	dbNamePikro := os.Getenv("AGROAPI_DB")   // db host pikro

	dsnString := []string{"sqlserver://", dbUserPikro, ":", dbPassPikro, "@", dbHostPikro, "?database=", dbNamePikro, "&encrypt=disable"} // connection string pikro
	dsn := strings.Join(dsnString, "")

	// fmt.Println("--CONNECTION STRING--")
	// fmt.Println(dsn)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{}) // open connection pinang mikro db

	if err != nil {
		return nil, err
	}

	db.Logger.LogMode(logger.Info)

	db.Use(dbresolver.Register(dbresolver.Config{}).SetMaxIdleConns(50).SetMaxOpenConns(100).SetConnMaxLifetime(time.Hour))

	return db, nil
}

func InitRedisConnection(redisDb string) (rdb *redis.Client) {
	redisUrl, _ := RedisUrl()
	redisPassword, _ := RedisPassword()
	// redisDB, _ := RedisDB()
	rDB := 0
	if redisDb != "" {
		rDB, _ = StringToInt(redisDb)
	}

	log.Println("RDB : ", rDB)

	rdb = redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: redisPassword, // no password set
		DB:       rDB,           // use default DB
	})

	return rdb
}

func LassAPIUserName() (result string, status int) {
	url := os.Getenv("LASS_API_USERNAME")

	// fmt.Println("--URL API SICD--")
	// fmt.Println(url)

	return url, 200
}

func LassAPIPassword() (result string, status int) {
	url := os.Getenv("LASS_API_PASSWORD")

	// fmt.Println("--URL API SICD--")
	// fmt.Println(url)

	return url, 200
}

func APIBERequestToken() (result string, status int) {
	url := os.Getenv("API_BE_REQUEST_TOKEN")

	// fmt.Println("--URL API CORE FACILITY URL--")
	// fmt.Println(url)

	return url, 200
}

func APIBEJWTUsername() (result string, status int) {
	url := os.Getenv("API_BE_USERNAME")

	// fmt.Println("--URL API CORE FACILITY URL--")
	// fmt.Println(url)

	return url, 200
}

func APIBEJWTPassword() (result string, status int) {
	url := os.Getenv("API_BE_PASSWORD")

	// fmt.Println("--URL API CORE FACILITY URL--")
	// fmt.Println(url)

	return url, 200
}

var TransportConfig *http.Transport = &http.Transport{
	Dial: (&net.Dialer{
		Timeout:   2 * time.Minute,
		KeepAlive: 30 * time.Second,
	}).Dial,
	// We use ABSURDLY large keys, and should probably not.
	TLSHandshakeTimeout: 60 * time.Second,
}

func SIKPWebService() (result string, status int) {
	url := os.Getenv("SIKP_WEB_SERVICE_URL")

	// fmt.Println("--URL API DATA AGRO PRESCREENING--")
	// fmt.Println(url)

	return url, 200
}
