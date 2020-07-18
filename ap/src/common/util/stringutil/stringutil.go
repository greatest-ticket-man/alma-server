package stringutil

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
