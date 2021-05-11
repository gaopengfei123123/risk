package utils

import (
	"bytes"
	"encoding/gob"
	"strconv"
)

// 任意类型转字节
func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// 移出字符串中的非制表符
func PureString(raw string) (res string, err error) {
	ttA := []rune(raw)
	tt := make([]rune, 0)
	for i := 0; i < len(ttA); i++ {
		if strconv.IsGraphic(ttA[i]) {
			tt = append(tt, ttA[i])
		}
	}
	res = string(tt)
	return
}
