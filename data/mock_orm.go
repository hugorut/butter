package data

import (
	sql "database/sql"
	"reflect"

	mock "github.com/stretchr/testify/mock"
)

// ORM is an autogenerated mock type for the ORM type
type MockORM struct {
	mock.Mock
}

// Assign provides a mock function with given fields: attrs
func (_m *MockORM) Assign(attrs ...interface{}) ORM {
	ret := _m.Called(attrs)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(...interface{}) ORM); ok {
		r0 = rf(attrs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Attrs provides a mock function with given fields: attrs
func (_m *MockORM) Attrs(attrs ...interface{}) ORM {
	ret := _m.Called(attrs)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(...interface{}) ORM); ok {
		r0 = rf(attrs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Begin provides a mock function with given fields:
func (_m *MockORM) Begin() ORM {
	ret := _m.Called()

	var r0 ORM
	if rf, ok := ret.Get(0).(func() ORM); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *MockORM) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *MockORM) Preload(column string, conditions ...interface{}) ORM {
	ret := _m.Called(column, conditions)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) ORM); ok {
		r0 = rf(column, conditions...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Commit provides a mock function with given fields:
func (_m *MockORM) Commit() ORM {
	ret := _m.Called()

	var r0 ORM
	if rf, ok := ret.Get(0).(func() ORM); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Count provides a mock function with given fields: value
func (_m *MockORM) Count(value interface{}) ORM {
	ret := _m.Called(value)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(interface{}) ORM); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Create provides a mock function with given fields: value
func (_m *MockORM) Create(value interface{}) ORM {
	ret := _m.Called(value)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(interface{}) ORM); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// DB provides a mock function with given fields:
func (_m *MockORM) DB() *sql.DB {
	ret := _m.Called()

	var r0 *sql.DB
	if rf, ok := ret.Get(0).(func() *sql.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.DB)
		}
	}

	return r0
}

// Debug provides a mock function with given fields:
func (_m *MockORM) Debug() ORM {
	ret := _m.Called()

	var r0 ORM
	if rf, ok := ret.Get(0).(func() ORM); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: value, where
func (_m *MockORM) Delete(value interface{}, where ...interface{}) ORM {
	ret := _m.Called(value, where)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) ORM); ok {
		r0 = rf(value, where...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Exec provides a mock function with given fields: _a0, values
func (_m *MockORM) Exec(_a0 string, values ...interface{}) ORM {
	ret := _m.Called(_a0, values)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(string, ...interface{}) ORM); ok {
		r0 = rf(_a0, values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Find provides a mock function with given fields: out, where
func (_m *MockORM) Find(out interface{}, where ...interface{}) ORM {
	ret := _m.Called(out, where)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) ORM); ok {
		r0 = rf(out, where...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// First provides a mock function with given fields: out, where
func (_m *MockORM) First(out interface{}, where ...interface{}) ORM {
	ret := _m.Called(out, where)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) ORM); ok {
		r0 = rf(out, where...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// FirstOrCreate provides a mock function with given fields: out, where
func (_m *MockORM) FirstOrCreate(out interface{}, where ...interface{}) ORM {
	ret := _m.Called(out, where)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) ORM); ok {
		r0 = rf(out, where...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// FirstOrInit provides a mock function with given fields: out, where
func (_m *MockORM) FirstOrInit(out interface{}, where ...interface{}) ORM {
	ret := _m.Called(out, where)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) ORM); ok {
		r0 = rf(out, where...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Get provides a mock function with given fields: name
func (_m *MockORM) Get(name string) (interface{}, bool) {
	ret := _m.Called(name)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Group provides a mock function with given fields: query
func (_m *MockORM) Group(query string) ORM {
	ret := _m.Called(query)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(string) ORM); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Having provides a mock function with given fields: query, values
func (_m *MockORM) Having(query string, values ...interface{}) ORM {
	ret := _m.Called(query, values)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(string, ...interface{}) ORM); ok {
		r0 = rf(query, values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Joins provides a mock function with given fields: query, args
func (_m *MockORM) Joins(query string, args ...interface{}) ORM {
	ret := _m.Called(query, args)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(string, ...interface{}) ORM); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Last provides a mock function with given fields: out, where
func (_m *MockORM) Last(out interface{}, where ...interface{}) ORM {
	ret := _m.Called(out, where)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) ORM); ok {
		r0 = rf(out, where...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Limit provides a mock function with given fields: limit
func (_m *MockORM) Limit(limit interface{}) ORM {
	ret := _m.Called(limit)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(interface{}) ORM); ok {
		r0 = rf(limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Model provides a mock function with given fields: value
func (_m *MockORM) Model(value interface{}) ORM {
	ret := _m.Called(value)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(interface{}) ORM); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Not provides a mock function with given fields: query, args
func (_m *MockORM) Not(query interface{}, args ...interface{}) ORM {
	ret := _m.Called(query, args)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) ORM); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Offset provides a mock function with given fields: offset
func (_m *MockORM) Offset(offset interface{}) ORM {
	ret := _m.Called(offset)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(interface{}) ORM); ok {
		r0 = rf(offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Omit provides a mock function with given fields: columns
func (_m *MockORM) Omit(columns ...string) ORM {
	ret := _m.Called(columns)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(...string) ORM); ok {
		r0 = rf(columns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Or provides a mock function with given fields: query, args
func (_m *MockORM) Or(query interface{}, args ...interface{}) ORM {
	ret := _m.Called(query, args)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) ORM); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Order provides a mock function with given fields: value, reorder
func (_m *MockORM) Order(value interface{}, reorder ...bool) ORM {
	ret := _m.Called(value, reorder)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(interface{}, ...bool) ORM); ok {
		r0 = rf(value, reorder...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Pluck provides a mock function with given fields: column, value
func (_m *MockORM) Pluck(column string, value interface{}) ORM {
	ret := _m.Called(column, value)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(string, interface{}) ORM); ok {
		r0 = rf(column, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Raw provides a mock function with given fields: _a0, values
func (_m *MockORM) Raw(_a0 string, values ...interface{}) ORM {
	ret := _m.Called(_a0, values)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(string, ...interface{}) ORM); ok {
		r0 = rf(_a0, values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Related provides a mock function with given fields: value, foreignKeys
func (_m *MockORM) Related(value interface{}, foreignKeys ...string) ORM {
	ret := _m.Called(value, foreignKeys)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(interface{}, ...string) ORM); ok {
		r0 = rf(value, foreignKeys...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Row provides a mock function with given fields:
func (_m *MockORM) Row() *sql.Row {
	ret := _m.Called()

	var r0 *sql.Row
	if rf, ok := ret.Get(0).(func() *sql.Row); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Row)
		}
	}

	return r0
}

// Rows provides a mock function with given fields:
func (_m *MockORM) Rows() (*sql.Rows, error) {
	ret := _m.Called()

	var r0 *sql.Rows
	if rf, ok := ret.Get(0).(func() *sql.Rows); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Rows)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: value
func (_m *MockORM) Save(value interface{}) ORM {
	ret := _m.Called(value)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(interface{}) ORM); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Scan provides a mock function with given fields: dest
func (_m *MockORM) Scan(dest interface{}) ORM {
	ret := _m.Called(dest)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(interface{}) ORM); ok {
		r0 = rf(dest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// ScanRows provides a mock function with given fields: rows, result
func (_m *MockORM) ScanRows(rows *sql.Rows, result interface{}) error {
	ret := _m.Called(rows, result)

	var r0 error
	if rf, ok := ret.Get(0).(func(*sql.Rows, interface{}) error); ok {
		r0 = rf(rows, result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Select provides a mock function with given fields: query, args
func (_m *MockORM) Select(query interface{}, args ...interface{}) ORM {
	ret := _m.Called(query, args)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) ORM); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Table provides a mock function with given fields: name
func (_m *MockORM) Table(name string) ORM {
	ret := _m.Called(name)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(string) ORM); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Update provides a mock function with given fields: attrs
func (_m *MockORM) Update(attrs ...interface{}) ORM {
	ret := _m.Called(attrs)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(...interface{}) ORM); ok {
		r0 = rf(attrs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// UpdateColumn provides a mock function with given fields: attrs
func (_m *MockORM) UpdateColumn(attrs ...interface{}) ORM {
	ret := _m.Called(attrs)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(...interface{}) ORM); ok {
		r0 = rf(attrs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// UpdateColumns provides a mock function with given fields: values
func (_m *MockORM) UpdateColumns(values interface{}) ORM {
	ret := _m.Called(values)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(interface{}) ORM); ok {
		r0 = rf(values)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Updates provides a mock function with given fields: values, ignoreProtectedAttrs
func (_m *MockORM) Updates(values interface{}, ignoreProtectedAttrs ...bool) ORM {
	ret := _m.Called(values, ignoreProtectedAttrs)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(interface{}, ...bool) ORM); ok {
		r0 = rf(values, ignoreProtectedAttrs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

// Where provides a mock function with given fields: query, args
func (_m *MockORM) Where(query interface{}, args ...interface{}) ORM {
	ret := _m.Called(query, args)

	var r0 ORM
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) ORM); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ORM)
		}
	}

	return r0
}

type MockORMSetsUser struct {
	*MockORM
	Email    string
	Password string
	ID       uint64
}

func (m *MockORMSetsUser) First(out interface{}, where ...interface{}) ORM {
	ps := reflect.ValueOf(out)
	s := ps.Elem()

	email := s.FieldByName("Email")
	email.SetString(m.Email)

	password := s.FieldByName("Password")
	password.SetString(m.Password)

	id := s.FieldByName("ID")
	id.SetUint(m.ID)

	return m
}

type MockORMCreatesUser struct {
	*MockORM
	Name     string
	Email    string
	Password string
	ID       uint64
}

func (m *MockORMCreatesUser) Create(value interface{}) ORM {
	ps := reflect.ValueOf(value)
	s := ps.Elem()

	name := s.FieldByName("Name")
	name.SetString(m.Name)

	email := s.FieldByName("Email")
	email.SetString(m.Email)

	password := s.FieldByName("Password")
	password.SetString(m.Password)

	id := s.FieldByName("ID")
	id.SetUint(m.ID)

	return m
}
