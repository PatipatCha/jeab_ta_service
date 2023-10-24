// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    timeAttendanceReportList, err := UnmarshalTimeAttendanceReportList(bytes)
//    bytes, err = timeAttendanceReportList.Marshal()

package model

import "encoding/json"

func UnmarshalTimeAttendanceReportList(data []byte) (TimeAttendanceReportList, error) {
	var r TimeAttendanceReportList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TimeAttendanceReportList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TimeAttendanceReportList struct {
	UserID  string  `json:"user_id"`
	Data    []Datum `json:"data"`
	Message string  `json:"message"`
}

type Datum struct {
	Date  Date   `json:"date"`
	Lists []List `json:"lists"`
}

type Date struct {
	Day   string `json:"day"`
	Month string `json:"month"`
}

type List struct {
	ProjectPlace   string `json:"project_place"`
	CheckInTime    string `json:"check_in_time"`
	CheckOutTime   string `json:"check_out_time"`
	CheckOutRemark string `json:"check_out_remark"`
	Total          string `json:"total"`
}
