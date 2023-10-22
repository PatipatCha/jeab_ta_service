package configuration

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func AzureConfig() *Config {
	return &Config{
		Host:     "jeab-project-server.postgres.database.azure.com",
		Database: "postgres",
		User:     "noppakrit",
		Password: "Ys12345#",
	}
}

func AzureAccountDBConfig() *Config {
	return &Config{
		Host:     "jeab-project-server.postgres.database.azure.com",
		Database: "account_db",
		User:     "noppakrit",
		Password: "Ys12345#",
	}
}
