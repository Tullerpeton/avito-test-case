package configer

import (
	"log"

	"github.com/spf13/viper"
)

type ServerData struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type PostgresqlData struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db-name"`
	Port     string `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	Sslmode  string `mapstructure:"sslmode"`
}

type CurrencyConverter struct {
	AccessSecret string `mapstructure:"access-secret"`
	BaseCurrency string `mapstructure:"base-currency"`
	Format       string `mapstructure:"format"`
	HttpUrl      string `mapstructure:"http-url"`
}

type Config struct {
	Server     ServerData        `mapstructure:"server-data"`
	Postgresql PostgresqlData    `mapstructure:"postgresql-data"`
	Secret     CurrencyConverter `mapstructure:"currency-converter"`
}

var AppConfig Config

func Init(configPath string) {
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatal(err)
	}
}
