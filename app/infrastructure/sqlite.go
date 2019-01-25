package infrastructure

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Conn *sql.DB
}

func (handler *Sqlite) Execute(statement string) {
	handler.Conn.Exec(statement)
}

func (handler *Sqlite) Query(statement string) Row {
	//fmt.Println(statement)
	rows, err := handler.Conn.Query(statement)
	if err != nil {
		fmt.Println(err)
		return new(SqliteRow)
	}
	row := new(SqliteRow)
	row.Rows = rows
	return row
}

type SqliteRow struct {
	Rows *sql.Rows
}

func (r SqliteRow) Scan(dest ...interface{}) {
	r.Rows.Scan(dest...)
}

func (r SqliteRow) Next() bool {
	return r.Rows.Next()
}

func NewSqlite(dbfileName string) *Sqlite {
	conn, _ := sql.Open("sqlite3", dbfileName)
	sqlite := new(Sqlite)
	sqlite.Conn = conn
	return sqlite
}
