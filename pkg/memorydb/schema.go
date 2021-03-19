package memorydb

import (
	"github.com/hashicorp/go-memdb"
)

func CreateSchema(tableName string, mappingKey string, indexName string) memdb.DBSchema {
	return  memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			tableName: &memdb.TableSchema{
				Name: tableName,
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    indexName,
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: mappingKey},
					},
				},
			},
		},
	}
}
