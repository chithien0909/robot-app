package model

type Filter struct {
	Key      string
	Value    interface{}
	Method   string
	OrFilter []*Filter
}

type Sort struct {
	Key    string
	SortBy string
}

type Query struct {
	Pagination   *Pagination
	Sort         *Sort
	Filters      []*Filter
	SearchFields []string
}

type Pagination struct {
	Page   int
	Limit  int
	Offset int
}

func (q *Query) AddFilter(f *Filter) {
	q.Filters = append(q.Filters, f)
}

func (q *Query) SetSort(s *Sort) {
	q.Sort = s
}

func (q *Query) SetSearchFields(searchFields ...string) {
	q.SearchFields = searchFields
	if len(searchFields) == 0 {
		q.SearchFields = []string{"*"}
	}
}

func (q *Query) GetPagination() *Pagination {
	if q.Pagination == nil {
		q.Pagination = &Pagination{
			Page:   1,
			Limit:  10,
			Offset: 0,
		}
	}

	if q.Pagination.Page <= 0 {
		q.Pagination.Page = 1
	}

	if q.Pagination.Limit <= 0 {
		q.Pagination.Limit = 10
	}

	q.Pagination.Offset = (q.Pagination.Page - 1) * q.Pagination.Limit

	return q.Pagination
}

func (q *Query) SetFilters(filters []*Filter) {
	q.Filters = filters
}

func (q *Query) SetPagination(p *Pagination) {
	q.Pagination = p
}
