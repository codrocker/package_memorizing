package shard

import (
	"fmt"
	"strconv"
)

const (
	DB10 int64 = 10
)
const (
	TABLE10   int64 = 10
	TABLE100  int64 = 100
	TABLE1000 int64 = 1000
)

// Shard returns the database and table suffixes for a given ID.
// For example, with id=4034, dbCount=100, tableCount=1000, it returns "04", "034".
func Shard(id int64, dbCount, tableCount int64) (string, string) {
	switch dbCount {
	case DB10:
	default:
		panic(fmt.Sprintf("invalid dbCount: %d", dbCount))
	}
	switch tableCount {
	case TABLE10, TABLE100, TABLE1000:
	default:
		panic(fmt.Sprintf("invalid tableCount: %d", tableCount))
	}

	dbIndex := (id / int64(tableCount)) % int64(dbCount)
	tableIndex := id % int64(tableCount)

	dbLen := len(strconv.Itoa(int(dbCount - 1)))
	tableLen := len(strconv.Itoa(int(tableCount - 1)))

	return fmt.Sprintf("%0*d", dbLen, dbIndex), fmt.Sprintf("%0*d", tableLen, tableIndex)
}
