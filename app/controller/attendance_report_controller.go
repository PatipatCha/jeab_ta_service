package controller

import (
	"os"

	"github.com/PatipatCha/jeab_ta_service/app/model"
	"github.com/PatipatCha/jeab_ta_service/app/services"
	"github.com/gofiber/fiber/v2"
)

func GetReport(c *fiber.Ctx) error {
	// var ta []model.TimeAttendanceEntity
	// userId := c.Params("userId")
	month := c.Query("month")
	userId := c.Query("user_id")
	resValidate, _ := services.VaildateUserId(userId)
	if !resValidate {
		output := model.TimeAttendanceResponse{
			UserId:  userId,
			Data:    nil,
			Message: os.Getenv("VAILD_USERID_NOT_FOUND"),
		}
		return c.JSON(output)
	}

	var data = services.GetReportByMonth(month)

	// res := services.GetReportNow()

	var res = model.TimeAttendanceReportResponse{
		UserId:  userId,
		Data:    data,
		Message: "Attendance Report List",
	}

	return c.JSON(res)
}

// for _, taParam := range ta {
// 	// fmt.Println(integ, " = ", spell.CheckInAT)
// 	taModel := model.TimeAttendance{CheckInDate: taParam.CheckInAT, CheckInPlace: taParam.CheckInPlace, CheckInTime: taParam.CheckInAT, CheckOutTime: taParam.CheckOutAt, CheckOutPlace: taParam.CheckOutPlace}
// }

// a := model.TimeAttendanceResponse{user_id: "", Message: "", list: taModel}

// bytes, _ := json.Marshal(ta)
// fmt.Println(string())
