package core

import (
	"db-core/pbfiles"
	"fmt"
)

type API struct {
	Name  string `yaml:"name"`
	Table string `yaml:"table"`
	Sql   string `yaml:"sql"`
}

func (a *API) Query(params *pbfiles.SimpleParams) ([]map[string]interface{}, error) {
	if a.Sql == "" || a.Table == "" {
		return nil, fmt.Errorf("error sql or table")
	}

	if a.Sql != "" {
		return a.QueryBySql(params)
	}

	return a.QueryByTableName(params)
}

func (a *API) QueryByTableName(params *pbfiles.SimpleParams) ([]map[string]interface{}, error) {
	dbResult := make([]map[string]interface{}, 0)
	db := GormDB.Table(a.Table)

	paramMap := params.Params.AsMap()
	for key, value := range paramMap {
		db = db.Where(key, value)
	}

	db.Find(&dbResult)

	return dbResult, db.Error
}

func (a *API) QueryBySql(params *pbfiles.SimpleParams) ([]map[string]interface{}, error) {
	dbResult := make([]map[string]interface{}, 0)
	db := GormDB.Raw(a.Sql, params.Params.AsMap()).Find(&dbResult)

	return dbResult, db.Error
}
