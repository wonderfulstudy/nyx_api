package main

import (
	"nyx_api/api/pubkey"
	"nyx_api/api/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	user.SetupUserRoutes(r)
	pubkey.SetupPubkeyRoutes(r)

	r.Run(":9528")
}
