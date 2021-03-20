package memorydb

import "github.com/hashicorp/go-memdb"

// Create a sample struct
type MemDB struct {
	DB        *memdb.MemDB
	TableName string
}

func New(schema *memdb.DBSchema, tableName string) (*MemDB, error) {
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		return nil, err
	}
	return &MemDB{
		DB:        db,
		TableName: tableName,
	}, nil
}

func (m *MemDB) Save(content interface{}) error {
	txn := m.DB.Txn(true)

	if err := txn.Insert(m.TableName, content); err != nil {
		return err
	}
	txn.Commit()
	return nil
}

func (m *MemDB) FetchById(index string, args interface{}) (interface{}, error) {
	txn := m.DB.Txn(false)
	defer txn.Abort()
	result, err := txn.First(m.TableName, index, args)
	if err != nil {
		return nil, err
	}
	txn.Commit()
	return result, nil
}

type Store interface {
	Save(content interface{}) error
	FetchById(index string, args interface{}) (interface{}, error)
}
