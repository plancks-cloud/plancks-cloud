package mem

import "github.com/hashicorp/go-memdb"

var (
	// Create the DB schema
	schema = &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"Service": {
				Name: "Service",
				Indexes: map[string]*memdb.IndexSchema{
					"Name": {
						Name:    "Name",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
					"contractId": {
						Name:    "contractId",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "ContractID"},
					},
				},
			},
			"Wallet": {
				Name: "Wallet",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
				},
			},
			"CancelContract": {
				Name: "CancelContract",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"contractId": {
						Name:    "contractId",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "ContractID"},
					},
				},
			},
		},
	}
)
