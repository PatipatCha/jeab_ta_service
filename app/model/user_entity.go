package model

type UserApiReponse struct {
	Data    Data   `json:"data"`
	Message string `json:"message"`
	UserID  string `json:"user_id"`
}

type Data struct {
	SeocID string `json:"seoc_id"`
}
