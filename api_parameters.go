package espocrm

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
	OrderBy     *string `json:"orderBy,omitempty"`
	Select      *string `json:"select,omitempty"`
	Order       *Order  `json:"order,omitempty"`
	MaxSize     *int    `json:"maxSize,omitempty"`
	Offset      *int    `json:"offset,omitempty"`
	ReturnTotal *bool   `json:"total,omitempty"`
	Where       []Where `json:"where,omitempty"`
}

func NewParameters(opts ...ParamOption) *Parameters {
	params := &Parameters{}

	for _, opt := range opts {
		opt(params)
	}

	return params
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
