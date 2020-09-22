package pkg

import (
	"log"
	"strings"
	"github.com/ishangupta-ds/plot/pkg/types"
)

func (v *validator) writeOne(kind string, query string) {
	m := v.values[query]
	one := m[0]
	formattedQuery := one.Metric.String()
	qElements := strings.Split(formattedQuery, "{")
	labels := strings.Split(strings.Trim(qElements[1], "}"), ",")
	matcher := "{"
	for _, label := range labels {
		parts := strings.Split(label, "=")
		matcher += parts[0] + ":"
		matcher += strings.Trim(parts[1], "\"")
		matcher += ","
	}
	matcher += "}"
	value := one.Values[0].String()
	valElements := strings.Split(value, " ")
	output := kind + ", " + qElements[0] + ", " + matcher + ", [" + valElements[0] + "],"
	v.values[formattedQuery] = m[1:]
	log.Println(output)
	v.out.WriteString(output + "\n")
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
