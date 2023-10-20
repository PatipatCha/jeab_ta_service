package model

type TimeAttendanceResponse struct {
	UserId  string           `json:"user_id"`
	Data    []TimeAttendance `json:"data"`
	Message string           `json:"message"`
}
