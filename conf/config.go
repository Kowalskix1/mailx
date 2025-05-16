package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	Mode    string `mapstructure:"mode"`
}

type LogConfig struct {
	LogLevel           string `mapstructure:"logLevel"`
	logFile            string `mapstructure:"logFile"`
	logFileMaxSize     int    `mapstructure:"logFileMaxSize"`
	logFileMaxNums     int    `mapstructure:"logFileMaxNums"`
	logFileMaxKeepTime int    `mapstructure:"loglogFileMaxKeepTimeLevel"`
	logFileCompress    bool   `mapstructure:"logFileCompress"`
}

type PostgresConfig struct {
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	User   string `mapstructure:"user"`
	Passwd string `mapstructure:"passwd"`
}

type RedisConfig struct {
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	Passwd string `mapstructure:"passwd"`
}

type SMTPConfig struct {
	Addr              string `mapstructure:"addr"`
	TLSCert           string `mapstructure:"tlsCert"`
	TLSKey            string `mapstructure:"tlsKey"`
	RequireTLS        bool   `mapstructure:"requireTls"`
	RequireStartTls   bool   `mapstructure:"requireStartTls"`
	AuthFile          string `mapstructure:"authFile"`
	AuthAllowInsecure bool   `mapstructure:"authAllowInsecure"`
	AuthAcceptAny     bool   `mapstructure:"authAcceptAny"`
	MaxRecipients     int    `mapstructure:"maxRecipients"`
	// reloay config
	// pop3 config
}

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Log      LogConfig      `mapstructure:"logger"`
	Postgres PostgresConfig `mapstructure:"postgresql"`
	Redis    RedisConfig    `mapstructure:"redis"`
	SMTPD    SMTPConfig     `mapstructure:"smtpd"`
}

var Conf Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("Load configuration file failed, error: %v", err)
	}
	viper.Unmarshal(&Conf)

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		logrus.Infoln("Config file has been changed.", in.Name)
		viper.ReadInConfig()
		viper.Unmarshal(&Conf)
	})

	// verify configurations
	if Conf.SMTPD.TLSCert != "" && Conf.SMTPD.TLSKey == "" || Conf.SMTPD.TLSCert == "" && Conf.SMTPD.TLSKey != "" {
		logrus.Fatalln("[SMTPD] You must provide both an SMTP tlsCert and tlsKey")
	}

	if Conf.SMTPD.RequireTLS || Conf.SMTPD.RequireStartTls {
		
	}
}
