package config

type Config struct {
	DbName     string
	DbHost     string
	DbUser     string
	DbPassword string
}

func GetConfig() Config {
	return Config{
		DbHost:     "localhost",
		DbName:     "cleverHouse-db",
		DbUser:     "postgres",
		DbPassword: "postgres",
	}
}
