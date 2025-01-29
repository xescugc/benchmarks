package inttostring

import (
	"fmt"
	"strconv"
)

const (
	iformat = "%d,%d,%d"
	sformat = "%s,%s,%s"
)

func StrconvItoaConcat(x, y, z int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y) + "," + strconv.Itoa(z)
}

func StrconvParseIntConcat(x, y, z int) string {
	return fmt.Sprintf(sformat, strconv.FormatInt(int64(x), 10), strconv.FormatInt(int64(y), 10), strconv.FormatInt(int64(z), 10))
}

func StrconvItoaSprintf(x, y, z int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y) + "," + strconv.Itoa(z)
}

func StrconvParseIntSprintf(x, y, z int) string {
	return fmt.Sprintf(sformat, strconv.FormatInt(int64(x), 10), strconv.FormatInt(int64(y), 10), strconv.FormatInt(int64(z), 10))
}

func Sprintf(x, y, z int) string {
	return fmt.Sprintf(iformat, x, y, z)
}
