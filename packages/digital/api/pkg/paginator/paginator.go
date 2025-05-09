package paginator

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math"
	"strconv"
	"strings"
)

type IMethods interface {
	Where(v string) IMethods
	Order(v string) IMethods
	Limit(v int) IMethods
	Select(v []string) IMethods
	Count(v any) struct {
		Error error
	}
	Preload(v any, vv ...any) IMethods
	Table(v string) IMethods
	Joins(v string) IMethods
	Offset(v any) IMethods
	Scan(v any) struct {
		Error error
	}
}

type functorFunc func(item map[string]any, tx *gorm.DB) *gorm.DB

type Paginator struct {
	Config         ConfigPaginator
	Table          string
	Limit          int
	Page           int
	Order          string
	Where          string
	UserData       string
	Join           string
	AcceptedFilter []string
	Select         []string
	Preloads       [][]interface{}
}

const (
	PAGE        = "page"
	LIMIT       = "limit"
	TOTAL_PAGES = "total"
	ORDERBY     = "order"
	FIELD       = "field"
	WHERE       = "where"
)

func (pag *Paginator) Configure(config ConfigPaginator) *Paginator {
	pag.Config = config
	return pag
}

func (pag *Paginator) SPage(page int) *Paginator {
	pag.Page = page
	return pag
}

func (pag *Paginator) SJoin(joinq string) *Paginator {
	pag.Join = joinq
	return pag
}

func (pag *Paginator) SSelect(slist []string) *Paginator {
	pag.Select = slist
	return pag
}

func (pag *Paginator) SLimit(limit int) *Paginator {
	pag.Limit = limit
	return pag
}

func (pag *Paginator) SUser(userData string) *Paginator {
	pag.UserData = userData
	return pag
}

func (pag *Paginator) SOrder(orderq string) *Paginator {
	pag.Order = pag.fromContextOrder(orderq)
	fmt.Println(pag.Order)
	return pag
}

func (pag *Paginator) SWhere(request string) *Paginator {
	pag.Where = pag.fromContextWhere(request)
	return pag
}

func (pag *Paginator) SAcceptedFilter(listing []string) *Paginator {
	pag.AcceptedFilter = listing
	return pag
}

func (pag *Paginator) STable(table string) *Paginator {
	pag.Table = table
	return pag
}

func (pag *Paginator) SPreloads(preloads [][]interface{}) *Paginator {
	pag.Preloads = preloads
	return pag
}

func (pag *Paginator) Pick(queries map[string][]string) *Paginator {
	if queries[PAGE] != nil {
		z, _ := strconv.Atoi(queries[PAGE][0])
		delete(queries, PAGE)
		pag.Page = z
	}
	if queries[LIMIT] != nil {
		z, _ := strconv.Atoi(queries[LIMIT][0])
		delete(queries, LIMIT)
		pag.Limit = z
	}

	if queries[ORDERBY] != nil {
		pag.SOrder(queries[ORDERBY][0])
		delete(queries, ORDERBY)
	}

	if queries[WHERE] != nil {
		pag.SWhere(queries[WHERE][0])
		delete(queries, WHERE)
	}

	return pag
}

// Ignite ?page=3&limit=10&order=id^asc
func (pag *Paginator) Ignite(obj *ObjectPaginator, w IMethods) error {
	var dataset []map[string]any
	q := w.Table(pag.Table)
	countResponse := new(int64)
	calculatedOffset := pag.Page * pag.Limit

	if pag.Where != "" {
		q.Where(pag.Where)
	}
	if pag.UserData != "" {
		q.Where(pag.UserData)
	}
	if pag.Order != "" {
		q.Order(pag.Order)
	}
	if pag.Limit != 0 {
		q.Limit(pag.Limit)
	}
	if pag.Join != "" {
		q.Joins(pag.Join)
	}

	if len(pag.Select) <= 0 {
		q.Select(pag.Select)
	}

	if len(pag.Preloads) > 0 {
		for _, v := range pag.Preloads {
			q.Preload(v[0].(string), v[1:]...)
		}
	}

	if countAll := q.Count(countResponse); countAll.Error != nil {
		return countAll.Error
	}

	obj.TotalPages = int(math.Floor(safeDivideInt(float64(int(*countResponse)), float64(pag.Limit))))
	obj.Page = pag.Page
	obj.Limit = pag.Limit
	obj.SortedBy = pag.Order

	if datasetRes := q.Offset(calculatedOffset).Scan(&dataset); datasetRes.Error != nil {
		return datasetRes.Error
	}
	obj.Data = dataset

	return nil
}

func safeDivideInt(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}

// a=desc|b=asc
func (pag *Paginator) fromContextOrder(orderStr string) string {
	result := ""
	splicedOrder := strings.Split(orderStr, "|")
	for k, v := range splicedOrder {
		splicedChunk := strings.Split(v, "^")
		if len(splicedChunk) >= 2 {
			result += fmt.Sprintf("%s.%s", pag.Table, splicedChunk[0]) + " " + splicedChunk[1]
			if k < len(splicedOrder)-1 {
				result += " ,"
			}
		}
	}
	return result
}

// id[1]|owner_user_hash['9fab84a340dc745aa74e0898985aa6cc190c3eee'^'9fab84a340dc745aa74e0898985aa6cc190c3eee']
func (pag *Paginator) fromContextWhere(querystring string) string {
	resultedQuery := ""

	queries := strings.Split(querystring, "|")

	for k, v := range queries {
		splicedChunk := strings.Split(
			strings.ReplaceAll(
				strings.ReplaceAll(
					v, "[", " "), "]", ""), " ")

		if len(pag.AcceptedFilter) > 0 && !includesBool(pag.AcceptedFilter, splicedChunk[0]) {
			continue
		}
		if len(splicedChunk) > 1 {
			if len(splicedChunk[1]) >= 1 {
				resultedQuery += accumulateOr(splicedChunk[0], strings.Split(splicedChunk[1], "^"))
			}
			if k < len(queries)-1 {
				resultedQuery += " AND "
			}
		}

	}
	return resultedQuery
}

func MergeTo[T any](instance *gorm.DB, obj *ObjectPaginator, toField string, functor functorFunc) error {
	if len(obj.Data) <= 0 {
		return nil
	}
	if obj.Data == nil || obj.Data[0][toField] != nil {
		return errors.New("list is invalid")
	}

	for _, v := range obj.Data {
		var result T
		if tx := functor(v, instance).Scan(&result); tx.Error != nil {
			return tx.Error
		}

		v[toField] = result
	}
	return nil
}

func accumulateOr(k string, list []string) string {
	result := ""
	for index, v := range list {
		result += fmt.Sprintf("%s = %s ", k, v)
		if index < len(list)-1 {
			result += "OR "
		}
	}
	return result
}

func includesBool[T comparable](list []T, item T) bool {
	if len(list) == 0 {
		return false
	}
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}

func NewPaginator() *Paginator {
	return &Paginator{}
}
