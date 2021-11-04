package flow

import (
	"hash/crc32"
)

func CalcHash(input string) uint32 {
	//res := sha256.Sum256([]byte(input))
	return crc32.ChecksumIEEE([]byte(input))
}
