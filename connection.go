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
	DB     *sql.DB
	tx     *sql.Tx
	Dialect DialectType
	DSN     string
}
func NewConnection(dialect DialectType, dsn string) (*Conn, error) {
	conn := &Conn{
		Dialect: dialect,
		DSN: dsn,
	}

	db, err := sql.Open(dialect.String(), dsn)

	if err != nil {
		return nil, fmt.Errorf("could not create a connection: %w", err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("database is not reachable: %w", err)
	}

	conn.DB = db
	return conn, nil
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
