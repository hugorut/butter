package data

import (
	"github.com/hugorut/butter/sys"
	"database/sql"

	"github.com/jinzhu/gorm"
)

type ORM interface {
	Close() error
	DB() *sql.DB
	Preload(column string, conditions ...interface{}) ORM
	Where(query interface{}, args ...interface{}) ORM
	Or(query interface{}, args ...interface{}) ORM
	Not(query interface{}, args ...interface{}) ORM
	Limit(limit interface{}) ORM
	Offset(offset interface{}) ORM
	Order(value interface{}, reorder ...bool) ORM
	Select(query interface{}, args ...interface{}) ORM
	Omit(columns ...string) ORM
	Group(query string) ORM
	Having(query string, values ...interface{}) ORM
	Joins(query string, args ...interface{}) ORM
	Attrs(attrs ...interface{}) ORM
	Assign(attrs ...interface{}) ORM
	First(out interface{}, where ...interface{}) ORM
	Last(out interface{}, where ...interface{}) ORM
	Find(out interface{}, where ...interface{}) ORM
	Scan(dest interface{}) ORM
	Row() *sql.Row
	Rows() (*sql.Rows, error)
	ScanRows(rows *sql.Rows, result interface{}) error
	Pluck(column string, value interface{}) ORM
	Count(value interface{}) ORM
	Related(value interface{}, foreignKeys ...string) ORM
	FirstOrInit(out interface{}, where ...interface{}) ORM
	FirstOrCreate(out interface{}, where ...interface{}) ORM
	Update(attrs ...interface{}) ORM
	Updates(values interface{}, ignoreProtectedAttrs ...bool) ORM
	UpdateColumn(attrs ...interface{}) ORM
	UpdateColumns(values interface{}) ORM
	Save(value interface{}) ORM
	Create(value interface{}) ORM
	Delete(value interface{}, where ...interface{}) ORM
	Raw(sql string, values ...interface{}) ORM
	Exec(sql string, values ...interface{}) ORM
	Model(value interface{}) ORM
	Table(name string) ORM
	Debug() ORM
	Begin() ORM
	Commit() ORM
	Get(name string) (value interface{}, ok bool)
}

type DB interface {
	DbBeginner
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type DbBeginner interface {
	Begin() (*sql.Tx, error)
}

type Tx interface {
	Commit() error
	Rollback() error
}

type GormORM struct {
	Gorm *gorm.DB
}

func (db *GormORM) Close() error {
	return db.Gorm.Close()
}

func (db *GormORM) DB() *sql.DB {
	return db.Gorm.DB()
}

func (db *GormORM) Preload(column string, conditions ...interface{}) ORM {
	return &GormORM{db.Gorm.Preload(column, conditions...)}
}

func (db *GormORM) Where(query interface{}, args ...interface{}) ORM {
	return &GormORM{db.Gorm.Where(query, args...)}
}

func (db *GormORM) Or(query interface{}, args ...interface{}) ORM {
	return &GormORM{db.Gorm.Or(query, args...)}
}

func (db *GormORM) Not(query interface{}, args ...interface{}) ORM {
	return &GormORM{db.Gorm.Not(query, args...)}
}

func (db *GormORM) Limit(limit interface{}) ORM {
	return &GormORM{db.Gorm.Limit(limit)}
}

func (db *GormORM) Offset(offset interface{}) ORM {
	return &GormORM{db.Gorm.Offset(offset)}
}

func (db *GormORM) Order(value interface{}, reorder ...bool) ORM {
	return &GormORM{db.Gorm.Order(value, reorder...)}
}

func (db *GormORM) Select(query interface{}, args ...interface{}) ORM {
	return &GormORM{db.Gorm.Select(query, args...)}
}

func (db *GormORM) Omit(columns ...string) ORM {
	return &GormORM{db.Gorm.Omit(columns...)}
}

func (db *GormORM) Group(query string) ORM {
	return &GormORM{db.Gorm.Group(query)}
}

func (db *GormORM) Having(query string, values ...interface{}) ORM {
	return &GormORM{db.Gorm.Having(query, values...)}
}

func (db *GormORM) Joins(query string, args ...interface{}) ORM {
	return &GormORM{db.Gorm.Joins(query, args...)}
}

func (db *GormORM) Attrs(attrs ...interface{}) ORM {
	return &GormORM{db.Gorm.Attrs(attrs...)}
}

func (db *GormORM) Assign(attrs ...interface{}) ORM {
	return &GormORM{db.Gorm.Assign(attrs...)}
}

func (db *GormORM) First(out interface{}, where ...interface{}) ORM {
	db.Gorm.First(out, where...)
	return db
}

func (db *GormORM) Last(out interface{}, where ...interface{}) ORM {
	db.Gorm.Last(out, where...)
	return db
}

func (db *GormORM) Find(out interface{}, where ...interface{}) ORM {
	db.Gorm.Find(out, where...)
	return db
}

func (db *GormORM) Scan(dest interface{}) ORM {
	db.Gorm.Scan(dest)
	return db
}

func (db *GormORM) Row() *sql.Row {
	return db.Gorm.Row()
}

func (db *GormORM) Rows() (*sql.Rows, error) {
	return db.Gorm.Rows()
}

func (db *GormORM) ScanRows(rows *sql.Rows, result interface{}) error {
	return db.Gorm.ScanRows(rows, result)
}

func (db *GormORM) Pluck(column string, value interface{}) ORM {
	db.Gorm.Pluck(column, value)
	return db
}

func (db *GormORM) Count(value interface{}) ORM {
	db.Gorm.Count(value)
	return db
}

func (db *GormORM) Related(value interface{}, foreignKeys ...string) ORM {
	db.Gorm.Related(value, foreignKeys...)
	return db
}

func (db *GormORM) FirstOrInit(out interface{}, where ...interface{}) ORM {
	db.Gorm.FirstOrInit(out, where...)
	return db
}

func (db *GormORM) FirstOrCreate(out interface{}, where ...interface{}) ORM {
	db.Gorm.FirstOrCreate(out, where...)
	return db
}

func (db *GormORM) Update(attrs ...interface{}) ORM {
	db.Gorm.Update(attrs...)
	return db
}

func (db *GormORM) Updates(values interface{}, ignoreProtectedAttrs ...bool) ORM {
	db.Gorm.Updates(values, ignoreProtectedAttrs...)
	return db
}

func (db *GormORM) UpdateColumn(attrs ...interface{}) ORM {
	db.Gorm.UpdateColumn(attrs...)
	return db
}

func (db *GormORM) UpdateColumns(values interface{}) ORM {
	db.Gorm.UpdateColumns(values)
	return db
}

func (db *GormORM) Save(value interface{}) ORM {
	db.Gorm.Save(value)
	return db
}

func (db *GormORM) Create(value interface{}) ORM {
	db.Gorm.Create(value)
	return db
}

func (db *GormORM) Delete(value interface{}, where ...interface{}) ORM {
	db.Gorm.Delete(value, where...)
	return db
}

func (db *GormORM) Raw(sql string, values ...interface{}) ORM {
	db.Gorm.Raw(sql, values...)
	return db
}

func (db *GormORM) Exec(sql string, values ...interface{}) ORM {
	db.Gorm.Exec(sql, values...)
	return db
}

func (db *GormORM) Model(value interface{}) ORM {
	return &GormORM{db.Gorm.Model(value)}
}

func (db *GormORM) Table(name string) ORM {
	return &GormORM{db.Gorm.Table(name)}
}

func (db *GormORM) Debug() ORM {
	return &GormORM{db.Gorm.Debug()}
}

func (db *GormORM) Begin() ORM {
	return &GormORM{db.Gorm.Begin()}
}

func (db *GormORM) Commit() ORM {
	return &GormORM{db.Gorm.Commit()}
}

func (db *GormORM) Get(name string) (value interface{}, ok bool) {
	return db.Gorm.Get(name)
}

// WrapSqlInGorm returns a GormORM interface from an existing connection
func WrapSqlInGorm(sqlConnection DB) (ORM, error) {
	db, err := gorm.Open("mysql", sqlConnection)
	return &GormORM{db}, err
}

// NewMySQLDBConnection returns a pointer to a mysql db
func NewMySQLDBConnection() (*sql.DB, error) {
	user := sys.EnvOrDefault("MYSQL_USER", "root")
	host := sys.EnvOrDefault("MYSQL_HOST", "127.0.0.1")
	port := sys.EnvOrDefault("MYSQL_PORT", "3306")
	database := sys.EnvOrDefault("MYSQL_DATABASE", "butter")
	password := sys.EnvOrDefault("MYSQL_PASSWORD", "")

	return sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+database+"?parseTime=true")
}
