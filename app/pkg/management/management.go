package management

import (
	"echosample/pkg/util/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
	"os"
	"gopkg.in/yaml.v3"
	dlog "log"
	"gorm.io/gorm/logger"
	"time"
	"gopkg.in/natefinch/lumberjack.v2"
)

// package global var
var dbmanager DBManager

func init(){
	dbOpen()
}

type Manager interface {
	GetConfig() config.Config
}

type DBManager interface {
	GetDSN() string
	GetDB() *gorm.DB
}

func GetDBManager() DBManager {
	return dbmanager
}

type GormManager struct {
	db *gorm.DB
	conf *config.GormConfig
}

func (m *GormManager) GetConfig() config.Config {
	if m.conf == nil {
		m.conf = &config.GormConfig{}
		b, _ := os.ReadFile("./config/gorm-config.yml")
		yaml.Unmarshal(b, m.conf)
	}
	return m.conf
}

func (m *GormManager) GetDSN() string {
	if m.conf == nil {
		m.GetConfig()
	}

	// やり方がダサい
	dsn := "host=" + m.conf.Host + " user=" + m.conf.User + " password=" + m.conf.PassWord + " dbname=" + m.conf.DBName + " port=" + m.conf.Port + " sslmode=" + m.conf.SSLMode + " TimeZone=" + m.conf.TimeZone

	return dsn
}

func (m *GormManager) GetDB() *gorm.DB {
	return m.db
}

func dbOpen() {
	dbm := &GormManager{}
	// gorm log writer
	w := &lumberjack.Logger {
		Filename: "/var/log/es/gorm.log", // output file
		MaxSize: 10, // MB
		MaxBackups: 10, // max quantity
		LocalTime: true, // default UTC
		Compress: true, // archive flag
	}
	db, err := gorm.Open(postgres.Open(dbm.GetDSN()), &gorm.Config{
		Logger: logger.New(
			dlog.New(w, "\r\n", dlog.LstdFlags),
			logger.Config{
				SlowThreshold: 1000 * time.Second,
				LogLevel: logger.Info,
				IgnoreRecordNotFoundError: false,
				Colorful: false,
			}),
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	dbm.db = db
	dbmanager = dbm
}

type LogManager struct {

}