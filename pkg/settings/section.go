package settings

type Config struct {
	// Server struct {
	// 	Port int `mapstructure:"port"`
	// } `mapstructure:"server"`
	PostgresConfig     PostgresConfig `mapstructure:"postgres"`
	LoggerConfig LoggerConfig   `mapstructure:"logger"`
	RedisConfig        RedisConfig    `mapstructure:"redis"`
}

type PostgresConfig struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	DBName          string `mapstructure:"dbname"`
	SSLMode         string `mapstructure:"sslmode"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

type LoggerConfig struct {
	LogLevel   string `mapstructure:"loglevel"`
	FilePath   string `mapstructure:"filePath"`
	MaxSize    int    `mapstructure:"maxSize"`    // Maximum size in megabytes before it is rotated
	MaxBackups int    `mapstructure:"maxBackups"` // Maximum number of old log files to retain
	MaxAge     int    `mapstructure:"maxAge"`     // Maximum number of days to retain old log files
	Compress   bool   `mapstructure:"compress"`   // Whether to compress rotated log files
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"poolSize"`
}
