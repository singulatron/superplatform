package sqlstore

import (
	"database/sql"
	"fmt"

	"github.com/singulatron/singulatron/localtron/logger"
)

type DebugDB struct {
	*sql.DB
	Debug     bool
	tableName string
}

func NewDebugDB(db *sql.DB, tableName string) *DebugDB {
	return &DebugDB{
		DB:        db,
		tableName: tableName,
	}
}

func (db *DebugDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	db.logQuery(query, args...)
	return db.DB.Query(query, args...)
}

func (db *DebugDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	db.logQuery(query, args...)
	return db.DB.Exec(query, args...)
}

func (db *DebugDB) Prepare(query string) (*sql.Stmt, error) {
	db.logQuery(query, nil)
	return db.DB.Prepare(query)
}

func (db *DebugDB) logQuery(query string, args ...interface{}) {
	if db.Debug {
		logger.Info(fmt.Sprintf("[%v] Executing query: %s, args: %v", db.tableName, query, args))
	}
}
