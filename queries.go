package coozy

import (
	"strings"

	"github.com/gobuffalo/pop/nulls"
	"github.com/gofrs/uuid"
)

// AppendQueryUUIDCriteria adds an AND clause to the query being constructed in with the provided Builder and parameter values slice.
func AppendQueryUUIDCriteria(queryBuilder *strings.Builder, queryParameters *[]interface{}, fieldName string, fieldValue uuid.UUID) (added bool) {
	if fieldValue != uuid.Nil {
		if queryBuilder.Len() > 0 {
			queryBuilder.WriteString(" AND ")
		}
		queryBuilder.WriteString(fieldName)
		queryBuilder.WriteString(" = ?")
		*queryParameters = append(*queryParameters, fieldValue)
		added = true
	}

	return
}

// AppendQueryStringCriteria adds an AND clause to the query being constructed in with the provided Builder and parameter values slice.
func AppendQueryStringCriteria(queryBuilder *strings.Builder, queryParameters *[]interface{}, fieldName string, fieldValue string) (added bool) {
	if fieldValue != "" {
		if queryBuilder.Len() > 0 {
			queryBuilder.WriteString(" AND ")
		}
		queryBuilder.WriteString(fieldName)
		queryBuilder.WriteString(" = ?")
		*queryParameters = append(*queryParameters, fieldValue)
		added = true
	}

	return
}

// AppendQueryNullableStringCriteria adds an AND clause to the query being constructed in with the provided Builder and parameter values slice.
func AppendQueryNullableStringCriteria(queryBuilder *strings.Builder, queryParameters *[]interface{}, fieldName string, fieldValue nulls.String) (added bool) {
	if fieldValue.Valid {
		if queryBuilder.Len() > 0 {
			queryBuilder.WriteString(" AND ")
		}
		queryBuilder.WriteString(fieldName)
		queryBuilder.WriteString(" = ?")
		*queryParameters = append(*queryParameters, fieldValue)
		added = true
	}

	return
}
