package conv

import (
	"bytes"
	"strconv"
	"strings"
)

func ParseI64s(s string, sep byte) ([]int64, error) {
	// split 比 readString 快
	ss := strings.Split(s, string(sep))
	i64s := make([]int64, 0, len(ss))
	for _, s := range ss {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		i64s = append(i64s, i)
	}
	return i64s, nil
}

func JoinI64s(i64s []int64, sep byte) string {
	buf := bytes.NewBuffer(nil)
	for i, i64 := range i64s {
		buf.WriteString(strconv.FormatInt(i64, 10))
		if i != len(i64s)-1 {
			buf.WriteByte(sep)
		}
	}
	return buf.String()
}
