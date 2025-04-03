package settings

type Config struct {
	Server         serverSetting
	ServiceSetting serviceSetting
}

type serviceSetting struct {
	PostgreSql   postgreSetting `mapstructure:"database"`
	KafkaSetting kafkaSetting   `mapstructure:"kafka"`
}

type serverSetting struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
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

type kafkaSetting struct {
	KafkaBroker_1 string `mapstructure:"kafka_broker_1"`
	KafkaBroker_2 string `mapstructure:"kafka_broker_2"`
	KafkaBroker_3 string `mapstructure:"kafka_broker_3"`
}
