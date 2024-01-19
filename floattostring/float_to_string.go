package floattostring

import (
	"fmt"
	"strconv"
)

func Sprintf(f float64) string {
	return fmt.Sprintf("%f", f)
}

func Strconv(f float64, prec byte) string {
	return strconv.FormatFloat(f, prec, -1, 64)
}
