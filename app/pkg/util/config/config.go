package config

type Config interface {
	get() Config
	getProperty(name string)
}

type GormConfig struct {
	Host string
	User string
	PassWord string
	DBName string
	Port string
	SSLMode string
	TimeZone string
}

func (c GormConfig) get() Config {
	return nil
}

func (c GormConfig) getProperty(name string) {
}

// type LogConfig struct {
// 	FileName string
// 	MaxSize int
// 	MaxAge int
// 	MaxBuckups int
// 	LocalTime bool
// 	Compress boll
// }