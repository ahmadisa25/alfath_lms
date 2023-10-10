package pagination

import (
	"alfath_lms/api/definitions"
	//"alfath_lms/api/instructor/domain/entity"
	"fmt"
	"strconv"
	"strings"

	"flamingo.me/dingo"
	"gorm.io/gorm"
)

type Module struct{}

type Paginator struct {
	PaginationReq definitions.PaginationRequest
	PaginationPrm definitions.PaginationParam
	db            *gorm.DB
}

func (paginator *Paginator) Inject(db *gorm.DB) {
	paginator.db = db
}

func (paginator *Paginator) Paginate(req definitions.PaginationRequest, prm definitions.PaginationParam) definitions.PaginationResult {
	paginator.PaginationReq = req
	paginator.PaginationPrm = prm

	whereClause := ""
	whereParams := map[string]interface{}{}

	sql := prm.Sql
	if req.SelectedColumns != "" {

		for _, selectField := range strings.Split(req.SelectedColumns, ",") {
			var isExist bool
			selectField = strings.ToLower(selectField)
			for _, element := range prm.SelectFields {
				if element == selectField {
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

	if req.Search != "" {
		if whereClause == "" {
			whereClause = "where "
		} else {
			whereClause = whereClause + " and "
		}

		i := 0
		for _, value := range prm.SearchFields {
			whereClause = whereClause + fmt.Sprintf("lower(%s)", value) + " like lower(@search_value) "
			whereParams["search_value"] = "%" + req.Search + "%"
			if i < len(prm.SearchFields)-1 {
				whereClause = whereClause + " or "
			}
			i++
		}
	}

	if req.Filter != "" {
		idx := 0
		if whereClause == "" {
			whereClause = "where "
		} else {
			whereClause = whereClause + " and "
		}

		filters := strings.Split(req.Filter, ",")
		for _, value := range filters {
			filterKey := strings.Split(value, ":")
			keyName := filterKey[0]
			_, keyOk := prm.FilterFields[keyName]
			if !keyOk {
				return definitions.PaginationResult{}
			}

			strIdx := strconv.Itoa(idx)
			whereClause = whereClause + " lower(" + prm.FilterFields[keyName] + ") like lower(@filter_value" + strIdx + ")"
			whereParams["filter_value"+strIdx] = "%" + filterKey[1] + "%"
			if idx < len(filters)-1 {
				whereClause = whereClause + " and "
			}
			idx++

		}

	}

	if whereClause != "" {
		sql = strings.Replace(sql, "-where-", whereClause, 1)
	} else {
		sql = strings.Replace(sql, "-where-", "", -1)
	}

	if req.OrderBy != "" {
		orderString := strings.Split(req.OrderBy, ":")
		sql = sql + " order by " + orderString[0] + " " + orderString[1]
	}

	if req.PerPage == "" {
		req.PerPage = "10"
	}

	perpage, convErr := strconv.Atoi(req.PerPage)

	sql = sql + " limit " + req.PerPage

	offset := 0
	offsetStr := "0"

	if req.Page != "" {
		pageInt, err := strconv.Atoi(req.Page)
		if err != nil {
			return definitions.PaginationResult{}
		}

		offset = (pageInt - 1) * perpage
		offsetStr = strconv.Itoa(offset)
		sql = sql + " offset " + offsetStr
	}

	if convErr != nil {
		return definitions.PaginationResult{}
	}

	total := 0
	mapResult := []interface{}{}
	rows, err := paginator.db.Raw(sql).Rows()
	if err != nil {
		return definitions.PaginationResult{}
	}
	if whereClause != "" {
		rows, err = paginator.db.Raw(sql, whereParams).Rows() //Limit in gorm just limits the rows you are taking from the database. It doesn't necessary add "Limit" to your SQL query probably, because if you iterate the rows with rows.Next(), rows that are outside of the limit is still accessed.
	}

	defer rows.Close()

	if err != nil {
		return definitions.PaginationResult{}
	}

	for rows.Next() {
		/*if i == req.PerPage {
			break
			//due to the quirk of Limit and Rows.Next logic, without this piece of code, it will just take everything
			//from the db table instead of limiting it. This only happens if you use rows.Next()
		}*/
		data := make(map[string]interface{})
		paginator.db.ScanRows(rows, &data)
		mapResult = append(mapResult, data)
		total++
	}

	result := definitions.PaginationResult{
		Data:    mapResult,
		Page:    1,
		PerPage: perpage,
		Total:   total,
		Status:  200,
	}

	return result
}

func (*Module) Configure(injector *dingo.Injector) {
	injector.Bind(new(Paginator)).ToProvider(func() *Paginator {
		return &Paginator{}
	}).In(dingo.Singleton)
}
