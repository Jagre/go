package sql

import (
	"fmt"
	"strings"
)

// SQLScript is sql script string
type SQLScript string

// Parse will replace "$n" with arguments
func (sql SQLScript) Parse(args ...interface{}) string {
	s := string(sql)
	if args == nil {
		return s
	}
	if len(args) == 0 {
		return s
	}
	for i, item := range args {
		s = strings.Replace(s, fmt.Sprintf("$%d", i+1), fmt.Sprintf("%v", item), -1)
	}
	return s
}
