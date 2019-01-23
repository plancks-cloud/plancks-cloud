package mem

import "github.com/hashicorp/go-memdb"

var (
	// Create the DB schema
	schema = &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"services": {
				Name: "services",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
				},
			},
			"routes": {
				Name: "routes",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
				},
			},
		},
	}
)
