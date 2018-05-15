package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgresql",
			Username: "test",
			Password: "test0000!",
			Name:     "todoapp",
			Charset:  "utf8",
		},
	}
}
