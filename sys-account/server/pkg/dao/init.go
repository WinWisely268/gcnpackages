package dao

import (
	"errors"
	"github.com/genjidb/genji"
	"github.com/genjidb/genji/document"
	"github.com/genjidb/genji/sql/query"

	"github.com/getcouragenow/packages/sys-core/server/pkg/db"
	"strings"
)

var (
	tablePrefix = ""
	modName     = "sys_accounts"
)

// QueryParams can be any condition
type QueryParams struct {
	Params map[string]interface{}
}

func (qp *QueryParams) ColumnsAndValues() ([]string, []interface{}) {
	var columns []string
	var values []interface{}
	for k, v := range qp.Params {
		columns = append(columns, k)
		values = append(values, v)
	}
	return columns, values
}

func tableName(name string) string {
	return tablePrefix + "_" + modName + "_" + name
}

// AccountDB struct satisfies store.DbTxer interface
type AccountDB struct {
	*genji.DB
}

func NewAccountDB(d *genji.DB) (*AccountDB, error) {
	tables := []db.DbModel{
		Account{},
		Project{},
		// Org{},
		// Roles{},
		// Permission{},
	}
	db.RegisterModels(modName, tables)
	if err := db.MakeSchema(d); err != nil {
		return nil, err
	}
	return &AccountDB{d}, nil
}

func (a *AccountDB) Exec(stmts []string, argSlices [][]interface{}) error {
	if len(stmts) != len(argSlices) {
		return errors.New("mismatch statements and argument counts")
	}
	tx, err := a.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	for i, stmt := range stmts {
		if err := tx.Exec(stmt, argSlices[i]...); err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (a *AccountDB) Query(stmt string, args ...interface{}) (*query.Result, error) {
	return a.Query(stmt, args...)
}

func (a *AccountDB) QueryOne(stmt string, args ...interface{}) (document.Document, error) {
	return a.QueryDocument(stmt, args...)
}

func (a *AccountDB) BuildSearchQuery(qs string) string {
	var sb strings.Builder
	sb.WriteString("%")
	sb.WriteString(qs)
	sb.WriteString("%")
	return sb.String()
}
