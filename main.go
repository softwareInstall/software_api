package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"software_api/pkg/init_data"

	"software_api/models"
	"software_api/pkg/gredis"
	"software_api/pkg/logging"
	"software_api/pkg/oss"
	"software_api/pkg/setting"
	"software_api/pkg/util"
	"software_api/routers"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
	oss.Setup()
	init_data.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}

	//If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless
	//endless.DefaultReadTimeOut = readTimeout
	//endless.DefaultWriteTimeOut = writeTimeout
	//endless.DefaultMaxHeaderBytes = maxHeaderBytes
	//server := endless.NewServer(endPoint, routersInit)
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}
}
