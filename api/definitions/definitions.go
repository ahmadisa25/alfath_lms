package definitions

type GenericAPIMessage struct {
	Status  int
	Message string
}

type GenericCreationMessage struct {
	Status     int
	InstanceID int
}

type StringSlice []string

type PaginationParam struct {
	Sql          string
	SearchFields map[string]string
	SelectFields StringSlice
	FilterFields map[string]string
}

type PaginationRequest struct {
	SelectedColumns string
	Page            string
	Search          string
	PerPage         string
	OrderBy         string
	Paginate        string
	Filter          string
}

type PaginationResult struct {
	Data    interface{}
	Page    int
	PerPage int
	Total   int
	Status  int
}
