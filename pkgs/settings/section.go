package settings

type Config struct {
	Server         serverSetting
	ServiceSetting serviceSetting
}

type serviceSetting struct {
	PostgreSql postgreSetting `mapstructure:"database"`
}

type serverSetting struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	APIKey     string `mapstructure:"API_KEY"`
	SercetKey  string `mapstructure:"SERCET_KEY"`
}

type postgreSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username,omitempty"`
	Password        string `mapstructure:"password,omitempty"`
	DbName          string `mapstructure:"db_name"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
}
