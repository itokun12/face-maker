package orm

import (
	"fmt"
	"github.com/itokun12/image-maker/utils"
	"strings"
	"time"
)

type databaseConfig struct {
	DSN            string
	LifetimeMillis int
	MaxIdle        int
	MaxOpen        int
}

func NewDBFromEnvValue(schemaName string) (*DB, error) {
	config := databaseConfigFromEnv(schemaName)
	db, err := Open("mysql", config.DSN)
	if err != nil {
		return nil, err
	}
	db.impl.DB().SetConnMaxLifetime(time.Duration(config.LifetimeMillis) * time.Millisecond)
	db.impl.DB().SetMaxIdleConns(config.MaxIdle)
	db.impl.DB().SetMaxOpenConns(config.MaxOpen)
	// default: logging is enable
	if utils.Env("GORM_SQL_LOG", "1") == "1" {
		db.LogMode(true)
	} else {
		db.LogMode(false)
	}
	return db, nil
}

func databaseConfigFromEnv(schemaName string) *databaseConfig {
	upperSchemaName := strings.ToUpper(schemaName)
	prefix := "DATABASE_" + upperSchemaName
	user := utils.Env(prefix+"_USERNAME", utils.Env("DATABASE_USERNAME", "root"))
	password := utils.Env(prefix+"_PASSWORD", utils.Env("DATABASE_PASSWORD", "root"))
	host := utils.Env(prefix+"_HOST", utils.Env("DATABASE_HOST", "127.0.0.1"))
	port := utils.Env(prefix+"_PORT", utils.Env("DATABASE_PORT", "3306"))
	schema := utils.Env(prefix+"_SCHEMA", schemaName)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", user, password, host, port, schema)

	return &databaseConfig{
		DSN:            dsn,
		LifetimeMillis: utils.EnvInt(prefix+"_LIFETIME_MILLIS", 1000*10),
		MaxIdle:        utils.EnvInt(prefix+"_MAX_IDLE", 5),
		MaxOpen:        utils.EnvInt(prefix+"_MAX_OPEN", 10),
	}
}
