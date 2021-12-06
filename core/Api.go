package core

type API struct {
	Name  string `yaml:"name"`
	Table string `yaml:"table"`
	Sql   string `yaml:"sql"`
}

func (a *API) QueryByTableName() ([]map[string]interface{}, error) {
	dbResult := make([]map[string]interface{}, 0)
	db := GormDB.Table(a.Table).Find(&dbResult)
	return dbResult, db.Error
}
