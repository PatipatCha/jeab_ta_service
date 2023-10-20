package services

type GetReportService interface{}

// func GetReport(c *fiber.Ctx) error {
// 	var ta []model.TimeAttendanceEntity
// 	db, err := databases.ConnectDB()
// 	if err != nil {
// 		return err
// 	}

// 	// userId := c.Params("userId")
// 	res := db.Table("time_attendance").Find(&ta).Error
// 	if err := res; err != nil {
// 		return err
// 	}

// 	println()

// 	return c.JSON(ta)
// }
