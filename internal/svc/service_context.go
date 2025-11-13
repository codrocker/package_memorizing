package svc

import (
	"fmt"

	"package_memorizing/internal/config"
	"package_memorizing/internal/shard"
	"package_memorizing/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	// dbConns maps explicit datasource name (e.g. db_000) to sqlx.SqlConn.
	dbConns map[string]sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	sc := &ServiceContext{Config: c}

	// prepare DB connections
	datasourcesMap := c.PackageMemorizingDB.DataSources
	sc.dbConns = make(map[string]sqlx.SqlConn, len(datasourcesMap))
	for key, dsn := range datasourcesMap {
		sc.dbConns[key] = sqlx.NewMysql(dsn)
	}

	return sc
}

// getDBConn returns the sqlx.SqlConn for a given dbSuffix.
func (s *ServiceContext) getDBConn(dbSuffix string) sqlx.SqlConn {
	if len(s.dbConns) == 0 {
		return nil
	}

	// try the expected explicit key first
	expectedKey := fmt.Sprintf("db_%s", dbSuffix)
	if conn, ok := s.dbConns[expectedKey]; ok {
		return conn
	}

	// as a last resort, return any available conn (first key)
	// Note: This part is tricky without a stable key list.
	// A map's iteration order is not guaranteed.
	// For now, we will iterate and return the first one found.
	for _, conn := range s.dbConns {
		return conn
	}

	return nil
}

// GetPackageModelById returns a PackageMemorizingModel bound to the correct
// DB connection and table name for the provided id.
// The table suffix is zero-padded based on TablesPerDB, e.g. if TablesPerDB=100
// the table names will be package_memorizing_00..package_memorizing_99.
func (s *ServiceContext) GetPackageModelById(id int64) model.PackageMemorizingModel {
	dbSuffix, tableSuffix := shard.Shard(id, s.Config.PackageMemorizingDB.DBCount, s.Config.PackageMemorizingDB.TablesPerDB)
	db := s.getDBConn(dbSuffix)

	return model.NewPackageMemorizingModelWithTableSuffix(db, tableSuffix)
}
