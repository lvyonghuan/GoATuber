package datasource

import (
	"github.com/aldarisbm/memory/pkg/types"
	"github.com/google/uuid"
)

// DataSourcer is an interface for data sources
type DataSourcer interface {
	// GetDocument returns the document with the given id
	GetDocument(id uuid.UUID) (*types.Document, error)
	// GetDocuments returns the documents with the given ids
	GetDocuments(ids []uuid.UUID) ([]*types.Document, error)
	// StoreDocument stores the given document
	StoreDocument(document *types.Document) error
	// Close closes the data source
	Close() error
}
