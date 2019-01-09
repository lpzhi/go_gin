package main

import (
	"fmt"
	"gin/pkg/setting"
	"gin/routers"
	"github.com/DeanThompson/ginpprof"
	"net/http"
)

func main() {
	router := routers.InitRouter()
	ginpprof.Wrap(router)
	//router.Run(":8000")
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.WriteTimeout,
		WriteTimeout:   setting.ServerSetting.ReadTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	//fmt.Println(setting.ServerSetting)
	//log.Println(s.ListenAndServe())
	s.ListenAndServe()
}