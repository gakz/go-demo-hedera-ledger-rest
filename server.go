package main

import (
	"github.com/droatl2000/demo-ledger-rest/controller"
	"github.com/droatl2000/demo-ledger-rest/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	vehicleEventService        service.VehicleEventService           = service.NewEvent()
	vehicleEventController     controller.VehicleEventController     = controller.NewEvent(vehicleEventService)
	verifiedServicerService    service.VerifiedServicerService       = service.NewServicer()
	verifiedServicerController controller.VerifiedServicerController = controller.NewServicer(verifiedServicerService)
)

func main() {
	router := gin.Default()
	// Enable CORS for requests UI domain (port)
	router.Use(cors.Default())

	router.GET("/vehicleEvents/:vin", func(ctx *gin.Context) {
		ctx.JSON(200, vehicleEventController.FindByVin(ctx.Param("vin")))
	})

	router.GET("/vehicleEvents", func(ctx *gin.Context) {
		ctx.JSON(200, vehicleEventController.FindAll())
	})

	router.POST("/vehicleEvents", func(ctx *gin.Context) {
		ctx.JSON(200, vehicleEventController.Save(ctx))
	})

	router.GET("/verifiedServicers", func(ctx *gin.Context) {
		ctx.JSON(200, verifiedServicerController.FindAll())
	})

	router.POST("/verifiedServicers", func(ctx *gin.Context) {
		println(ctx.Request.GetBody)
		ctx.JSON(200, verifiedServicerController.Save(ctx))
	})

	router.Run(":8082")

}
