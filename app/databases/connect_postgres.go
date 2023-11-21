package databases

import (
	"fmt"

	"github.com/PatipatCha/jeab_ta_service/app/configuration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectTADB() (*gorm.DB, error) {
	// Initialize connection string.
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require", configuration.AzureTimeAttendanceDBConfig().Host, configuration.AzureTimeAttendanceDBConfig().User, configuration.AzureTimeAttendanceDBConfig().Password, configuration.AzureTimeAttendanceDBConfig().Database)

	// Initialize connection object using GORM.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully created connection to database")

	return db, nil
}
