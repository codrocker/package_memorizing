package shard

import "fmt"

// ModShard 对ID进行取模操作并返回补零的字符串
// id: 要进行取模的ID
// count: 取模的基数，必须是10的N次方（10, 100, 1000, ...）
// 返回: 取模后补零的字符串
func ModShard(id int64, count int) string {
	// 快速查表法：校验并获取宽度，常见的10的幂次方
	var width int
	switch count {
	case 10:
		width = 1
	case 100:
		width = 2
	case 1000:
		width = 3
	default:
		panic(fmt.Sprintf("count must be a power of 10 (10, 100, 1000, ...), got %d", count))
	}

	if id <= 0 {
		panic(fmt.Sprintf("id must be non-negative, got %d", id))
	}

	mod := id % int64(count)
	return fmt.Sprintf("%0*d", width, mod)
}
