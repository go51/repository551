package repository551_test

import (
	"database/sql"
	"github.com/go51/model551"
	"github.com/go51/mysql551"
	"github.com/go51/repository551"
	"os"
	"testing"
	"time"
)

var db *mysql551.Mysql
var mm *model551.Model

func TestMain(m *testing.M) {

	dbConfig := &mysql551.Config{
		Host:     "tcp(localhost:3306)",
		User:     "root",
		Password: "",
		Database: "repository551",
	}
	db = mysql551.New(dbConfig)
	db.Open()
	defer db.Close()

	// Truncate
	sql := "TRUNCATE TABLE sample"
	_, _ = db.Exec(sql)
	sql = "TRUNCATE TABLE sample_delete"
	_, _ = db.Exec(sql)

	// Insert
	sql = "INSERT INTO sample (`group_id`, `name`, `description`) VALUES (?, ?, ?)"
	_, _ = db.Exec(sql, 2, "pubapp.biz_1", "domain_1")
	_, _ = db.Exec(sql, 1, "pubapp.biz_2", "domain_2")
	_, _ = db.Exec(sql, 2, "pubapp.biz_3", "domain_3")
	_, _ = db.Exec(sql, 2, "pubapp.biz_4", "domain_4")
	_, _ = db.Exec(sql, 3, "pubapp.biz_5", "domain_5")

	mm = model551.Load()
	mm.Add(NewSampleModel, NewSampleModelPointer)

	code := m.Run()

	os.Exit(code)

}

func TestLoad(t *testing.T) {
	m1 := repository551.Load()
	m2 := repository551.Load()

	if m1 == nil {
		t.Errorf("インスタンスの作成に失敗しました。")
	}
	if m2 == nil {
		t.Errorf("インスタンスの作成に失敗しました。")
	}
}

func BenchmarkLoad(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = repository551.Load()
	}
}

type Sample struct {
	Id          int64
	GroupId     int64
	Name        string
	Description string
	DeletedAt   time.Time `db_delete:"true"`
}

func NewSampleModel() interface{} {
	return Sample{}
}

func NewSampleModelPointer() interface{} {
	return &Sample{}
}

func (m *Sample) Scan(rows *sql.Rows) error {

	err := rows.Scan(
		&m.Id,
		&m.GroupId,
		&m.Name,
		&m.Description,
	)

	return err
}

func TestFind(t *testing.T) {

	repo := repository551.Load()
	mtSample := mm.Get("Sample")

	ret := repo.Find(db, mtSample, 1)
	if ret == nil {
		t.Errorf("取得に失敗しました。")
	}

	sample, _ := ret.(*Sample)
	if sample.Id != 1 {
		t.Errorf("取得に失敗しました。")
	}
	if sample.GroupId != 2 {
		t.Errorf("取得に失敗しました。")
	}
	if sample.Name != "pubapp.biz_1" {
		t.Errorf("取得に失敗しました。")
	}
	if sample.Description != "domain_1" {
		t.Errorf("取得に失敗しました。")
	}

	ret = repo.Find(db, mtSample, 2)
	if ret == nil {
		t.Errorf("取得に失敗しました。")
	}

	sample, _ = ret.(*Sample)
	if sample.Id != 2 {
		t.Errorf("取得に失敗しました。")
	}
	if sample.GroupId != 1 {
		t.Errorf("取得に失敗しました。")
	}
	if sample.Name != "pubapp.biz_2" {
		t.Errorf("取得に失敗しました。")
	}
	if sample.Description != "domain_2" {
		t.Errorf("取得に失敗しました。")
	}

}

func TestFindBy(t *testing.T) {

	repo := repository551.Load()
	mtSample := mm.Get("Sample")

	param := map[string]interface{}{
		"group_id": 1,
	}
	ret := repo.FindBy(db, mtSample, param)
	if ret == nil {
		t.Errorf("取得に失敗しました。")
	}
	if len(ret) != 1 {
		t.Errorf("取得に失敗しました。")
	}

	param = map[string]interface{}{
		"group_id": 2,
	}
	ret = repo.FindBy(db, mtSample, param)
	if ret == nil {
		t.Errorf("取得に失敗しました。")
	}
	if len(ret) != 3 {
		t.Errorf("取得に失敗しました。")
	}

	param = map[string]interface{}{
		"group_id": 3,
	}
	ret = repo.FindBy(db, mtSample, param)
	if ret == nil {
		t.Errorf("取得に失敗しました。")
	}
	if len(ret) != 1 {
		t.Errorf("取得に失敗しました。")
	}

	param = map[string]interface{}{
		"group_id": 4,
	}
	ret = repo.FindBy(db, mtSample, param)
	if ret != nil {
		t.Errorf("取得に失敗しました。")
	}
	if len(ret) != 0 {
		t.Errorf("取得に失敗しました。")
	}

}
