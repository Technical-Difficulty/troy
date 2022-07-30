package util

import (
	"encoding/hex"
	"strconv"
)

func BytesToUInt64(bytes []byte) (uint64, error) {
	return strconv.ParseUint(hex.EncodeToString(bytes), 16, 64)
}
