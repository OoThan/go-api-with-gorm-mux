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
			Dialect:  "mysql",
			Username: "maungoothan",
			Password: "Mg12345!@#$%SQL",
			Name:     "employee",
			Charset:  "utf8",
		},
	}
}
