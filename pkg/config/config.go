package config

type Config struct {
	Logger `yaml:"logger"`
	File   `yaml:"file"`
	Mode   string `yaml:"mode"`
}

type File struct {
	Path string `yaml:"path"`
	Name string `yaml:"name"`
}

type Logger struct {
	LogLevel        string `mapstructure:"logLevel"`
	LogFileEnable   bool   `mapstructure:"logFileEnable"`
	LogStdoutEnable bool   `mapstructure:"logStdoutEnable"`
	LogFile         string `mapstructure:"logFile"`
	RewriteLog      bool   `mapstructure:"rewriteLog"`
}
