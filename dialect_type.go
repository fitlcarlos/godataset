package godata

import "fmt"

type DialectType uint8

const (
	FIREBIRD   DialectType = 1
	INTERBASE              = 2
	MYSQL                  = 3
	ORACLE                 = 4
	POSTGRESQL             = 5
	SQLSERVER              = 6
	SQLITE                 = 7
)

func (d DialectType) String() string {
	switch d {
	case FIREBIRD:
		return "firebird"
	case INTERBASE:
		return "interbase"
	case MYSQL:
		return "mysql"
	case ORACLE:
		return "oracle"
	case POSTGRESQL:
		return "pgx"
	case SQLSERVER:
		return "sqlserver"
	case SQLITE:
		return "sqlite"
	default:
		return fmt.Sprintf("%d", int(d))
	}
}
