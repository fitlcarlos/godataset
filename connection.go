package godata

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "github.com/sijms/go-ora/v2"
	"strconv"
	"time"
)

type Conn struct {
	DB           *sql.DB
	Dialect      DialectType
	DSN          string
	log          bool
	PoolSize     int
	PoolLifetime time.Duration
	MaxOpenConns int
	ConnLifetime time.Duration
	connContext  bool
}

func NewConnection(dialect DialectType, dsn string) (*Conn, error) {
	conn := &Conn{
		Dialect:  dialect,
		DSN:      dsn,
		PoolSize: 20,
	}

	err := conn.Open()

	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewConnectionOracle(dsn string) (*Conn, error) {

	timout := strconv.FormatInt(int64(time.Minute*60), 10)

	return NewConnection(ORACLE, dsn+"?connection timeout="+timout+"&lob fetch=post")
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

	db.SetMaxIdleConns(co.PoolSize)
	db.SetMaxOpenConns(co.MaxOpenConns)
	db.SetConnMaxIdleTime(co.PoolLifetime)
	db.SetConnMaxLifetime(co.ConnLifetime)

	if err = db.Ping(); err != nil {
		return fmt.Errorf("database is not reachable: %w", err)
	}

	co.DB = db
	co.connContext = false

	return nil
}

// SetSizePool
// Tamanho maximo do Pool de conexão
func (co *Conn) SetSizePool(n int) {
	co.PoolSize = n
	co.DB.SetMaxIdleConns(n)
}

// SetPoolLifeTime
// Tempo de vida do Pool de conexões
func (co *Conn) SetPoolLifeTime(d time.Duration) {
	co.PoolLifetime = d
	co.DB.SetConnMaxIdleTime(d)
}

// SetMaxOpenConns
// Maximo de conexões abertas
func (co *Conn) SetMaxOpenConns(n int) {
	co.MaxOpenConns = n
	co.DB.SetMaxOpenConns(n)
}

// SetConnLifeTime
// Tempo de vida das conexões
func (co *Conn) SetConnLifeTime(d time.Duration) {
	co.ConnLifetime = d
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

func (co *Conn) StartTransaction() (*Transaction, error) {
	tx, err := NewTransaction(co)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (co *Conn) StartTransactionContext(ctx context.Context) (*Transaction, error) {
	tx, err := NewTransactionCtx(co, ctx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}
func (co *Conn) Exec(sql string, arg ...any) (sql.Result, error) {
	return co.DB.Exec(sql, arg)
}

func (co *Conn) Close() {
	if err := co.DB.Close(); err != nil {
		return
	}
}

func (co *Conn) NewDataSet() *DataSet {
	ds := &DataSet{
		Connection:   co,
		Index:        0,
		Recno:        0,
		Fields:       NewFields(),
		Params:       NewParams(),
		Macros:       NewMacros(),
		MasterSource: NewMasterSource(),
	}
	ds.Fields.Owner = ds

	return ds
}
