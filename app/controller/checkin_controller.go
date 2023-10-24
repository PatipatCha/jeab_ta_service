package controller

import (
	"os"

	"github.com/PatipatCha/jeab_ta_service/app/model"
	"github.com/PatipatCha/jeab_ta_service/app/services"
	"github.com/gofiber/fiber/v2"
)

func CheckIn(c *fiber.Ctx) error {
	var checkStatus = "checkin"
	// var res model.TimeAttendanceEntity
	var request model.TimeAttendanceCheckInRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	userId := request.UserId
	resValidate, _ := services.VaildateUserId(userId)
	if !resValidate {
		output := model.TimeAttendanceResponse{
			UserId:  userId,
			Data:    nil,
			Message: os.Getenv("VAILD_USERID_NOT_FOUND"),
		}
		return c.JSON(output)
	}

	resValidate, resTA, resValidateMsg := services.VaildateCheckStatus(userId, "", checkStatus)
	if !resValidate {

		output := fiber.Map{
			"message": resValidateMsg,
			"user_id": request.UserId,
			"data":    resTA,
		}

		return c.JSON(output)
	}

	request.CheckStatus = checkStatus
	// request.RefId = services.StringWithCharset(userId, request.ProjectId)
	request.RefId = services.GenUUID()
	res, err := services.SaveData(request)
	if err != nil {
		return err
	}

	output := fiber.Map{
		"message": os.Getenv("SAVE_CHECKIN_SUCCESS"),
		"user_id": request.UserId,
		"data":    res,
	}

	return c.JSON(output)
}

func CheckOut(c *fiber.Ctx) error {
	var checkStatus = "checkout"
	var request model.TimeAttendanceCheckInRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}
	userId := request.UserId
	refId := request.RefId

	if refId == "" {
		output := model.TimeAttendanceResponse{
			UserId:  userId,
			Data:    nil,
			Message: "RefId ????",
		}
		return c.JSON(output)
	}

	resValidate, _ := services.VaildateUserId(userId)
	if !resValidate {
		output := model.TimeAttendanceResponse{
			UserId:  userId,
			Data:    nil,
			Message: os.Getenv("VAILD_USERID_NOT_FOUND"),
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
	res, err := services.SaveData(request)
	if err != nil {
		return err
	}

	output := fiber.Map{
		"message": os.Getenv("SAVE_CHECKOUT_SUCCESS"),
		"user_id": request.UserId,
		"data":    res,
	}

	return c.JSON(output)
}