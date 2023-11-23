package controller

import (
	"os"

	"github.com/PatipatCha/jeab_ta_service/app/model"
	"github.com/PatipatCha/jeab_ta_service/app/services"
	"github.com/gofiber/fiber/v2"
)

func CheckInHandler(c *fiber.Ctx) error {
	var checkStatus = "checkin"
	var request model.TimeAttendanceCheckInOutRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	userId := request.UserId
	resValidate, resTA, resValidateMsg := services.VaildateCheckStatus(userId, "", checkStatus)
	if !resValidate {
		output := fiber.Map{
			"message": resValidateMsg,
			"user_id": userId,
			"data":    resTA,
		}
		return c.JSON(output)
	}

	request.CheckStatus = checkStatus
	request.RefId = services.StringWithCharset(userId, request.ProjectId)
	res := services.SaveData(request)

	output := fiber.Map{
		"message": os.Getenv("SAVE_CHECKIN_SUCCESS"),
		"user_id": request.UserId,
		"data":    res,
	}

	return c.JSON(output)
}

func CheckOutHandler(c *fiber.Ctx) error {
	var checkStatus = "checkout"
	var data = []model.TimeAttendanceReportList{}
	var request model.TimeAttendanceCheckInOutRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}
	userId := request.UserId
	refId := request.RefId

	if refId == "" {
		output := model.TimeAttendanceReportForMobileResponse{
			UserId:  userId,
			Data:    data,
			Message: "RefId ????",
		}
		return c.JSON(output)
	}

	resValidate, resTA, resValidateMsg := services.VaildateCheckStatus(userId, refId, checkStatus)
	if !resValidate {
		output := fiber.Map{
			"message": resValidateMsg,
			"user_id": request.UserId,
			"data":    resTA,
		}

		return c.JSON(output)
	}

	request.CheckStatus = checkStatus
	res := services.SaveData(request)

	output := fiber.Map{
		"message": os.Getenv("SAVE_CHECKOUT_SUCCESS"),
		"user_id": request.UserId,
		"data":    res,
	}

	return c.JSON(output)
}

func CheckStatusHandler(c *fiber.Ctx) error {

	userId := c.Query("user_id")
	if userId == "" {
		output := fiber.Map{
			"message": "UserId is Null",
			"user_id": userId,
			"data":    fiber.Map{},
		}
		return c.JSON(output)
	}

	output := fiber.Map{
		"message":        "CheckIn Screen",
		"user_id":        userId,
		"status_type_id": 10001,
		"data":           fiber.Map{},
	}
	data, err := services.CheckStatus(userId)
	if err != nil {
		output["message"] = os.Getenv("DB_CONNECT_ERROR")
		return c.JSON(output)
	}

	if data.CheckStatus != "checkout" {
		output["message"] = "CheckOut Screen"
		output["status_type_id"] = 10002
		output["data"] = data
	}

	return c.JSON(output)
}
