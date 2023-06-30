package godata

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "github.com/sijms/go-ora/v2"
	"log"
	"time"
)

type Conn struct {
	DB          *sql.DB
	tx          *sql.Tx
	Dialect     DialectType
	DSN         string
	maxLifetime time.Duration
	log         bool
}

func NewConnection(dialect DialectType, dsn string) (*Conn, error) {
	conn := &Conn{
		Dialect: dialect,
		DSN:     dsn,
	}

	err := conn.Open()

	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewConnectionOracle(dsn string) (*Conn, error) {
	return NewConnection(ORACLE, dsn)
}

func NewConnectionPostgres(dsn string) (*Conn, error) {
	return NewConnection(POSTGRESQL, dsn)
}

func NewConnectionFirebird(dsn string) (*Conn, error) {
	return NewConnection(FIREBIRD, dsn)
}

func NewConnectionInterbase(dsn string) (*Conn, error) {
	return NewConnection(INTERBASE, dsn)
}

func NewConnectionMySql(dsn string) (*Conn, error) {
	return NewConnection(MYSQL, dsn)
}

func NewConnectionSqLite(dsn string) (*Conn, error) {
	return NewConnection(SQLITE, dsn)
}

func NewConnectionSqlServer(dsn string) (*Conn, error) {
	return NewConnection(SQLSERVER, dsn)
}

func (co *Conn) EnableLog() {
	co.log = true
}

func (co *Conn) DisableLog() {
	co.log = false
}

func (co *Conn) Open() error {
	db, err := sql.Open(co.Dialect.String(), co.DSN)

	if err != nil {
		return fmt.Errorf("could not create a connection: %w", err)
	}

	if err = db.Ping(); err != nil {
		return fmt.Errorf("database is not reachable: %w", err)
	}

	co.DB = db

	return nil
}

func (co *Conn) SetConnMaxLifeTime(d time.Duration) {
	co.maxLifetime = d
	co.DB.SetConnMaxLifetime(d)
}

func (co *Conn) Ping() error {
	if err := co.DB.Ping(); err != nil {
		return fmt.Errorf("database is not reachable: %w", err)
	}
	return nil
}

func (co *Conn) CreateContext(ctx context.Context) (context.Context, context.CancelFunc) {
	timeout := 5 * time.Second
	return context.WithTimeout(ctx, timeout)
}

func (co *Conn) StartTransaction() error {
	if co.tx != nil {
		t, err := co.DB.Begin()
		if err != nil {
			return err
		}
		co.tx = t
	}
	return nil

}
func (co *Conn) Commit() error {
	err := co.tx.Commit()
	co.tx = nil
	return err
}
func (co *Conn) Rollback() error {
	err := co.tx.Rollback()
	co.tx = nil
	return err
}
func (co *Conn) Exec(sql string, arg ...any) (sql.Result, error) {
	return co.tx.Exec(sql, arg)
}

func (co *Conn) Close() {
	if err := co.DB.Close(); err != nil {
		log.Printf("could not close the database connection %v\n", err)
		return
	}
	log.Printf("database connection released succesfully")
}
