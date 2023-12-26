package database

import (
	"database/sql"
	"log"

	"github.com/HelloLingC/moon-counter/common"
	_ "github.com/mattn/go-sqlite3"
)

const CREATE_TABLE = `CREATE TABLE IF NOT EXISTS counters (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	identifier TEXT NOT NULL UNIQUE,
	count INTEGER NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
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
	DB    *sql.DB
	DBCfg *common.DBConfig
}

func (s *SQLite) openDB() *sql.DB {
	db, err := sql.Open("sqlite3", s.DBCfg.Dbname)
	if err != nil {
		log.Fatal("Error connecting SQLite DB: ", err)
	}
	return db
}

func (s *SQLite) CloseDB() {
	s.DB.Close()
	stmt_add_counter.Close()
	stmt_insert_counter.Close()
	stmt_query_counter.Close()
}

func (s *SQLite) InitDB() {
	s.DB = s.openDB()
	_, err := s.DB.Exec(CREATE_TABLE)
	if err != nil {
		log.Fatal("Error creating Sqlite DB table", err)
	}
	stmt_query_counter, _ = s.DB.Prepare("SELECT count FROM counters WHERE identifier = ?")
	stmt_add_counter, _ = s.DB.Prepare(ADD_COUNTER)
	stmt_insert_counter, _ = s.DB.Prepare("INSERT OR REPLACE INTO counters (identifier, count) VALUES (?, ?)")
}

func (s *SQLite) AddCounter(identifer string) (int, error) {
	c, _ := s.GetCounter(identifer)
	if c == 0 {
		stmt_insert_counter.Exec(identifer, 1)
		return 1, nil
	}

	if _, err := stmt_add_counter.Exec(identifer); err != nil {
		return 0, err
	}
	return c, nil
}

func (s *SQLite) GetCounter(identifer string) (int, error) {
	row := stmt_query_counter.QueryRow(identifer)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *SQLite) WriteCounter(identifer string, v int) error {
	return nil
}

func (s *SQLite) QueryCounter(skip int, limit int) ([]common.Counter, error) {
	rows, err := s.DB.Query("SELECT * FROM counters LIMIT ? OFFSET ?", limit, skip)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var counters []common.Counter
	for rows.Next() {
		counter := &common.Counter{}
		if err := rows.Scan(&counter.Id, &counter.Identifier, &counter.Count, &counter.CreatedTime); err != nil {
			return nil, err
		}
		counters = append(counters, *counter)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return counters, nil
}

func (s *SQLite) Exec(st string) error {
	return nil
}
