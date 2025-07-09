package godataset

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/sijms/go-ora/v2"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Conn struct {
	DB           *sql.DB
	Dialect      DialectType
	DSN          string
	Schema       string
	log          bool
	PoolSize     int           // Máximo de conexões abertas
	PoolLifetime time.Duration // Tempo máximo ocioso antes de fechar a conexão
	MaxOpenConns int           // Máximo de conexões ociosas
	ConnLifetime time.Duration // Tempo máximo de vida de uma conexão
	connContext  bool
}

func NewConnection(dialect DialectType, dsn string) (*Conn, error) {
	switch dialect {
	case ORACLE:
		return NewConnectionOracle(dsn)
	case POSTGRESQL:
		return NewConnectionPostgres(dsn)
	case FIREBIRD:
		return NewConnectionFirebird(dsn)
	case INTERBASE:
		return NewConnectionInterbase(dsn)
	case MYSQL:
		return NewConnectionMySql(dsn)
	case SQLITE:
		return NewConnectionSqLite(dsn)
	case SQLSERVER:
		return NewConnectionSqlServer(dsn)
	default:
		return nil, nil
	}
}

func newConnection(dialect DialectType, dsn string) (*Conn, error) {
	conn := &Conn{
		Dialect:      dialect,
		DSN:          dsn,
		PoolSize:     20,
		PoolLifetime: time.Second * 20,
	}

	err := conn.Open()

	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewConnectionOracle(dsn string) (*Conn, error) {

	timout := strconv.FormatInt(int64(time.Minute*60), 10)
	return newConnection(ORACLE, dsn+"?connection timeout="+timout+"&lob fetch=post")
}

func NewConnectionPostgres(dsn string) (*Conn, error) {
	if !strings.Contains(dsn, "application_name") {
		if strings.Contains(dsn, "?") {
			dsn = dsn + "&application_name=" + GetApplicationName()
		} else {
			dsn = dsn + "?application_name=" + GetApplicationName()
		}
	}

	return newConnection(POSTGRESQL, dsn)
}

func NewConnectionFirebird(dsn string) (*Conn, error) {
	return newConnection(FIREBIRD, dsn)
}

func NewConnectionInterbase(dsn string) (*Conn, error) {
	return newConnection(INTERBASE, dsn)
}

func NewConnectionMySql(dsn string) (*Conn, error) {
	return newConnection(MYSQL, dsn)
}

func NewConnectionSqLite(dsn string) (*Conn, error) {
	return newConnection(SQLITE, dsn)
}

func NewConnectionSqlServer(dsn string) (*Conn, error) {
	return newConnection(SQLSERVER, dsn)
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

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("could not connect: %w", err)
	}

	db.SetMaxIdleConns(co.PoolSize)
	db.SetMaxOpenConns(co.MaxOpenConns)
	db.SetConnMaxIdleTime(co.PoolLifetime)
	db.SetConnMaxLifetime(co.ConnLifetime)

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

func (co *Conn) SetSchema(s string) {
	co.Schema = s
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
	co.DB = nil
}

func (co *Conn) NewDataSet() *DataSet {
	return NewDataSet(co)
}

func (co *Conn) AddOracleSessionParam(key, value string) error {
	_, err := co.DB.Exec(fmt.Sprintf("alter session set %s=%s", key, value))
	if err != nil {
		return err
	}

	return nil
}

func GetApplicationName() string {
	execPath, _ := os.Executable()
	_, execName := filepath.Split(execPath)
	return execName
}
