package config

func getDefaultConfig(config *Config) *Config {
	if config.HttpListenPort == "" {
		config.HttpListenPort = "8080"
	}
	if config.PostgresHost == "" {
		config.PostgresHost = "postgres"
	}
	if config.PostgresPort == "" {
		config.PostgresPort = "5432"
	}
	if config.PostgresDbUser == "" {
		config.PostgresDbUser = "product_services"
	}
	if config.PostgresDbPass == "" {
		config.PostgresDbPass = "product_services"
	}
	if config.PostgresDbName == "" {
		config.PostgresDbName = "product_services"
	}
	if config.SqLiteDbPath == "" {
		config.SqLiteDbPath = "database/product_services.db"
	}

	return config
}
