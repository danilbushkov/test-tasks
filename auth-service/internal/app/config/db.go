package config

type ConnConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
}

type DBConfig struct {
	Conn *ConnConfig
}
