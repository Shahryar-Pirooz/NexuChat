package config

type Config struct {
	Database Postgres `mapstructure:"database"`
	Server   Server   `mapstructure:"server"`
	Nats     Nats     `mapstructure:"nats"`
}

type Postgres struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	DBName   string `mapstructure:"db-name"`
	Password string `mapstructure:"password"`
	SSLMode  string `mapstructure:"ssl-mode"`
}

type Server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Nats struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
