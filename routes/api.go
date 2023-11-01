package routes

import (
	"github.com/PatipatCha/jeab_ta_service/app/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func SetupApiRoutes(app *fiber.App, store *session.Store) {

	apita := app.Group("/api")

	//
	// *TA Menu
	apita.Get("/get-ta-report-mockup/", controller.GetReportMockUp)
	apita.Get("/get-ta-report-web/", controller.GetReportWeb)
	apita.Get("/get-ta-report-mobile/", controller.GetReportMobile)
	apita.Post("/checkin", controller.CheckIn)
	apita.Post("/checkout", controller.CheckOut)
	//
	//

}
