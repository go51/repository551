package repository551

import (
	"errors"
	"github.com/go51/model551"
	"github.com/go51/mysql551"
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
