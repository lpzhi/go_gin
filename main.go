package main

import (
	"fmt"
	"gin/pkg/setting"
	"gin/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	//router.Run(":8000")
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}


	s.ListenAndServe()
}