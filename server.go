package main

import (
	"github.com/droatl2000/demo-ledger-rest/controller"
	"github.com/droatl2000/demo-ledger-rest/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	vehicleActionService    service.VehicleActionService       = service.New()
	vehicleActionController controller.VehicleActionController = controller.New(vehicleActionService)
)

func main() {
	router := gin.Default()
	// Enable CORS for requests UI domain (port)
	router.Use(cors.Default())

	router.GET("/vehicleActions/:vin", func(ctx *gin.Context) {
		ctx.JSON(200, vehicleActionController.FindByVin(ctx.Param("vin")))
	})

	router.GET("/vehicleActions", func(ctx *gin.Context) {
		ctx.JSON(200, vehicleActionController.FindAll())
	})

	router.POST("/vehicleActions", func(ctx *gin.Context) {
		ctx.JSON(200, vehicleActionController.Save(ctx))
	})

	router.Run(":8082")

}
