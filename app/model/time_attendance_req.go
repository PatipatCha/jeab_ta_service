package model

type TimeAttendanceRequest struct {
	UserId         string         `json:"user_id"`
	TimeAttendance TimeAttendance `json:"time_attendance"`
}
