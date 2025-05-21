package main

import (
	"nyx_api/api/pubkey"
	"nyx_api/api/user"
	"nyx_api/configs"

	"github.com/gin-gonic/gin"
)



func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
