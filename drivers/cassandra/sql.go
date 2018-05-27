package cassandra

import (
	"context"
	"database/sql"
	"database/sql/driver"
)

// gocql does not follow database/sql interface, this file contains wrappers for it

func init() {
	sql.Register("cassandra", Driver{})
}

type sqlDriver interface {
	driver.Driver
}

type sqlConn interface {
	driver.Conn
	driver.ConnBeginTx
	driver.ConnPrepareContext
}

type sqlStmt interface {
	driver.Stmt
	driver.StmtExecContext
	driver.StmtQueryContext
}

type sqlTx interface {
	driver.Tx
	// TODO: any extra interface?
}

var _ sqlDriver = Driver{}
var _ sqlConn = (*Conn)(nil)
var _ sqlStmt = (*Stmt)(nil)
var _ sqlTx = (*Tx)(nil)

type Driver struct{}

type Conn struct {
}

type Stmt struct {
}

type Tx struct {
}

// TODO: might change to pointer receiver...
func (d Driver) Open(dsn string) (driver.Conn, error) {
	return &Conn{}, nil
}

func (c *Conn) Prepare(query string) (driver.Stmt, error) {
	return c.PrepareContext(context.Background(), query)
}

func (c *Conn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	return &Stmt{}, nil
}

func (c *Conn) Close() error {
	return nil
}

func (c *Conn) Begin() (driver.Tx, error) {
	return c.BeginTx(context.Background(), driver.TxOptions{})
}

func (c *Conn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	return &Tx{}, nil
}

func (s *Stmt) Close() error {
	return nil
}

func (s *Stmt) NumInput() int {
	return -1
}

func (s *Stmt) Exec(args []driver.Value) (driver.Result, error) {
	// TODO: how to convert to ExecContext, need to change it to named value ... use slice index and ordinal?
	return nil, nil
}

func (s *Stmt) Query(args []driver.Value) (driver.Rows, error) {
	return nil, nil
}

func (s *Stmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	return nil, nil
}

func (s *Stmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	return nil, nil
}

func (t *Tx) Commit() error {
	return nil
}

func (t *Tx) Rollback() error {
	return nil
}
