package config

type Config struct {
	Logger `yaml:"logger"`
	File   `yaml:"file"`
	// mode: add - добавить объект, remove - удалить
	Mode string `yaml:"mode"`
}

type File struct {
	// путь до файла
	Path string `yaml:"path"`
	// имя файла с указанием расширения (docx, doc, txt)
	Name string `yaml:"name"`
	// какую строку нужно вставить или удалить в начале абзаца
	Object string `yaml:"object"`
}

type Logger struct {
	LogLevel        string `mapstructure:"logLevel"`
	LogFileEnable   bool   `mapstructure:"logFileEnable"`
	LogStdoutEnable bool   `mapstructure:"logStdoutEnable"`
	LogFile         string `mapstructure:"logFile"`
	RewriteLog      bool   `mapstructure:"rewriteLog"`
}
