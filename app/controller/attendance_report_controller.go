package controller

import (
	"os"

	"github.com/PatipatCha/jeab_ta_service/app/model"
	"github.com/PatipatCha/jeab_ta_service/app/services"
	"github.com/gofiber/fiber/v2"
)

func GetReportWeb(c *fiber.Ctx) error {
	_ = c.Query("month")
	userId := c.Query("user_id")

	findUserId := c.Query("find_user_id")
	data, _ := services.GetReportForWeb(findUserId)

	var res = model.TimeAttendanceDashboardForWebResponse{
		UserId:  userId,
		Data:    data,
		Message: "Attendance Report List",
	}

	return c.JSON(res)
}

func GetReportMobile(c *fiber.Ctx) error {
	var taReportList = []model.TimeAttendanceReportList{}
	userId := c.Query("user_id")

	month := c.Query("month")
	_, data, msg := services.GetReportForMobile(userId, month)
	if len(data) > 0 {
		taReportList = services.RecordListForMobile(data)
	} else {
		msg = os.Getenv("NO_RECORD_LISTS")
	}

	var res = fiber.Map{
		"user_id": userId,
		"message": msg,
		"data":    taReportList,
	}

	return c.JSON(res)
}

func GetReportHandler(c *fiber.Ctx) error {
	var taReportList = []model.TimeAttendanceReportList{}
	userId := c.Query("user_id")
	month := c.Query("month")
	_, data, msg := services.GetReportJGuard(userId, month)
	if len(data) > 0 {
		taReportList = services.RecordListJGuard(data)
	} else {
		msg = os.Getenv("NO_RECORD_LISTS")
	}

	var res = fiber.Map{
		"user_id": userId,
		"message": msg,
		"data":    taReportList,
	}

	return c.JSON(res)
}
