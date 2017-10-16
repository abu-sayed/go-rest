package list

type Query struct {
	Select map[string]string
	Join map[string]string
	Where map[string]string
	GroupBy map[string]string
	Having map[string]string
	OrderBy map[string]string

	Params map[string]interface{}
	Store map[string]interface{}

	Fields map[string]func()
	Filters map[string]func(...interface{})
	Sorts map[string]func(string)
	Data map[string]func(interface{}, bool)

	Columns map[string]Column
	DefaultSort bool
}

type Column struct {
	Fields []string
	Default bool
}

type Response struct {
	Metadata
	Results []interface{}
}

type Metadata struct {
	Total int
	PerPage int
	Page int
}

type QueryService interface {
	Execute() Response
	Init()
	Filter() []int
	Gather([]int) []interface{}
	GetSelect() string
	GetJoin() string
	GetWhere() string
	GetGroupBy() string
	GetHaving() string
	GetOrderBy() string
}