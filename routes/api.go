package routes

import (
	"github.com/PatipatCha/jeab_ta_service/app/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func SetupApiRoutes(app *fiber.App, store *session.Store) {

	apita := app.Group("/api")
	v1 := apita.Group("/v1")

	// *TA Menu
	jguard := v1.Group("/jguard")
	jguard.Get("/status", controller.CheckStatusHandler)
	jguard.Get("/report", controller.GetReportHandler)
	jguard.Post("/checkin", controller.CheckInHandler)
	jguard.Post("/checkout", controller.CheckOutHandler)
	//
	jcenter := v1.Group("/jcenter")
	jcenter.Get("/report/", controller.GetReportWeb)
	//

}
