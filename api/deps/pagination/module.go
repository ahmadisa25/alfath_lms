package pagination

import(
	"alfath_lms/api/definitions"
	//"alfath_lms/api/instructor/domain/entity"
	"strings"
	"flamingo.me/dingo"
	"gorm.io/gorm"
	"fmt"
)


type Module struct{}

type Paginator struct {
	PaginationReq definitions.PaginationRequest
	PaginationPrm definitions.PaginationParam
	db *gorm.DB
}


func (paginator *Paginator) Inject(db *gorm.DB) {
	paginator.db = db
}

func (paginator *Paginator) Paginate(req definitions.PaginationRequest, prm definitions.PaginationParam) (definitions.PaginationResult){
	paginator.PaginationReq = req
	paginator.PaginationPrm = prm

	sql := prm.Sql
	if req.SelectedColumns != ""{

		for _, selectField := range strings.Split(req.SelectedColumns,","){
			var isExist bool
			for _, element := range prm.SelectFields{
				if element == selectField{
					continue
				}
			}
			if !isExist {
				return definitions.PaginationResult{}
			}
		}

		sql = strings.Replace(sql, "-select-", req.SelectedColumns, -1)
	}

	fmt.Println(sql)

	//var instructor entity.Instructor
	i := 0
	mapResult := []interface{}{}
	rows, err := paginator.db.Raw(sql).Rows()
	if err != nil{
		return definitions.PaginationResult{}
	}

	for rows.Next(){
		data := make(map[string]interface{})
		paginator.db.ScanRows(rows, &data)
		mapResult = append(mapResult,data)
		i++
	}

	result:= definitions.PaginationResult{
		Data: mapResult,
		Page: 1,
		PerPage: 10,
		Total: 1,
		Status: 200,
	}

	return result
}

func (*Module) Configure(injector *dingo.Injector) {
	injector.Bind(new(Paginator)).ToProvider(func() *Paginator {
		return &Paginator{}
	}).In(dingo.Singleton)
}
