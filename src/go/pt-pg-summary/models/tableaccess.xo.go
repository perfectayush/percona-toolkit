// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
)

// Table Access
type TableAccess struct {
	Relname string         // relname
	Relkind string         // relkind
	Datname sql.NullString // datname
	Count   sql.NullInt64  // count
}

// GetTableAccesses runs a custom query, returning results as TableAccess.
func GetTableAccesses(db XODB) ([]*TableAccess, error) {
	var err error

	// sql query
	sqlstr := `SELECT c.relname, c.relkind, b.datname datname, count(*) FROM pg_locks a ` +
		`JOIN pg_stat_database b ` +
		`ON a.database=b.datid ` +
		`JOIN pg_class c ` +
		`ON a.relation=c.oid ` +
		`WHERE a.relation IS NOT NULL ` +
		`AND a.database IS NOT NULL ` +
		`GROUP BY 1,2,3`

	// run query
	XOLog(sqlstr)
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*TableAccess{}
	for q.Next() {
		ta := TableAccess{}

		// scan
		err = q.Scan(&ta.Relname, &ta.Relkind, &ta.Datname, &ta.Count)
		if err != nil {
			return nil, err
		}

		res = append(res, &ta)
	}

	return res, nil
}