package model

type TimeAttendance struct {
	CheckInDate   string `json:"check_in_date"`
	CheckInPlace  string `json:"check_in_place"`
	CheckInTime   string `json:"check_in_time"`
	CheckOutTime  string `json:"check_out_time"`
	CheckOutPlace string `json:"check_out_place"`
	WorkingHours  string `json:"working_hours"`
}
