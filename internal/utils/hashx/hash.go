package hashx

import "github.com/cespare/xxhash/v2"

func Sum64String(s string) uint64 {
	return xxhash.Sum64String(s)
}

func Sum64(b []byte) uint64 {
	return xxhash.Sum64(b)
}
