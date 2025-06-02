// Package aes provides AES-CBC encryption/decryption middleware for Gin framework
package aes

import (
	"bytes"
	"encoding/base64"
	"io"
	"io/ioutil"
	"net/http"
	"nyx_api/pkg/setting" // 导入配置包获取密钥

	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/gin-gonic/gin"
)

// aesKey 从配置中心获取AES密钥（原硬编码已改进为动态配置）
var aesKey = []byte(setting.AESKey)

// middlewareDecryptReq 返回Gin中间件函数，处理请求的AES解密流程
// 功能：
// 1. 解密URL查询参数
// 2. 解密请求体数据
// 3. 错误处理及请求终止
func middlewareDecryptReq() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 处理URL查询参数解密
		if c.Request.URL.RawQuery != "" {
			res, err := AesCbcDecryptBase64([]byte(c.Request.URL.RawQuery), aesKey)
			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				c.Abort()
				return
			}
			c.Request.URL.RawQuery = string(res)
		}

		// 读取并处理请求体
		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}
		defer c.Request.Body.Close()
		
		// 空数据直接放行
		if len(data) == 0 {
			c.Next()
			return
		}

		// 解密请求体数据
		plainBuf, err := AesCbcDecryptBase64(data, aesKey)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}

		// 替换原始请求体为解密后数据
		r := bytes.NewBuffer(plainBuf)
		rd := io.NopCloser(r)
		c.Request.Body = rd
	}
}

// EncryptWriter 加密数据并写入HTTP响应
// 参数：
// c - Gin上下文
// data - 需加密的原始数据
func EncryptWriter(c *gin.Context, data []byte) {
	cipherBuf := AesCbcEncryptBase64(data, aesKey)
	c.String(http.StatusOK, string(cipherBuf))
}

// AesCbcEncrypt AES-CBC模式加密基础方法
func AesCbcEncrypt(plainText, secretKey []byte) []byte {
	return cryptor.AesCbcEncrypt(plainText, secretKey)
}

// AesCbcDecrypt AES-CBC模式解密基础方法
func AesCbcDecrypt(cipherText, key []byte) []byte {
	return cryptor.AesCbcDecrypt(cipherText, key)
}

// AesCbcEncryptBase64 执行加密并进行Base64编码
func AesCbcEncryptBase64(plainText, secretKey []byte) (cipherTextBase64 []byte) {
	encryBytes := AesCbcEncrypt(plainText, secretKey)
	cipherTextBase64 = make([]byte, base64.StdEncoding.EncodedLen(len(encryBytes)))
	base64.StdEncoding.Encode(cipherTextBase64, encryBytes)
	return
}

// AesCbcDecryptBase64 执行Base64解码并进行解密
func AesCbcDecryptBase64(cipherTextBase64, key []byte) (res []byte, err error) {
	plainTextBytes := make([]byte, base64.StdEncoding.DecodedLen(len(cipherTextBase64)))
	n, err := base64.StdEncoding.Decode(plainTextBytes, cipherTextBase64)
	if err != nil {
		return
	}
	res = AesCbcDecrypt(plainTextBytes[:n], key)
	return
}