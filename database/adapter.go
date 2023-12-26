package database

import (
	"errors"

	"github.com/HelloLingC/moon-counter/common"
)

type IDatabase interface {
	CloseDB()
	InitDB()
	AddCounter(identifer string) (int, error)
	WriteCounter(identifer string, v int) error
	GetCounter(identifer string) (int, error)
	QueryCounter(skip int, limit int) ([]common.Counter, error)
}

func NewDBAdapter(dbType string, dbcfg *common.DBConfig) (IDatabase, error) {
	switch dbType {
	case "sqlite":
		return &SQLite{DBCfg: dbcfg}, nil
	default:
		return nil, errors.New("Unknow Database Type")
	}
}
