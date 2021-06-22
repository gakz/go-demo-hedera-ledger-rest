package main

import (
	"github.com/droatl2000/demo-ledger-rest/controller"
	"github.com/droatl2000/demo-ledger-rest/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	ledgerMessageService    service.LedgerMessageService       = service.New()
	ledgerMessageController controller.LedgerMessageController = controller.New(ledgerMessageService)
)

func main() {
	router := gin.Default()
	// Enable CORS for requests UI domain (port)
	router.Use(cors.Default())

	router.GET("/ledgerMessages/:vin", func(ctx *gin.Context) {
		ctx.JSON(200, ledgerMessageController.FindByVin(ctx.Param("vin")))
	})

	router.GET("/ledgerMessages", func(ctx *gin.Context) {
		ctx.JSON(200, ledgerMessageController.FindAll())
	})

	router.POST("/ledgerMessages", func(ctx *gin.Context) {
		ctx.JSON(200, ledgerMessageController.Save(ctx))
	})

	router.Run(":8082")

}
