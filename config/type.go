package config

type Config struct {
	Database Postgres `mapstructure:"database"`
	Server   Server   `mapstructure:"server"`
}

type Postgres struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	SSL      string `mapstructure:"ssl"`
}

type Server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
