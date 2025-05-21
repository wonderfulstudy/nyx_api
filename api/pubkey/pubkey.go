package pubkey

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"

	"github.com/gin-gonic/gin"
)

var privateKey *rsa.PrivateKey

func init() {
	// 生成 RSA 密钥对
	var err error
	privateKey, err = rsa.GenerateKey(rand.Reader, 2048) // 2048 位密钥
	fmt.Println(privateKey)
	if err != nil {
		panic("Failed to generate RSA key pair: " + err.Error())
	}
}

func SetupPubkeyRoutes(r *gin.Engine) {
	// 获取公钥的接口
	r.GET("/pubkey", func(c *gin.Context) {
		// 将公钥转换为 PEM 格式
		publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to marshal public key"})
			return
		}

		publicKeyPEM := pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: publicKeyBytes,
		})

		c.JSON(200, gin.H{"publicKey": string(publicKeyPEM)})
	})

	// 解密接口（示例）
	r.POST("/decrypt", func(c *gin.Context) {
		var request struct {
			EncryptedData string `json:"encryptedData"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		decryptedData, err := decryptData(request.EncryptedData)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to decrypt data"})
			return
		}

		c.JSON(200, gin.H{"decryptedData": decryptedData})
	})
}

func decryptData(encryptedData string) (string, error) {
	// 解码 Base64 加密数据
	decodedData, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}

	// 使用私钥解密
	decryptedData, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, decodedData)
	if err != nil {
		return "", err
	}

	return string(decryptedData), nil
}
