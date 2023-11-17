package definitions

import "go.mongodb.org/mongo-driver/bson/primitive"

type GenericAPIMessage struct {
	Status  int
	Message string
}

type GenericCreationMessage struct {
	Status     int
	InstanceID int
}

type GenericMongoCreationMessage struct {
	Status     int
	InstanceID primitive.ObjectID
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
	Total   int64
	Status  int
}
