package vectorstore

import (
	"github.com/aldarisbm/memory/pkg/types"
	"github.com/google/uuid"
)

// VectorStorer is an interface for vector stores
type VectorStorer interface {
	// StoreVector stores the given Document
	StoreVector(document *types.Document) error
	// QuerySimilarity returns the k most similar documents to the given vector
	QuerySimilarity(vector []float32, k int64) ([]uuid.UUID, error)
}
