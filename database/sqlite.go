package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const CREATE_TABLE = `
CREATE TABLE IF NOT EXISTS counters (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	identifier TEXT NOT NULL UNIQUE,
	count INTEGER NOT NULL
);
`
const ADD_COUNTER = `UPDATE counters
SET count = count + 1
WHERE identifier = ?;
`

var stmt_add_counter *sql.Stmt
var stmt_insert_counter *sql.Stmt
var stmt_query_counter *sql.Stmt

type SQLite struct {
	DB     *sql.DB
	Dbname string
}

func (s *SQLite) openDB() *sql.DB {
	db, err := sql.Open("sqlite3", s.Dbname)
	if err != nil {
		log.Fatal("Error connecting SQLite DB: ", err)
	}
	return db
}

func (s *SQLite) CloseDB() {
	s.DB.Close()
}

func (s *SQLite) InitDB(dbname string) {
	s.Dbname = dbname
	s.DB = s.openDB()
	stmt_query_counter, _ = s.DB.Prepare("SELECT count FROM counters WHERE identifier = ?")
	stmt_add_counter, _ = s.DB.Prepare(ADD_COUNTER)
	stmt_insert_counter, _ = s.DB.Prepare("INSERT OR REPLACE INTO counters (identifier, count) VALUES (?, ?)")

	_, err := s.DB.Exec(CREATE_TABLE)
	if err != nil {
		log.Fatal("Error creating Sqlite DB table", err)
	}
}

func (s *SQLite) AddCounter(identifer string) (int, error) {
	c, _ := s.ReadCounter(identifer)
	if c == 0 {
		stmt_insert_counter.Exec(identifer, 1)
		return 1, nil
	}

	if _, err := stmt_add_counter.Exec(identifer); err != nil {
		return 0, err
	}
	return c, nil
}

func (s *SQLite) WriteCounter(identifer string, v int) error {
	return nil
}

func (s *SQLite) ReadCounter(identifer string) (int, error) {
	row := stmt_query_counter.QueryRow(identifer)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
