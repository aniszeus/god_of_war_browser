package colexpu

import (
	"bytes"
	"fmt"
	"strings"
)

func Floats64ToString(a []float64) string {
	var b bytes.Buffer
	for _, f := range a {
		b.WriteString(fmt.Sprintf("%f ", f))
	}
	return strings.TrimRight(b.String(), " ")
}

func Floats32ToString(a []float32) string {
	var b bytes.Buffer
	for _, f := range a {
		b.WriteString(fmt.Sprintf("%f ", f))
	}
	return strings.TrimRight(b.String(), " ")
}

func IntsToString(a []int) string {
	var b bytes.Buffer
	for _, i := range a {
		b.WriteString(fmt.Sprintf("%d ", i))
	}
	return strings.TrimRight(b.String(), " ")
}

func CreateAccessor(count, offset int, source string, stride int, paramsAndTypes ...string) string {
	a := fmt.Sprintf("<accessor count=\"%d\" offset=\"%d\" source=\"%s\" stride=\"%d\">\n",
		count, offset, source, stride)
	for i := 0; i < len(paramsAndTypes); i += 2 {
		a += fmt.Sprintf("<param name=\"%s\" type=\"%s\"/>\n", paramsAndTypes[i], paramsAndTypes[i+1])
	}
	return a + "</accessor>\n"
}
