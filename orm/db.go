package orm

/* gorm wrapper for scalability (ex. transaction) */

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DB struct {
	impl *gorm.DB
}

func Open(dialect string, args ...interface{}) (*DB, error) {
	db, err := gorm.Open(dialect, args...)
	return &DB{
		impl: db,
	}, err
}

// Wrapped gorm.DB.New
func (s *DB) New() *DB {
	return wrapDB(s.impl.New())
}

// Wrapped gorm.DB.Close
func (s *DB) Close() error {
	return s.impl.Close()
}

// Wrapped gorm.DB.Error
func (s *DB) Error() error {
	return s.impl.Error
}

// Wrapped gorm.DB.Create
func (s *DB) Create(value interface{}) *DB {
	return wrapDB(s.impl.Create(value))
}

// Wrapped gorm.DB.Save
func (s *DB) Save(value interface{}) *DB {
	return wrapDB(s.impl.Save(value))
}

// Wrapped gorm.DB.Update
func (s *DB) Update(attrs ...interface{}) *DB {
	return wrapDB(s.impl.Update(attrs...))
}

// Wrapped gorm.DB.Updates
func (s *DB) Updates(values interface{}, ignoreProtectedAttrs ...bool) *DB {
	return wrapDB(s.impl.Updates(values, ignoreProtectedAttrs...))
}

// Wrapped gorm.DB.UpdateColumn
func (s *DB) UpdateColumn(attrs ...interface{}) *DB {
	return wrapDB(s.impl.UpdateColumn(attrs...))
}

// Wrapped gorm.DB.Delete
func (s *DB) Delete(value interface{}, where ...interface{}) *DB {
	return wrapDB(s.impl.Delete(value, where...))
}

func (s *DB) Count(value interface{}) *DB {
	return wrapDB(s.impl.Count(value))
}

// Wrapped gorm.DB.LogMode
func (s *DB) LogMode(enable bool) *DB {
	s.impl.LogMode(enable)
	return s
}

func (s *DB) Raw(sql string, values ...interface{}) *DB {
	return wrapDB(s.impl.Raw(sql, values...))
}

// Wrapped gorm.DB.Rows
func (s *DB) Rows() (*sql.Rows, error) {
	return s.impl.Rows()
}

// ------- Internal use

func wrapDB(db *gorm.DB) *DB {
	return &DB{
		impl: db,
	}
}
