package config

import (
	"github.com/zeromicro/go-zero/rest"
)

// Config is the application configuration including DB sharding settings.
type Config struct {
	rest.RestConf
	PackageMemorizingDB PackageMemorizingDBConfig `json:"PackageMemorizingDB" yaml:"PackageMemorizingDB"`
}

// PackageMemorizingDBConfig holds database connection and sharding settings.
// DataSources is now a map from explicit datasource name (e.g. "db_000") to DSN.
// This allows configuration to name each physical DB explicitly rather than relying on
// array index ordering.
type PackageMemorizingDBConfig struct {
	// DataSources maps explicit datasource names (like db_000, db_001) to DSNs.
	// If a key for a desired shard (e.g. db_000) is missing, the code will fall back
	// to a stable lexicographic selection strategy.
	DataSources map[string]string `json:"DataSources" yaml:"DataSources"`

	// DBCount is the number of database shards (physical DBs).
	DBCount int64 `json:"DBCount" yaml:"DBCount"`

	// TablesPerDB is the number of logical tables per database.
	TablesPerDB int64 `json:"TablesPerDB" yaml:"TablesPerDB"`
}
