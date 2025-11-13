package model

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PackageMemorizingModel = (*customPackageMemorizingModel)(nil)

type (
	// PackageMemorizingModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPackageMemorizingModel.
	PackageMemorizingModel interface {
		packageMemorizingModel
		withSession(session sqlx.Session) PackageMemorizingModel
	}

	customPackageMemorizingModel struct {
		*defaultPackageMemorizingModel
	}
)

// NewPackageMemorizingModel returns a model for the database table.
func NewPackageMemorizingModel(conn sqlx.SqlConn) PackageMemorizingModel {
	return &customPackageMemorizingModel{
		defaultPackageMemorizingModel: newPackageMemorizingModel(conn),
	}
}

// NewPackageMemorizingModelWithTable returns a model bound to a specific table name.
// table should be the logical table name without backticks, e.g. "package_memorizing_0".
func NewPackageMemorizingModelWithTable(conn sqlx.SqlConn, table string) PackageMemorizingModel {
	m := newPackageMemorizingModel(conn)
	if table == "" {
		m.table = "`package_memorizing`"
	} else {
		m.table = fmt.Sprintf("`%s`", table)
	}
	return &customPackageMemorizingModel{defaultPackageMemorizingModel: m}
}

// NewPackageMemorizingModelWithTableSuffix returns a model bound to a specific table name constructed from a suffix.
func NewPackageMemorizingModelWithTableSuffix(conn sqlx.SqlConn, tableSuffix string) PackageMemorizingModel {
	tableName := fmt.Sprintf("package_memorizing_%s", tableSuffix)
	return NewPackageMemorizingModelWithTable(conn, tableName)
}

func (m *customPackageMemorizingModel) withSession(session sqlx.Session) PackageMemorizingModel {
	return NewPackageMemorizingModel(sqlx.NewSqlConnFromSession(session))
}
