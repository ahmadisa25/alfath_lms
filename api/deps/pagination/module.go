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

	whereClause := ""
	whereParams := map[string]interface{}{}

	sql := prm.Sql
	if req.SelectedColumns != ""{

		for _, selectField := range strings.Split(req.SelectedColumns,","){
			var isExist bool
			selectField = strings.ToLower(selectField)
			for _, element := range prm.SelectFields{
				if element == selectField{
					isExist = true
					break
				}
			}
			if !isExist {
				return definitions.PaginationResult{}
			}
		}
		sql = strings.Replace(sql, "-select-", req.SelectedColumns, -1)
	} else {
		sql = strings.Replace(sql, "-select-", "*", -1)
	}

	if req.Search !="" {
		fmt.Println(req.Search)
		if whereClause == ""{
			whereClause = "where "
		} else {
			whereClause = " and "
		}

		i:=0
		for _, value := range prm.SearchFields{
			fmt.Println(value)
			whereClause = whereClause + fmt.Sprintf("lower(%s)", value) + " like lower(@search_value) "
			whereParams["search_value"] = "%" + req.Search + "%"
			if i < len(prm.SearchFields)-1{
				whereClause = whereClause + " or "
			}
			i++
		}
	}
	if whereClause != ""{
		sql = strings.Replace(sql, "-where-", whereClause, 1)
	} else {
		sql = strings.Replace(sql, "-where-", "", -1)
	}
	//var instructor entity.Instructor
	i := 0
	mapResult := []interface{}{}
	rows, err := paginator.db.Raw(sql).Rows()
	if whereClause != "" {
		fmt.Println(whereParams)
		rows, err = paginator.db.Raw(sql, whereParams).Rows()
	} 

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
