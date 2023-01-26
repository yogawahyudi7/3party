package routes

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	Config "3party/config"
	Constants "3party/constants"
	pb "3party/delivery/proto/3party"
)

func OutOfServiceTime() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()
		hr, _, _ := t.Clock()

		log.Println("JAM : ", hr)

		log.Println("DAY : ", t.Day())
		log.Println("MONTH : ", int(t.Month()))
		log.Println("YEAR : ", t.Year())

		if t.Day() != 25 && int(t.Month()) != 12 && t.Year() != 2022 {
			// 23.00 - 06.00
			if hr >= 23 {
				ctx.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
					"status":  503,
					"message": Constants.MessageOutOfService,
					"desc":    "Out Of Service",
				})

				return
			}

			if hr < 6 {
				ctx.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
					"status":  503,
					"message": Constants.MessageOutOfService,
					"desc":    "Out Of Service",
				})

				return
			}
		}

		ctx.Next()
	}
}

func RouterMain() http.Handler {
	// Force log's color
	gin.DisableConsoleColor()

	// Logging to a file.
	t := time.Now()
	formatDate := t.Format("20060102")
	logJoin := []string{"logs", "/", "clients", "/", "3party", "/", "log", "-", formatDate, ".log"}
	logFile := strings.Join(logJoin, "")
	f, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.New()

	// use middleware out of service
	router.Use(OutOfServiceTime())

	// use middleware logging
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	appEnv, _ := Config.AppEnv()
	log.Println("APP ENV : ", appEnv)

	if appEnv == "DEV" {
		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"*"}, //http or https
			AllowMethods:     []string{"HEAD", "GET", "POST", "PUT", "PATCH", "OPTIONS", "DELETE"},
			AllowHeaders:     []string{"*"},
			AllowCredentials: false,
			MaxAge:           1 * time.Minute,
		}))

	} else {

		router.Use(cors.New(cors.Config{
			AllowOrigins: []string{
				"https://webview.pinang-performa.bankraya.co.id",
				"https://webview.pinang-performa.staging.rayain.net",
			}, //http or https
			AllowMethods: []string{
				"GET",
				"POST",
			},
			AllowHeaders: []string{
				"Authorization",
				"Content-Type",
				"Origin",
				"Accept-Encoding",
				"Accept-Language",
				"Host",
			},
			AllowCredentials: false,
			AllowOriginFunc: func(origin string) bool {
				return origin == "https://gitlab.com"
			},
			MaxAge: 12 * time.Hour,
		}))

	}

	// use middleware gin recovery
	router.Use(gin.Recovery())

	router.GET("/testing", func(ctx *gin.Context) {
		envronment, _ := Config.AppEnv()
		ctx.JSON(200, gin.H{
			"status":      200,
			"environment": envronment,
		})
	})

	//SETUP PORT
	portMain, _ := Config.PortClient()
	addressService, _ := Config.PortService()

	//OPEN CONNECTION GRPC SERVICE
	maxMsgSize := 1024 * 1024 * 20

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(maxMsgSize),
		grpc.MaxCallSendMsgSize(maxMsgSize),
	))

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	serviceConnecion, err := grpc.Dial(
		addressService,
		opts...,
	)
	if err != nil {
		log.Println("did not connect: ", err.Error())
	}

	log.Println("START OPEN SERVER IN ", portMain, " HIT : ", time.Now())
	log.Println("RUN SERVICE V1 IN ", addressService)

	//CONNECT TO SERVICE
	client := pb.NewThirdPartyServiceClient(serviceConnecion)

	// authorizeBasicAuth := gin.Accounts{
	// 	Constants.UserBasic: Constants.PassBasic,
	// }

	thirdPartyEndpoint := router.Group("/v/1/3party")
	{
		thirdPartyEndpoint.GET("/testing", func(ctx *gin.Context) {
			req := &pb.TestingRequest{Text: "HIT"}

			if response, err := client.Testing(ctx, req); err == nil {
				ctx.JSON(http.StatusOK, gin.H{
					"status":  200,
					"message": "Welcome To 3PARTY Pinang Performa",
					"desc":    "Behasil Konek 3PARTY Pinang Performa",
					"data":    response.Result,
				})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":  500,
					"message": Constants.ErorrGeneralMessage,
					"desc":    err.Error(),
					"data":    "ERROR",
				})
			}
		})

		thirdPartyEndpoint.POST("/sikp/verification", func(ctx *gin.Context) {
			type DataSIKP struct {
				StatusCode        int    `json:"statusCode"`
				StatusDescription string `json:"statusDescription"`
				BankCode          string `json:"bankCode,omitempty"`
				UploadDate        string `json:"uploadDate,omitempty"`
			}

			userId := ctx.DefaultPostForm("userId", "")
			ktpNumber := ctx.DefaultPostForm("ktpNumber", "")

			fmt.Println("ISER ID", userId)
			fmt.Println("KTP NUMBER", ktpNumber)

			req := &pb.VerificationSIKPRequest{
				UserId:    userId,
				KtpNumber: ktpNumber,
			}

			if response, err := client.VerificationSIKP(ctx, req); err == nil {
				if response.Status == 400 {
					ctx.JSON(http.StatusBadRequest, gin.H{
						"status":  response.GetStatus(),
						"message": response.GetMessage(),
						"desc":    response.GetMessageLocal(),
						"data":    nil,
					})
					return
				}

				if response.Status != 200 {
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"status":  response.GetStatus(),
						"message": response.GetMessage(),
						"desc":    response.GetMessageLocal(),
						"data":    nil,
					})
					return
				}

				statusCode := response.GetEmbedDataVerificationSIKP().GetStatusCode()
				statusDescription := response.GetEmbedDataVerificationSIKP().GetStatusDescription()
				data := response.GetEmbedDataVerificationSIKP().GetDataVerificationSIKP()
				bankCode := data.GetBankCode()
				uploadDate := data.GetUploadDate()

				if response.GetStatus() == 200 {
					if statusCode == 12 && statusDescription == "Data ditemukan" {
						if data.GetBankCode() == "494" {
							statusDescription = "NIK telah terdaftar di SIKP sebagai debitur Bank BRI AGRO."
							ctx.JSON(http.StatusOK, gin.H{
								// "StatusCode":        1,
								// "StatusDescription": "NIK " + ktpNumber + " telah terdaftar di SIKP sebagai debitur Bank BRI AGRO",
								"status":  200,
								"message": Constants.DataFoundMessage,
								"desc":    Constants.SuccessProccessMessage,
								"data": DataSIKP{
									StatusCode:        1,
									StatusDescription: statusDescription,
									BankCode:          bankCode,
									UploadDate:        uploadDate,
								},
							})
							return
						} else {
							statusDescription = "NIK telah terdaftar di SIKP sebagai debitur Bank Lain."
							ctx.JSON(http.StatusOK, gin.H{
								// "StatusCode":        -1,
								// "StatusDescription": "NIK " + ktpNumber + " telah terdaftar di SIKP sebagai debitur Bank Lain",
								"status":  200,
								"message": Constants.DataFoundMessage,
								"desc":    Constants.FailedProccessMessage,
								"data": DataSIKP{
									StatusCode:        -1,
									StatusDescription: statusDescription,
									BankCode:          bankCode,
									UploadDate:        uploadDate,
								},
							})
							return
						}
					} else if statusCode == 15 && statusDescription == "Data ditemukan (Penyalur Lain)" {
						statusDescription = "NIK telah terdaftar di SIKP sebagai debitur Bank Lain."
						ctx.JSON(http.StatusOK, gin.H{
							// "StatusCode":        -1,
							// "StatusDescription": "NIK " + ktpNumber + " telah terdaftar di SIKP sebagai debitur Bank Lain",
							"status":  200,
							"message": Constants.DataFoundMessage,
							"desc":    Constants.SuccessProccessMessage,
							"data": DataSIKP{
								StatusCode:        -1,
								StatusDescription: statusDescription,
								BankCode:          bankCode,
								UploadDate:        uploadDate,
							},
						})
						return
					} else if statusCode == 07 && statusDescription == "Data tidak ditemukan" {
						statusDescription = "NIK belum terdaftar di SIKP."
						ctx.JSON(http.StatusOK, gin.H{
							// "StatusCode":        0,
							// "StatusDescription": "NIK " + ktpNumber + " belum terdaftar di SIKP",
							"status":  404,
							"message": Constants.DataNotFoundMessage,
							"desc":    Constants.SuccessProccessMessage,
							"data": DataSIKP{
								StatusCode:        0,
								StatusDescription: statusDescription,
								BankCode:          bankCode,
								UploadDate:        uploadDate,
							},
						})
						return
					} else {
						statusDescription = "response Blank dari SIKP."
						ctx.JSON(http.StatusInternalServerError, gin.H{
							// "StatusCode":        -2,
							// "StatusDescription": "response Blank dari SIKP",
							"status":  500,
							"message": Constants.ErorrGeneralMessage,
							"desc":    Constants.FailedProccessMessage,
							"data": DataSIKP{
								StatusCode:        -2,
								StatusDescription: statusDescription,
								BankCode:          bankCode,
								UploadDate:        uploadDate,
							},
						})
						return
					}
				}
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":  500,
					"message": Constants.ErorrGeneralMessage,
					"desc":    err.Error(),
					"data":    nil,
				})
			}
		})

		thirdPartyEndpoint.POST("/sikp/check-plafond", func(ctx *gin.Context) {
			type DataPlafondSIKP struct {
				KtpNumber          string `json:"ktNumber"`
				Scheme             int64  `json:"scheme"`
				SchemeDescription  string `json:"schemeDescription"`
				TotalLimitDefault  int64  `json:"totalLimitDefault"`
				TotalLimit         int64  `json:"totalLimit"`
				LimitActiveDefault int64  `json:"LimitActiveDefault"`
				LimitActive        int64  `json:"LimitActive"`
				BankCode           int64  `json:"bankCode,omitempty"`
			}

			userId := ctx.DefaultPostForm("userId", "")
			ktpNumber := ctx.DefaultPostForm("ktpNumber", "")

			fmt.Println("ISER ID", userId)
			fmt.Println("KTP NUMBER", ktpNumber)

			req := &pb.CheckPlafondSIKPRequest{
				UserId:    userId,
				KtpNumber: ktpNumber,
			}

			if response, err := client.CheckPlafondSIKP(ctx, req); err == nil {
				if response.Status == 400 {
					ctx.JSON(http.StatusBadRequest, gin.H{
						"status":  response.GetStatus(),
						"message": response.GetMessage(),
						"desc":    response.GetMessageLocal(),
						"data":    nil,
					})
					return
				}

				if response.Status != 200 {
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"status":  response.GetStatus(),
						"message": response.GetMessage(),
						"desc":    response.GetMessageLocal(),
						"data":    nil,
					})
					return
				}

				if response.GetMessage() == "Data Ditemukan" {

					resultData := []DataPlafondSIKP{}

					for _, vData := range response.GetEmbedDataCheckPlafondSIKP().GetDataCheckPlafondSIKP() {

						data := DataPlafondSIKP{
							KtpNumber:          vData.GetKtpNumber(),
							Scheme:             vData.GetScheme(),
							SchemeDescription:  "",
							TotalLimitDefault:  vData.GetTotalLimitDefault(),
							TotalLimit:         vData.GetTotalLimit(),
							LimitActiveDefault: vData.GetLimitActiveDefault(),
							LimitActive:        vData.GetLimitActive(),
							BankCode:           vData.GetBankCode(),
						}

						// resultKtpNumber := vData.GetKtpNumber()
						// sheme, _ := Helpers.
						// // schemeDescription :=
						// totalLimitDefault := vData.GetTotalLimitDefault()
						// totalLimit := vData.GetTotalLimit()
						// limitActiveDefault := vData.GetLimitActiveDefault()
						// limitActive := vData.GetLimitActive()
						// bankCode := vData.GetBankCode()

						resultData = append(resultData, data)
					}
					ctx.JSON(http.StatusOK, gin.H{
						"status":  200,
						"message": Constants.DataFoundMessage,
						"desc":    Constants.SuccessProccessMessage,
						"data":    resultData,
					})
					return
				} else {
					ctx.JSON(http.StatusOK, gin.H{
						"status":  404,
						"message": Constants.DataNotFoundMessage,
						"desc":    Constants.SuccessProccessMessage,
						"data":    nil,
					})
					return

				}

			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":  500,
					"message": Constants.ErorrGeneralMessage,
					"desc":    err.Error(),
					"data":    nil,
				})
			}
		})

	}
	return router
}
