package memorydb

import "github.com/hashicorp/go-memdb"

// Create a sample struct
type MemDB struct {
	DB *memdb.MemDB
}
func New(schema *memdb.DBSchema) (*MemDB, error){
	db, err := memdb.NewMemDB(schema)
	if err != nil {
	 	return nil, err
	}
	return &MemDB{
		DB: db,
	}, nil
}

func (m *MemDB) Save(tableName string, content interface{}) error {
	txn := m.DB.Txn(true)

	if err := txn.Insert(tableName, content ); err != nil {
		return err
	}
	txn.Commit()
	return nil
}

func (m *MemDB) FetchById(tableName string, index string, args interface{}) (interface{}, error) {
	txn := m.DB.Txn(false)
	defer txn.Abort()
	result, err := txn.First(tableName, index, args)
	if err != nil {
		return nil, err
	}
	txn.Commit()
	return result, nil
}

type Store interface {
	Save(tableName string, content interface{}) error
	FetchById(tableName string, index string, args interface{}) (interface{}, error)
}
