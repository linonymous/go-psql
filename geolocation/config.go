package geolocation

type Config struct {
	DBDriver string
	Enabled  bool
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
	SSLMode  string
}

func NewConfig() *Config {
	return &Config{
		DBDriver: "postgres",
		Enabled:  true,
		Host:     "localhost",
		Port:     "5432",
		DBName:   "linonymous",
		User:     "linonymous",
		Password: "5fE5AAns",
		SSLMode:  "enable",
	}
}
