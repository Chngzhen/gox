package dates

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// FormatDate2String 按照指定的格式将时间格式化为字符串。若datetime为空，则返回空字符串。
func FormatDate2String(datetime time.Time, format string) string {
	if 0 == len(format) {
		return ""
	}
	return doFormatDate2String(format, 0, 0, datetime)
}

// FormatNow2String 按照指定的格式将当前时间格式化为字符串。
func FormatNow2String(format string) string {
	if 0 == len(format) {
		return ""
	}
	return doFormatDate2String(format, 0, 0, time.Now())
}

// 利用起止游标遍历时间格式的片段，如yyyy、MM、dd、HH、mm、ss、SSS及连接符等，然后从时间对象中提取相应片段的数值字符串，并补位，使数值字符串的位数与片段相同。
// 例如，8时在格式化为HH时会提取成08。
func doFormatDate2String(str string, stIndex, enIndex int, dt time.Time) string {
	var result strings.Builder
	if stIndex > len(str) {
		// 首字节溢出，终止格式化
		return ""
	} else if enIndex >= len(str) {
		// 尾字节溢出，格式化时间片段
		fragment := str[stIndex:]
		if fragment == "" {
			return ""
		}
		result.WriteString(doFormatDateFragment2String(fragment, dt))
	} else if str[stIndex] != 'y' && str[stIndex] != 'M' && str[stIndex] != 'd' &&
		str[stIndex] != 'H' && str[stIndex] != 'm' && str[stIndex] != 's' && str[stIndex] != 'S' {
		// 首字节非时间字符，直接拼接，片段首字符索引继续向右浮动寻找新片段
		result.WriteByte(str[stIndex])
		result.WriteString(doFormatDate2String(str, stIndex+1, stIndex+1, dt))
	} else if enIndex == stIndex || str[enIndex] == str[stIndex] {
		// 首尾索引相等或首尾字符相同，锁定当前片段的首字符索引，尾字符索引继续向右游动
		result.WriteString(doFormatDate2String(str, stIndex, enIndex+1, dt))
	} else {
		// 首尾索引不等且首尾字符不同，意味着当前片段检索结束，格式化时间片段，并以尾索引为首索引向右浮动检索新片段。
		fragment := str[stIndex:enIndex]
		result.WriteString(doFormatDateFragment2String(fragment, dt))
		result.WriteString(doFormatDate2String(str, enIndex, enIndex, dt))
	}
	return result.String()
}

func doFormatDateFragment2String(fragment string, dateTime time.Time) string {
	switch fragment[0] {
	case 'y':
		yearString := strconv.Itoa(dateTime.Year())
		return doLpad(fragment, yearString, "0")
	case 'M':
		monthString := strconv.Itoa(int(dateTime.Month()))
		return doLpad(fragment, monthString, "0")
	case 'd':
		dateString := strconv.Itoa(dateTime.Day())
		return doLpad(fragment, dateString, "0")
	case 'H':
		hourString := strconv.Itoa(dateTime.Hour())
		return doLpad(fragment, hourString, "0")
	case 'm':
		minuteString := strconv.Itoa(dateTime.Minute())
		return doLpad(fragment, minuteString, "0")
	case 's':
		secondString := strconv.Itoa(dateTime.Second())
		return doLpad(fragment, secondString, "0")
	case 'S':
		microString := strconv.Itoa(dateTime.Nanosecond() / 1000000)
		return doLpad(fragment, microString, "0")
	default:
		return ""
	}
}

// 按照fragment的长度对source向左补str。
func doLpad(fragment string, source string, str string) string {
	sourceLength := len(source)
	fragmentLength := len(fragment)
	if sourceLength < fragmentLength {
		// 若source的长度小于fragment的，则补位
		// fmt支持空位填充。例如，%4s表示字符串长度为4，不足的在左边补空格；%04s表示字符串长度为4，不足的左边补0。
		fragmentFormat := "%" + str + strconv.Itoa(fragmentLength) + "s"
		return fmt.Sprintf(fragmentFormat, source)
	} else if sourceLength == fragmentLength {
		// 若source的长度等于fragment的，则不做任何处理，直接返回。
		return source
	} else {
		// 若source的长度大于fragment的，则从源串的右边开始截取。例如以yyy格式化2022，则截取022。
		return source[sourceLength-fragmentLength:]
	}
}
