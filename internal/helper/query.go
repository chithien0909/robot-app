package helper

import (
	"fmt"
	"github.com/fatih/structs"
	"golang.org/x/net/html"
	"robot-app/internal/model"
	"strconv"
	"strings"
)

func BuildFilters(filters []*model.Filter) (fields []string, values []interface{}) {
	for _, filter := range filters {
		if filter == nil {
			continue
		}
		if filter.Value == nil {
			fields = append(fields, fmt.Sprintf("%s %s", filter.Key, filter.Method))
			continue
		}
		fields = append(fields, fmt.Sprintf("%s %s ?", filter.Key, filter.Method))
		values = append(values, filter.Value)
	}
	return fields, values
}

func BuildSearch(q string, searchFields []string) (filters []*model.Filter) {
	if q == "" {
		return nil
	}
	q = strings.ToLower(q)
	for _, searchField := range searchFields {
		filters = append(filters, &model.Filter{
			Key:    fmt.Sprintf("LOWER(%s)", searchField),
			Value:  strings.ToLower(fmt.Sprint("%" + html.EscapeString(q) + "%")),
			Method: "LIKE",
		})
	}
	return filters
}

func BuildFiltersFromReq(target interface{}) (filters []*model.Filter) {
	s := structs.New(target)
	for _, f := range s.Fields() {
		name := s.Field(f.Name())
		filterTag := name.Tag("filter")
		op := name.Tag("op")
		if op == "" {
			op = "="
		}
		if !f.IsZero() && filterTag != "" {
			filters = append(filters, &model.Filter{
				Key:    filterTag,
				Method: op,
				Value:  f.Value(),
			})
		}
	}
	return filters
}

func BuildQuery(q string, filters []*model.Filter, sort *model.Sort, pagination *model.Pagination) *model.Query {
	query := &model.Query{}
	//Search by keyword
	query.SetQ(q)
	//Build Filters
	query.SetFilters(filters)
	//Build sort
	if sort != nil && sort.Key != "" {
		query.SetSort(sort)
	}
	query.SetPagination(pagination)
	return query
}

func BuildPagination(page, limit int) *model.Pagination {
	return &model.Pagination{
		Page:   page,
		Limit:  limit,
		Offset: (page - 1) * limit,
	}
}

/**
 * GetSliceOfValue
 * @param value: string
 * @param typeValue
 */
func GetSliceStr(value string, caseSensitive bool) []string {
	value = strings.TrimSpace(value)
	if !caseSensitive {
		value = strings.ToLower(value)
	}
	var result []string
	for _, v := range strings.Split(value, ",") {
		result = append(result, strings.TrimSpace(v))
	}
	return result
}

func GetSliceInt(value string) []int {
	value = strings.TrimSpace(value)
	var result []int
	for _, v := range strings.Split(value, ",") {
		i, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			continue
		}
		result = append(result, i)
	}
	return result
}

func GetSliceFloat(value string) []float64 {
	value = strings.TrimSpace(value)
	var result []float64
	for _, v := range strings.Split(value, ",") {
		f, err := strconv.ParseFloat(strings.TrimSpace(v), 64)
		if err != nil {
			continue
		}
		result = append(result, f)
	}
	return result
}

func FilterInString(value, name string) *model.Filter {
	if value == "" {
		return nil
	}
	return &model.Filter{
		Key:    fmt.Sprintf("LOWER(%s)", name),
		Method: "IN",
		Value:  GetSliceStr(value, false),
	}
}

func FilterInInt(value, name string) *model.Filter {
	if value == "" {
		return nil
	}
	return &model.Filter{
		Key:    name,
		Method: "IN",
		Value:  GetSliceInt(value),
	}
}

func FilterInFloat(value, name string) *model.Filter {
	if value == "" {
		return nil
	}
	return &model.Filter{
		Key:    name,
		Method: "IN",
		Value:  GetSliceFloat(value),
	}
}

func FilterInBool(value bool, name string) *model.Filter {
	return &model.Filter{
		Key:    name,
		Method: "=",
		Value:  value,
	}
}

func FilterDateRange(value, name string) (filters []*model.Filter) {
	arr := strings.Split(value, ",")
	if len(arr) != 2 {
		return nil
	}
	start := strings.TrimSpace(arr[0])
	end := strings.TrimSpace(arr[1])

	return []*model.Filter{
		{
			Key:    fmt.Sprintf("DATE(%s)", name),
			Method: ">=",
			Value:  start,
		},
		{
			Key:    fmt.Sprintf("DATE(%s)", name),
			Method: "<=",
			Value:  end,
		},
	}
}

func BuildSortFromString(sortStr string) *model.Sort {
	if sortStr == "" {
		return &model.Sort{
			Key:    "id",
			SortBy: "DESC",
		}
	}
	firstChar := sortStr[0:1]
	if firstChar == "-" {
		return &model.Sort{
			Key:    strings.ReplaceAll(sortStr, "-", ""),
			SortBy: "DESC",
		}
	}
	return &model.Sort{
		Key:    sortStr,
		SortBy: "ASC",
	}
}
