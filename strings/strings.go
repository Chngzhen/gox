package strings

// ArrayToCommaDelimitedString 将字符串数组arr拼接成英文逗号分隔的字符串。
func ArrayToCommaDelimitedString(arr []string) string {
	if arr == nil || len(arr) == 0 {
		return ""
	}
	var result string
	for _, i := range arr {
		result += "," + i
	}
	return result[1:]
}

// ArrayToDelimitedString 将字符串数组arr拼接成字符串delimiter分隔的字符串，每个元素使用前缀字符串prefix和后缀字符串suffix包围。
func ArrayToDelimitedString(arr []string, prefix, suffix, delimiter string) string {
	if arr == nil || len(arr) == 0 {
		return ""
	}
	var result []byte
	d := []byte(delimiter)
	p := []byte(prefix)
	s := []byte(suffix)
	for _, i := range arr {
		result = append(result, d...)
		result = append(result, p...)
		result = append(result, []byte(i)...)
		result = append(result, s...)
	}
	return string(result)[len(delimiter):]
}
