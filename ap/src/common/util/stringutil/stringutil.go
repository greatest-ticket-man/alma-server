package stringutil

import (
	"fmt"
	"strconv"
	"strings"
)

// Uint32ToString .
func Uint32ToString(num uint32) string {
	return strconv.FormatUint(uint64(num), 10)
}

// SplitLen 指定文字列で文字列を分割して返す
func SplitLen(src string, splitLen int32) []string {

	result := []string{"", ""}
	srcLen := int32(len(src))
	for i := int32(0); i < srcLen; i += splitLen {
		if i+splitLen < srcLen {
			result[0] = src[i:(i + splitLen)]
		} else {
			result[1] = src[i:]
		}
	}

	return result
}

// AddComma 数字をStringに変換して、3桁ずつにカンマをつける
func AddComma(num int32) string {
	arr := strings.Split(fmt.Sprintf("%d", num), "")
	cnt := len(arr) - 1
	res := ""
	i2 := 0

	for i := cnt; i >= 0; i-- {

		if i2 > 2 && i2%3 == 0 {
			res = fmt.Sprintf(",%s", res)
		}
		res = fmt.Sprintf("%s%s", arr[i], res)
		i2++
	}

	return res
}
