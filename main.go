package main

import (
	"fmt"
	"net/http"
	"nyx_api/middleware/aes"
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

	de, err := aes.AesDecryptCBCBase64("8y6ASnvXE//McWTUobRU+kQ8XVqsi3XH9N1sAJuNf44qAjdLW+qtlqPy1mLenTh5BRbcIHC/gSYM+AQ2ArPypQ+ajv1Rw7G+75nthrbclI0RaMILbxjhkyuVQTlHw0GQ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("解密结果：", de)

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
