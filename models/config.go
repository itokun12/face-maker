package models

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/itokun12/image-maker/orm"
	"github.com/pkg/errors"
)

var (
	masterDB *orm.DB
	logger   = logrus.New()
)

type ConnectionType int

const (
	_ ConnectionType = iota
	MySQL
)

func Initialize() {
	masterDB = connectToDB(MySQL, "image-maker")
}

func Finalize() {
	masterDB.Close()
}

func connectToDB(typ ConnectionType, schemaName string) *orm.DB {
	var db *orm.DB
	var err error
	if typ == MySQL {
		db, err = orm.NewDBFromEnvValue(schemaName)
	} else {
		e := fmt.Errorf("Unknown connection type %v", typ)
		logger.Errorln(e)
		panic(e)
	}

	if err != nil {
		e := errors.Wrapf(err, "Connect to database failed (schemaName:%v)", schemaName)
		logger.Errorln(e)
		panic(e)
	}
	return db
}
