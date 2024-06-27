/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 *
 * For commercial use, a separate license must be obtained by purchasing from The Authors.
 * For commercial licensing inquiries, please contact The Authors listed in the AUTHORS file.
 */
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

type DebugTx struct {
	*sql.Tx
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
	res, err := db.DB.Query(query, args...)
	db.logQuery(query, err, args...)
	return res, err
}

func (db *DebugDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	res, err := db.DB.Exec(query, args...)
	db.logQuery(query, err, args...)
	return res, err
}

func (db *DebugDB) Prepare(query string) (*sql.Stmt, error) {
	res, err := db.DB.Prepare(query)
	db.logQuery(query, err, nil)
	return res, err
}

func (db *DebugDB) logQuery(query string, err error, args ...interface{}) {
	if db.Debug {
		logger.Info(fmt.Sprintf("[%v] [ERROR: %v] Executing query: %s, args: %v", db.tableName, err, query, args))
	}
}

func (db *DebugDB) Begin() (*DebugTx, error) {
	tx, err := db.DB.Begin()
	if err != nil {
		return nil, err
	}
	return &DebugTx{
		Tx:        tx,
		Debug:     db.Debug,
		tableName: db.tableName,
	}, nil
}

func (db *DebugTx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	res, err := db.Tx.Query(query, args...)
	db.logQuery(query, err, args...)
	return res, err
}

func (db *DebugTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	res, err := db.Tx.Exec(query, args...)
	db.logQuery(query, err, args...)
	return res, err
}

func (db *DebugTx) Prepare(query string) (*sql.Stmt, error) {
	db.logQuery(query, nil)
	return db.Tx.Prepare(query)
}

func (db *DebugTx) logQuery(query string, err error, args ...interface{}) {
	if db.Debug {
		logger.Info(fmt.Sprintf("[%v] [ERROR: %v] [TRANSACTION] Executing query: %s, args: %v", db.tableName, err, query, args))
	}
}
