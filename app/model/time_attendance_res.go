package model

type TimeAttendanceResponse struct {
	UserId         string         `json:"user_id"`
	TimeAttendance TimeAttendance `json:"time_attendance"`
	Message        string         `json:"message"`
}
