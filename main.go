package main

import (
	"fmt"
	"net/http"
	"nyx_api/pkg/setting"
	"nyx_api/routers"
)

func main() {
	// var src = "CPPE-cd@2019"
	// en, err := aes.AesEncryptCBCBase64(src)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("加密结果：", en)

	// de, err := aes.AesDecryptCBCBase64(en)
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
