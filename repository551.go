package repository551

import (
	"errors"
	"github.com/go51/model551"
	"github.com/go51/mysql551"
	"github.com/go51/string551"
)

type Repository struct{}

var repositoryInstance *Repository

func Load() *Repository {
	if repositoryInstance != nil {
		return repositoryInstance
	}

	repositoryInstance = &Repository{}

	return repositoryInstance
}

func (r *Repository) Find(db *mysql551.Mysql, mInfo *model551.ModelInformation, id int64) interface{} {

	sql := mInfo.SqlInformation.Select + " AND `" + mInfo.TableInformation.PrimaryKey + "` = ?"
	param := []interface{}{id}

	rows := db.Query(sql, param...)
	defer rows.Close()
	if !rows.Next() {
		return nil
	}

	model := mInfo.NewFuncPointer()
	modelScan, ok := model.(model551.ScanInterface)
	if !ok {
		panic(errors.New("Not found: 'T.Scan(rows *sql.rows) error' method"))
	}

	err := modelScan.Scan(rows)
	if err != nil {
		panic(err)
	}

	return model

}

func (r *Repository) FindBy(db *mysql551.Mysql, mInfo *model551.ModelInformation, param map[string]interface{}) []interface{} {

	sql := mInfo.SqlInformation.Select
	var values []interface{}
	for key, value := range param {
		sql = string551.Join(sql, " AND `"+key+"` = ?")
		values = append(values, value)
	}

	rows := db.Query(sql, values...)
	defer rows.Close()

	var models []interface{}

	for rows.Next() {
		model := mInfo.NewFuncPointer()
		modelScan, ok := model.(model551.ScanInterface)
		if !ok {
			panic(errors.New("Not found: 'T.Scan(rows *sql.rows) error' method"))
		}

		err := modelScan.Scan(rows)
		if err != nil {
			panic(err)
		}

		models = append(models, model)
	}

	return models

}

func (r *Repository) Create(db *mysql551.Mysql, mInfo *model551.ModelInformation, model interface{}) interface{} {

	sql := mInfo.SqlInformation.Insert
	sqlValueModel, ok := model.(model551.ValuesInterface)
	if !ok {
		panic(errors.New("Not found: 'T.SqlValues(sqlType SqlType) []interface{}' method"))
	}
	param := sqlValueModel.SqlValues(model551.SQL_INSERT)

	_, lastInsertId := db.Exec(sql, param...)

	idModel, ok := model.(model551.PrimaryInterface)
	if !ok {
		panic(errors.New("Not found: 'T.SetId(int64)', 'T.Id() int64' method"))
	}

	idModel.SetId(lastInsertId)

	return model
}

func (r *Repository) Update(db *mysql551.Mysql, mInfo *model551.ModelInformation, model interface{}) interface{} {

	sql := mInfo.SqlInformation.Update
	sqlValueModel, ok := model.(model551.ValuesInterface)
	if !ok {
		panic(errors.New("Not found: 'T.SqlValues(sqlType SqlType) []interface{}' method"))
	}
	param := sqlValueModel.SqlValues(model551.SQL_UPDATE)

	db.Exec(sql, param...)

	return model
}
