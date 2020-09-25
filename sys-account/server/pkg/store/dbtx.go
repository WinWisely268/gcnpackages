package store

import (
	"github.com/genjidb/genji/document"
	"github.com/genjidb/genji/sql/query"
)

// DbTxer provides common interface for executing sql statements
type DbTxer interface {
	Exec([]string, [][]interface{}) error
	Query(string, ...interface{}) (query.Result, error)
	QueryOne(string, ...interface{}) (document.Document, error)
	BuildSearchQuery(string) string
}
