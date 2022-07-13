package config

import (
	"reflect"
)

type Config interface {
	GetProperty(key string) string
}

type DBConfig interface {
	Config
	GetOutput() string
}

type GormConfig struct {
	Host string
	User string
	PassWord string
	DBName string
	Port string
	SSLMode string
	TimeZone string
	Output string
}

func (c GormConfig) GetProperty(key string) string {
	r := reflect.ValueOf(c)
	v := r.FieldByName(key).Interface()
	return v.(string)
}

func (c GormConfig) GetOutput() string {
	return c.Output
}

// type LogConfig struct {
// 	FileName string
// 	MaxSize int
// 	MaxAge int
// 	MaxBuckups int
// 	LocalTime bool
// 	Compress boll
// }