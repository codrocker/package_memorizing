package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

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

func (m *customPackageMemorizingModel) withSession(session sqlx.Session) PackageMemorizingModel {
	return NewPackageMemorizingModel(sqlx.NewSqlConnFromSession(session))
}
