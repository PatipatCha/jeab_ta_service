package controller

import (
	"os"

	"github.com/PatipatCha/jeab_ta_service/app/databases"
	"github.com/PatipatCha/jeab_ta_service/app/model"
	"github.com/PatipatCha/jeab_ta_service/app/services"
	"github.com/gofiber/fiber/v2"
)

func CheckIn(c *fiber.Ctx) error {
	var checkStatus = "checkin"
	var request model.TimeAttendanceCheckInRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	var res model.TimeAttendanceEntity

	request.CheckStatus = checkStatus
	resValidate := services.VaildateCheckStatus(request.UserId, checkStatus)
	if !resValidate {
		var msg = os.Getenv("VAILD_OTHER_ERROR")
		if request.CheckStatus == "checkin" {
			msg = os.Getenv("VAILD_CHECKOUT_ERROR")
		}

		output := model.TimeAttendanceResponse{
			UserId:  request.UserId,
			Data:    nil,
			Message: msg,
		}

		return c.JSON(output)
	}

	res, err := services.SaveData(request)
	if err != nil {
		return err
	}

	output := fiber.Map{
		"user_id": request.UserId,
		"data":    res,
		"message": os.Getenv("SAVE_DATA_SUCCESS"),
	}

	return c.JSON(output)
}

func CheckOut(c *fiber.Ctx) error {
	var checkStatus = "checkout"
	var request model.TimeAttendanceCheckInRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	var res model.TimeAttendanceEntity

	request.CheckStatus = checkStatus
	resValidate := services.VaildateCheckStatus(request.UserId, checkStatus)
	if !resValidate {
		var msg = os.Getenv("VAILD_OTHER_ERROR")
		if request.CheckStatus == "checkout" {
			msg = os.Getenv("VAILD_CHECKIN_ERROR")
		}

		output := model.TimeAttendanceResponse{
			UserId:  request.UserId,
			Data:    nil,
			Message: msg,
		}

		return c.JSON(output)
	}

	res, err := services.SaveData(request)
	if err != nil {
		return err
	}

	output := fiber.Map{
		"user_id": request.UserId,
		"data":    res,
		"message": os.Getenv("SAVE_DATA_SUCCESS"),
	}

	return c.JSON(output)
}

func GetReport(c *fiber.Ctx) error {
	var ta []model.TimeAttendanceEntity
	db, err := databases.ConnectDB()
	if err != nil {
		return err
	}

	// userId := c.Params("userId")
	res := db.Table("time_attendance").Find(&ta).Error
	if err := res; err != nil {
		return err
	}

	// for _, taParam := range ta {
	// 	// fmt.Println(integ, " = ", spell.CheckInAT)
	// 	taModel := model.TimeAttendance{CheckInDate: taParam.CheckInAT, CheckInPlace: taParam.CheckInPlace, CheckInTime: taParam.CheckInAT, CheckOutTime: taParam.CheckOutAt, CheckOutPlace: taParam.CheckOutPlace}
	// }

	// a := model.TimeAttendanceResponse{user_id: "", Message: "", list: taModel}

	// bytes, _ := json.Marshal(ta)
	// fmt.Println(string())

	return res
}

// var unit []models.Unit
// if err := db.Where("project_id = ?", projectID).Find(&unit).Error; err != nil {
// 	return err
// }

// func convertTimeUnix(myDateString string) int64 {
// 	myDate, err := time.Parse("2006-01-02 15:04:05 +0700 UTC", myDateString)
// 	if err != nil {
// 		panic(err)
// 	}
// 	res := myDate.Unix()

// 	return res
// }

// func TestCurrentTime(t *testing.T) {
// 	currentTime := time.Now()

// 	fmt.Println(
// 		"\n>> Current Time = ", currentTime,
// 		"\n>> Day = ", currentTime.Day(),
// 		"\n>> Month = ", strconv.Itoa(int(currentTime.Month())),
// 		"\n>> Year = ", strconv.Itoa(currentTime.Year()), "\n ",
// 	)
// }
