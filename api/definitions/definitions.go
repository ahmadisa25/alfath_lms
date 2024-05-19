package definitions

import (
	"context"
	"time"

	"flamingo.me/flamingo/v3/framework/web"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MiddlewareFunc func(ctx context.Context, req *web.Request, actionStack []web.Action) web.Result

type GenericAPIMessage struct {
	Status  int
	Message string
}

type GenericGetMessage[T any] struct {
	Status int
	Data   T
}

type LoginResponse struct {
	Status       int
	Message      string
	Token        string
	RefreshToken string
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
	Sql              string
	SearchFields     map[string]string
	SelectFields     StringSlice
	FilterFields     map[string]string
	NullFilterFields map[string]bool
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

type SubmittedQuizAnswer struct {
	Answer    string
	StudentID string
	CreatedAt time.Time
}

type SimpleDashboardData struct {
	TotalCourses     int64
	TotalStudents    int64
	TotalSubmissions int64
}
