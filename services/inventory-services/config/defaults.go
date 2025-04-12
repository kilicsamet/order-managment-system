package config

func getDefaultConfig(config *Config) *Config {
	if config.HttpListenPort == "" {
		config.HttpListenPort = "8081"
	}
	if config.PostgresHost == "" {
		config.PostgresHost = "postgres"
	}
	if config.PostgresPort == "" {
		config.PostgresPort = "5432"
	}
	if config.PostgresDbUser == "" {
		config.PostgresDbUser = "inventory_services"
	}
	if config.PostgresDbPass == "" {
		config.PostgresDbPass = "inventory_services"
	}
	if config.PostgresDbName == "" {
		config.PostgresDbName = "inventory_services"
	}
	if config.SqLiteDbPath == "" {
		config.SqLiteDbPath = "database/inventory_services.db"
	}

	return config
}
