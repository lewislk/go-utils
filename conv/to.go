package conv

import "strconv"

func Bool2Int8(v bool) int8 {
	if v {
		return int8(1)
	}
	return int8(0)
}

func Bool2Int(v bool) int {
	return int(Bool2Int8(v))
}

func Bool2Int32(v bool) int32 {
	return int32(Bool2Int8(v))
}

func Str2Int64E(v string) (int64, error) {
	return strconv.ParseInt(v, 10, 64)
}

func Str2Int64(v string, defaultValue int64) int64 {
	rv, err := Str2Int64E(v)
	if err != nil {
		return defaultValue
	}
	return rv
}

func Int642Str(v int64) string {
	return strconv.FormatInt(v, 10)
}
