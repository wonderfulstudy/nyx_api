package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"nyx_api/pkg/setting"
)

// CBC 加密 按照 golang 标准库的例子代码
// 不过里面没有填充的部分, 所以补上

// 使用 PKCS7 进行填充, iOS 也是 7
// 只要少于 256 就能放到一个 byte 中, 默认的 blockSize=16 (即采用 16*8=128, AES-128 长的密钥)
// 最少填充 1 个 byte, 如果明文刚好是 blocksize 的整数倍, 则再填充一个 blocksize
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	// 计算需要 padding 的数目, 并生成填充的文本
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PaddingLeft(ori []byte, pad byte, length int) []byte {
	if len(ori) >= length {
		return ori[:length]
	}
	pads := bytes.Repeat([]byte{pad}, length-len(ori))
	return append(pads, ori...)
}

// AES 加密, 填充秘钥 key 的 16 位, 24, 32 分别对应 AES-128, AES-192, or AES-256
func AesCBCEncrypt(rawData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// 填充原文
	blockSize := block.BlockSize()
	rawData = PKCS7Padding(rawData, blockSize)
	// 初始向量 IV 必须是唯一, 但不需要保密
	cipherText := make([]byte, blockSize+len(rawData))
	// block 大小 16
	iv := cipherText[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	// block 大小和初始向量大小一定要一致
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[blockSize:], rawData)

	return cipherText, nil
}

func AesCBCDecrypt(ciphertext, key []byte) ([]byte, error) {
	pkey := PaddingLeft(key, '0', 16)
	block, err := aes.NewCipher(pkey)
	if err != nil {
		panic(err)
	}

	blockModel := cipher.NewCBCDecrypter(block, pkey)
	plantText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, ciphertext)
	plantText = PKCS7UnPadding(plantText)

	return plantText, nil
}

func AesEncryptCBCBase64(rawData string) (string, error) {
	data, err := AesCBCEncrypt([]byte(rawData), []byte(setting.AESKey))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func AesDecryptCBCBase64(rawData string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(rawData)
	if err != nil {
		return "", err
	}
	dnData, err := AesCBCDecrypt(data, []byte(setting.AESKey))
	if err != nil {
		return "", err
	}
	return string(dnData), nil
}
