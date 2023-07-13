package config

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

func NewConfig() *Config {
	return &Config{}
}

// GetConfig инициализирует и заполняет структуру конфигурационного файла
func GetConfig() (*Config, error) {
	var cfg Config

	configPath := "./"

	var v = viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(configPath)

	err := cfg.readParametersFromConfig(v)
	if err != nil {
		return &cfg, err
	}

	// Проверка наличия параметров в командной строке
	readFlags(&cfg)

	return &cfg, err
}

// readParametersFromConfig производит чтение и десериализацию конфигурационного файла
func (cfg *Config) readParametersFromConfig(v *viper.Viper) error {
	// Попытка чтения конфига
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("cannot read config: %v", err)
	}
	// Попытка заполнение структуры Config полученными данными
	if err := v.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("cannot read config: %v", err)
	}
	return nil
}

// readFlags реализует возможность передачи параметров
// конфигурационного файла при запуске из командной строки
func readFlags(cfg *Config) {
	flag.StringVar(&cfg.Path, "path", cfg.Path, "path to file for processing")
	flag.StringVar(&cfg.Name, "name", cfg.Name, "name of file for processing")
	flag.StringVar(&cfg.Object, "object", cfg.Object, "object for processing")

	flag.StringVar(&cfg.LogLevel, "logLevel", cfg.LogLevel, "The level of logging parameter")
	flag.BoolVar(&cfg.LogFileEnable, "logFileEnable", cfg.LogFileEnable, "The statement whether to log to a file")
	flag.BoolVar(&cfg.LogStdoutEnable, "logStdoutEnable", cfg.LogStdoutEnable, "The statement whether to log to console")
	flag.StringVar(&cfg.LogFile, "logpath", cfg.LogFile, "The path to file of logging out")
	flag.BoolVar(&cfg.RewriteLog, "rewriteLog", cfg.RewriteLog, "Is rewrite log file")

	flag.Parse()
}
