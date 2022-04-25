package postgresql_utilits

import (
	"strconv"
	"strings"
)

func CustomRebind(startIndex int, query string) string {
	// Add space enough for 10 params before we have to allocate
	rqb := make([]byte, 0, len(query)+10)

	var i int
	j := startIndex - 1
	for i = strings.Index(query, "?"); i != -1; i = strings.Index(query, "?") {
		rqb = append(rqb, query[:i]...)

		rqb = append(rqb, '$')

		j++
		rqb = strconv.AppendInt(rqb, int64(j), 10)

		query = query[i+1:]
	}

	return string(append(rqb, query...))
}
