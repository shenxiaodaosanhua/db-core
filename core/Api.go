package core

import (
	"db-core/pbfiles"
	"fmt"
	"gorm.io/gorm"
)

type Select struct {
	Sql string `yaml:"sql"`
}

type API struct {
	Name   string  `yaml:"name"`
	Table  string  `yaml:"table"`
	Sql    string  `yaml:"sql"`
	Select *Select `yaml:"select"`
	DB     *gorm.DB
}

func (a *API) getDB() *gorm.DB {
	if a.DB == nil {
		return GormDB
	}
	return a.DB
}

func (a *API) SetDB(db *gorm.DB) {
	a.DB = db
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

//deprecated
// QueryByTableName 未来将要过期
func (a *API) QueryByTableName(params *pbfiles.SimpleParams) ([]map[string]interface{}, error) {
	dbResult := make([]map[string]interface{}, 0)
	db := a.getDB().Table(a.Table)

	paramMap := params.Params.AsMap()
	for key, value := range paramMap {
		db = db.Where(key, value)
	}

	db.Find(&dbResult)

	return dbResult, db.Error
}

func (a *API) QueryBySql(params *pbfiles.SimpleParams) ([]map[string]interface{}, error) {
	result := make([]map[string]interface{}, 0)
	db := a.getDB().Raw(a.Sql, params.Params.AsMap()).Find(&result)

	return result, db.Error
}

func (a *API) ExecBySql(params *pbfiles.SimpleParams) (int64, map[string]interface{}, error) {
	if a.Select != nil {
		selectKey := make(map[string]interface{})
		var rows int64 = 0
		err := a.getDB().Transaction(func(tx *gorm.DB) error {
			db := tx.Exec(a.Sql, params.Params.AsMap())
			if db.Error != nil {
				return db.Error
			}

			rows = db.RowsAffected
			db = tx.Raw(a.Select.Sql).Find(&selectKey)
			if db.Error != nil {
				return db.Error
			}

			return nil
		})
		if err != nil {
			return 0, nil, err
		}
		return rows, selectKey, err
	} else {
		result := a.getDB().Exec(a.Sql, params.Params.AsMap())
		return result.RowsAffected, nil, result.Error
	}
}
