package instance

import (
	"fmt"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	defaultFields  = []string{"timestamp", "cluster", "namespace", "app", "pod_name", "container_name", "host", "log"}
	defaultColumns = "timestamp, cluster, namespace, app, pod_name, container_name, host, fields_string.key, fields_string.value, fields_number.key, fields_number.value, log"

	fieldsMetric = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "kobs",
		Name:      "klogs_fields_total",
		Help:      "Number how often a field was used in a query.",
	}, []string{"field"})
)

// parseLogsQuery parses the given query string and return the conditions for the where statement in the sql query. We
// are providing a very simple query language where the user can use "(", ")", "_not_", "_and_" and "_or_" operators.
// Then we are splitting the string again for the other operators "=", "!=", ">", ">=", "<", "<=" and "~" which are used
// to check the value of a field.
// Once we have build all the conditions we concate all the strings to the final sql statement for the where clause.
func parseLogsQuery(query string, materializedColumns []string) (string, error) {
	var newOpenBrackets []string
	openBrackets := strings.Split(query, "(")
	for _, openBracket := range openBrackets {
		var newCloseBrackets []string
		closeBrackets := strings.Split(openBracket, ")")
		for _, closeBracket := range closeBrackets {
			var newNots []string
			nots := strings.Split(closeBracket, "_not_")
			for _, not := range nots {
				var newAnds []string
				ands := strings.Split(not, "_and_")
				for _, and := range ands {
					var newOrs []string
					ors := strings.Split(and, "_or_")
					for _, or := range ors {
						condition, err := splitOperator(or, materializedColumns)
						if err != nil {
							return "", err
						}

						newOrs = append(newOrs, condition)
					}
					newAnds = append(newAnds, strings.Join(newOrs, " OR "))
				}
				newNots = append(newNots, strings.Join(newAnds, " AND "))
			}
			newCloseBrackets = append(newCloseBrackets, strings.Join(newNots, " NOT "))
		}
		newOpenBrackets = append(newOpenBrackets, strings.Join(newCloseBrackets, ")"))
	}

	return strings.Join(newOpenBrackets, "("), nil
}

// splitOperator splits the given string by the following operators "=", "!=", ">", ">=", "<", "<=" and "~". If the
// result is a slice with two items we found the operator which was used by the user to check the value of a field. So
// that we pass the key (first item), value (second item) and the operator to the handleConditionParts to build the
// where condition.
func splitOperator(condition string, materializedColumns []string) (string, error) {
	greaterThanOrEqual := strings.Split(condition, ">=")
	if len(greaterThanOrEqual) == 2 {
		return handleConditionParts(greaterThanOrEqual[0], greaterThanOrEqual[1], ">=", materializedColumns)
	}

	greaterThan := strings.Split(condition, ">")
	if len(greaterThan) == 2 {
		return handleConditionParts(greaterThan[0], greaterThan[1], ">", materializedColumns)
	}

	lessThanOrEqual := strings.Split(condition, "<=")
	if len(lessThanOrEqual) == 2 {
		return handleConditionParts(lessThanOrEqual[0], lessThanOrEqual[1], "<=", materializedColumns)
	}

	lessThan := strings.Split(condition, "<")
	if len(lessThan) == 2 {
		return handleConditionParts(lessThan[0], lessThan[1], "<", materializedColumns)
	}

	ilike := strings.Split(condition, "=~")
	if len(ilike) == 2 {
		return handleConditionParts(ilike[0], ilike[1], "=~", materializedColumns)
	}

	notEqual := strings.Split(condition, "!=")
	if len(notEqual) == 2 {
		return handleConditionParts(notEqual[0], notEqual[1], "!=", materializedColumns)
	}

	notIlike := strings.Split(condition, "!~")
	if len(notIlike) == 2 {
		return handleConditionParts(notIlike[0], notIlike[1], "!~", materializedColumns)
	}

	regex := strings.Split(condition, "~")
	if len(regex) == 2 {
		return handleConditionParts(regex[0], regex[1], "~", materializedColumns)
	}

	equal := strings.Split(condition, "=")
	if len(equal) == 2 {
		return handleConditionParts(equal[0], equal[1], "=", materializedColumns)
	}

	if strings.Contains(condition, "_exists_ ") {
		return handleExistsCondition(strings.TrimLeft(strings.TrimSpace(condition), "_exists_ "), materializedColumns), nil
	}

	if strings.TrimSpace(condition) == "" {
		return "", nil
	}

	return "", fmt.Errorf("invalid operator: %s", condition)
}

// handleConditionParts converts the given key, value and operator to it's sql representation. This is required because
// some fields like "timestamp", "cluster", "namespace", etc. are a seperate column in the sql table, where others like
// "content.level" or "content.response_code" are only available via the fields_strings / fields_numbers column. For
// these nested columns we have to use a special query syntax. We also have to use the match function when the operator
// is "~" which says that the user checks the field value against a regular expression.
//
// See: https://gist.github.com/alexey-milovidov/d6ffc9e0bc0bc72dd7bca90e76e3b83b
// See: https://clickhouse.tech/docs/en/sql-reference/functions/string-search-functions/#matchhaystack-pattern
func handleConditionParts(key, value, operator string, materializedColumns []string) (string, error) {
	key = strings.TrimSpace(key)
	value = strings.TrimSpace(value)

	// The kobs_klogs_fields_total metric can be used to determine how often a field is used. This information can
	// then be used to create an additional column for this field via the following SQL commands:
	// ALTER TABLE logs.logs ON CLUSTER '{cluster}' ADD COLUMN <FIELD> String DEFAULT fields_string.value[indexOf(fields_string.key, '<FIELD>')];
	// ALTER TABLE logs.logs ON CLUSTER '{cluster}' ADD COLUMN <FIELD> String DEFAULT fields_number.value[indexOf(fields_number.key, '<FIELD>')];
	fieldsMetric.WithLabelValues(key).Inc()

	if contains(defaultFields, key) || contains(materializedColumns, key) {
		if operator == "=~" {
			return fmt.Sprintf("%s ILIKE %s", key, value), nil
		}

		if operator == "!~" {
			return fmt.Sprintf("%s NOT ILIKE %s", key, value), nil
		}

		if operator == "~" {
			return fmt.Sprintf("match(%s, %s)", key, value), nil
		}

		return fmt.Sprintf("%s%s%s", key, operator, value), nil
	}

	if value != "" && string(value[0]) == "'" && string(value[len(value)-1]) == "'" {
		if operator == "=~" {
			return fmt.Sprintf("fields_string.value[indexOf(fields_string.key, '%s')] ILIKE %s", key, value), nil
		}

		if operator == "!~" {
			return fmt.Sprintf("fields_string.value[indexOf(fields_string.key, '%s')] NOT ILIKE %s", key, value), nil
		}

		if operator == "~" {
			return fmt.Sprintf("match(fields_string.value[indexOf(fields_string.key, '%s')], %s)", key, value), nil
		}

		return fmt.Sprintf("fields_string.value[indexOf(fields_string.key, '%s')] %s %s", key, operator, value), nil
	}

	if operator == "=~" {
		return fmt.Sprintf("fields_number.value[indexOf(fields_number.key, '%s')] ILIKE %s", key, value), nil
	}

	if operator == "!~" {
		return fmt.Sprintf("fields_number.value[indexOf(fields_number.key, '%s')] NOT ILIKE %s", key, value), nil
	}

	if operator == "~" {
		return fmt.Sprintf("match(fields_number.value[indexOf(fields_number.key, '%s')], %s)", key, value), nil
	}

	return fmt.Sprintf("fields_number.value[indexOf(fields_number.key, '%s')] %s %s", key, operator, value), nil
}

func handleExistsCondition(key string, materializedColumns []string) string {
	if contains(defaultFields, key) || contains(materializedColumns, key) {
		return fmt.Sprintf("%s IS NOT NULL", key)
	}

	return fmt.Sprintf("(has(fields_string.key, '%s') = 1 OR has(fields_number.key, '%s') = 1)", key, key)
}

func parseOrder(order, orderBy string, materializedColumns []string) string {
	if order == "" || orderBy == "" {
		return "timestamp DESC"
	}

	if order == "ascending" {
		order = "ASC"
	} else {
		order = "DESC"
	}

	orderBy = strings.TrimSpace(orderBy)
	if contains(defaultFields, orderBy) || contains(materializedColumns, orderBy) {
		return fmt.Sprintf("%s %s", orderBy, order)
	}

	return fmt.Sprintf("fields_string.value[indexOf(fields_string.key, '%s')] %s, fields_number.value[indexOf(fields_number.key, '%s')] %s", orderBy, order, orderBy, order)
}
