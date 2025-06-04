package main

import (
	"fmt"
	"net/http"
	"nyx_api/pkg/setting"
	"nyx_api/routers"
)

func main() {
	// var src = "{admin: 123, user: 456}"
	// en, err := aes.AesEncryptCBCBase64(src)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("加密结果：", string(en))

	// de, err := aes.AesDecryptCBCBase64("1830b0b2d2faf4cfe24f1591f41fb0bcb909087905b4ac93acb2c5ec5dafa62a5b5a11136c569c06304ddfec38b1559d")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("解密结果：", de)
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
