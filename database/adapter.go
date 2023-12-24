package database

import (
	"errors"
)

type IDatabase interface {
	CloseDB()
	InitDB(dbname string)
	AddCounter(identifer string) (int, error)
	WriteCounter(identifer string, v int) error
	ReadCounter(identifer string) (int, error)
}

func NewDBAdapter(dbType string, dbname string) (IDatabase, error) {
	switch dbType {
	case "sqlite":
		return &SQLite{Dbname: dbname}, nil
	default:
		return nil, errors.New("Unknow Database Type")
	}
}
