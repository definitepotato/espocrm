package espocrm

import (
	"fmt"
	"net/url"
)

type Order string

const (
	Ascending Order = "asc"
	Decending Order = "desc"
)

type ParamOption func(p *Parameters)

func SetMaxSize(maxSize int) ParamOption {
	return func(param *Parameters) {
		param.MaxSize = &maxSize
	}
}

func SetOffset(offset int) ParamOption {
	return func(param *Parameters) {
		param.Offset = &offset
	}
}

func SetWhere(where []Where) ParamOption {
	return func(param *Parameters) {
		param.Where = append(param.Where, where...)
	}
}

func SetOrderBy(orderBy string) ParamOption {
	return func(param *Parameters) {
		param.OrderBy = &orderBy
	}
}

func SetSelect(selectAttributes string) ParamOption {
	return func(param *Parameters) {
		param.Select = &selectAttributes
	}
}

func SetTotal(total bool) ParamOption {
	return func(param *Parameters) {
		param.ReturnTotal = &total
	}
}

type Parameters struct {
	OrderBy     *string
	Select      *string
	Order       *Order
	MaxSize     *int
	Offset      *int
	ReturnTotal *bool
	Where       []Where
}

func NewParameters(opts ...ParamOption) *Parameters {
	params := &Parameters{}

	for _, opt := range opts {
		opt(params)
	}

	return params
}

func (params *Parameters) Encode() string {
	if params == nil {
		return ""
	}

	paramValues := url.Values{}

	if params.MaxSize != nil {
		maxSize := fmt.Sprintf("%v", *params.MaxSize)
		paramValues.Add("maxSize", maxSize)
	}

	if params.OrderBy != nil {
		paramValues.Add("orderBy", *params.OrderBy)
	}

	if params.Select != nil {
		paramValues.Add("select", *params.Select)
	}

	if params.Order != nil {
		order := "desc"

		if *params.Order == Ascending {
			order = "asc"
		}
		paramValues.Add("order", order)
	}

	if params.Offset != nil {
		offset := fmt.Sprintf("%d", params.Offset)
		paramValues.Add("offset", offset)
	}

	if params.ReturnTotal != nil {
		returnTotal := "false"

		if *params.ReturnTotal {
			returnTotal = "true"
		}
		paramValues.Add("returnTotal", returnTotal)
	}

	if params.Where != nil {
		for i := 0; i < len(params.Where); i++ {
			paramValues.Add("where[%d][type]=%s", string(params.Where[i].Type))
			paramValues.Add("where[%d][attribute]=%s", params.Where[i].Attribute)
			paramValues.Add("where[%d][value]=%s", params.Where[i].Value)
		}
	}

	return paramValues.Encode()
}

type FilterType string

const (
	Equals               FilterType = "equals"
	NotEquals            FilterType = "notEquals"
	GreaterThan          FilterType = "greaterThan"
	LessThan             FilterType = "lessThan"
	GreaterThanOrEquals  FilterType = "greaterThanOrEquals"
	LessThanOrEquals     FilterType = "lessThanOrEquals"
	IsNull               FilterType = "isNull"
	IsNotNull            FilterType = "isNotNull"
	IsTrue               FilterType = "isTrue"
	IsFalse              FilterType = "isFalse"
	LinkedWith           FilterType = "linkedWith"
	NotLinkedWith        FilterType = "notLinkedWith"
	IsLinked             FilterType = "isLinked"
	IsNotLinked          FilterType = "isNotLinked"
	In                   FilterType = "in"
	NotIn                FilterType = "notIn"
	Contains             FilterType = "contains"
	NotContains          FilterType = "notContains"
	StartsWith           FilterType = "startsWith"
	EndsWith             FilterType = "endsWith"
	Like                 FilterType = "like"
	NotLike              FilterType = "notLike"
	Or                   FilterType = "or"
	AndToday             FilterType = "andToday"
	Past                 FilterType = "past"
	Future               FilterType = "future"
	LastSevenDays        FilterType = "lastSevenDays"
	CurrentMonth         FilterType = "currentMonth"
	LastMonth            FilterType = "lastMonth"
	NextMonth            FilterType = "nextMonth"
	CurrentQuarter       FilterType = "currentQuarter"
	LastQuarter          FilterType = "lastQuarter"
	CurrentYear          FilterType = "currentYear"
	LastYear             FilterType = "lastYear"
	CurrentFiscalYear    FilterType = "currentFiscalYear"
	LastFiscalYear       FilterType = "lastFiscalYear"
	CurrentFiscalQuarter FilterType = "currentFiscalQuarter"
	LastFiscalQuarter    FilterType = "lastFiscalQuarter"
	LastXDays            FilterType = "lastXDays"
	NextXDays            FilterType = "nextXDays"
	OlderThanXDays       FilterType = "olderThanXDays"
	AfterXDays           FilterType = "afterXDays"
	Between              FilterType = "between"
	ArrayAnyOf           FilterType = "arrayAnyOf"
	ArrayNoneOf          FilterType = "arrayNoneOf"
	ArrayAllOf           FilterType = "arrayAllOf"
	ArrayIsEmpty         FilterType = "arrayIsEmpty"
	ArrayIsNotEmpty      FilterType = "arrayIsNotEmpty"
)

type Where struct {
	Type      FilterType `json:"type"`
	Attribute string     `json:"attribute"`
	Value     string     `json:"value"`
}
