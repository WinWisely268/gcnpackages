package dao

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/genjidb/genji/document"
	"github.com/getcouragenow/packages/sys-account/pkg/store"
)

type Account struct {
	ID                string
	Name              string
	Email             string
	Password          string
	RoleId            string
	UserDefinedFields map[string]interface{}
	CreatedAt         sql.NullTime
	UpdatedAt         sql.NullTime
	LastLogin         sql.NullTime
	Disabled          bool
}

func (a Account) TableName() string {
	return tableName("Accounts")
}

/*
,
			role_id TEXT NOT NULL,
			user_defined_fields JSON,
			created_at TIMESTAMP WITH TIME ZONE,
			updated_at TIMESTAMP WITH TIME ZONE,
			last_login TIMESTAMP WITH TIME ZONE,
			disabled BOOL
*/

func (a Account) CreateSQL() []string {
	return []string{
		`CREATE TABLE IF NOT EXISTS ` + a.TableName() + ` (
			id TEXT NOT NULL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT NOT NULL,
			password TEXT NOT NULL,
			role_id TEXT NOT NULL,
			user_defined_fields BLOB
		);`,
		"CREATE INDEX IF NOT EXISTS idx_accounts ON " + a.TableName() + "(name);",
	}
}

func (a *Account) getColumns() string {
	return `id, name, email, password, role_id, 
 				user_defined_fields, created_at, updated_at, last_login, disabled`
}

func (a *Account) getAccountSelectStatement(aqp *QueryParams) (string, []interface{}, error) {
	baseStmt := sq.Select(a.getColumns()).From(a.TableName())
	for k, v := range aqp.Params {
		baseStmt = baseStmt.Where(k, v)
	}
	return baseStmt.ToSql()
}

func (a *Account) GetAccount(db store.DbTxer, aqp *QueryParams) (*Account, error) {
	var acc Account
	selectStmt, args, err := a.getAccountSelectStatement(aqp)
	if err != nil {
		return nil, err
	}
	doc, err := db.QueryOne(selectStmt, args...)
	if err != nil {
		return nil, err
	}
	err = document.StructScan(doc, &acc)
	return &acc, err
}

func (a *Account) ListAccount(db store.DbTxer, aqp *QueryParams) ([]*Account, error) {
	var accs []*Account
	selectStmt, args, err := a.getAccountSelectStatement(aqp)
	if err != nil {
		return nil, err
	}
	res, err := db.Query(selectStmt, args...)
	if err != nil {
		return nil, err
	}
	err = res.Iterate(func(d document.Document) error {
		var acc *Account
		if err = document.StructScan(d, acc); err != nil {
			return err
		}
		accs = append(accs, acc)
		return nil
	})
	return accs, err
}

func (a *Account) Insert(db store.DbTxer, aqp *QueryParams) error {
	var allVals [][]interface{}
	columns, values := aqp.ColumnsAndValues()
	stmt, args, err := sq.Insert(a.TableName()).
		Columns(columns...).
		Values(values...).
		Suffix(`RETURNING ` + a.getColumns()).
		ToSql()
	if err != nil {
		return err
	}
	allVals = append(allVals, args)
	return db.Exec([]string{stmt}, allVals)
}

func (a *Account) Update(db store.DbTxer, aqp *QueryParams) error {
	var values [][]interface{}
	stmt, args, err := sq.Update(a.TableName()).SetMap(aqp.Params).ToSql()
	if err != nil {
		return err
	}
	values = append(values, args)
	return db.Exec([]string{stmt}, values)
}
