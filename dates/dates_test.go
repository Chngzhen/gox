package xdates

import (
	"fmt"
	"testing"
	"time"
)

func TestFormatDate2String(t *testing.T) {
	fmt.Println(FormatDate2String(time.Now(), FmtFullDateTime))
	fmt.Println(FormatNow2String(FmtFullDateTime))
}
