package services

import (
	"encoding/json"

	"github.com/PatipatCha/jeab_ta_service/app/configuration"
	"github.com/PatipatCha/jeab_ta_service/app/model"

	"github.com/gofiber/fiber/v2"
)

func GetUser(seoc_id string) string {
	var user model.UserApiReponse
	var userReq = "&seoc_id=" + seoc_id
	var baseUrl = configuration.ApiUserConfig()
	agent := fiber.Get(baseUrl + userReq)

	_, body, _ := agent.Bytes()

	json.Unmarshal(body, &user)

	return user.Data.SeocID
}
