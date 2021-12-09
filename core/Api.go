package core

import "db-core/pbfiles"

type API struct {
	Name  string `yaml:"name"`
	Table string `yaml:"table"`
	Sql   string `yaml:"sql"`
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
