package pc

import (
	"context"
	"fmt"
	"github.com/aldarisbm/memory/pkg/types"
	"github.com/aldarisbm/memory/pkg/vectorstore"
	"github.com/google/uuid"
	"github.com/nekomeowww/go-pinecone"
)

const Namespace = "asltm"

type storer struct {
	client    *pinecone.IndexClient
	namespace string
}

// NewStorer returns a new storer
func NewStorer(opts ...CallOptions) *storer {
	o := applyCallOptions(opts, options{
		namespace: Namespace,
	})
	// TODO we should be able to create a new index if it doesn't exist
	c, err := pinecone.NewIndexClient(
		pinecone.WithAPIKey(o.apiKey),
		pinecone.WithIndexName(o.indexName),
		pinecone.WithEnvironment(o.environment),
		pinecone.WithProjectName(o.projectName),
	)
	if err != nil {
		panic(err)
	}
	return &storer{
		client:    c,
		namespace: o.namespace,
	}
}

// StoreVector stores the given Document
// it attempts to save the metadata
func (p *storer) StoreVector(doc *types.Document) error {
	ctx := context.Background()
	req := pinecone.UpsertVectorsParams{
		Vectors: []*pinecone.Vector{
			{
				ID:       doc.ID.String(),
				Values:   doc.Vector,
				Metadata: doc.Metadata,
			},
		},
		Namespace: p.namespace,
	}

	resp, err := p.client.UpsertVectors(ctx, req)
	if err != nil {
		return fmt.Errorf("storing vector: %w", err)
	}
	if resp.UpsertedCount != 1 {
		return fmt.Errorf("storing vector: upserted count is not 1")
	}
	return nil
}

// QuerySimilarity returns the k most similar documents to the given vector
func (p *storer) QuerySimilarity(vector []float32, k int64) ([]uuid.UUID, error) {
	ctx := context.Background()
	req := pinecone.QueryParams{
		Vector:    vector,
		Namespace: p.namespace,
		TopK:      k,
	}
	resp, err := p.client.Query(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("querying vector: %w", err)
	}
	if len(resp.Matches) == 0 {
		// should we return an error here?
		return nil, nil
	}

	var uuids []uuid.UUID
	for _, match := range resp.Matches {
		id, err := uuid.Parse(match.ID)
		if err != nil {
			return nil, fmt.Errorf("querying vector: %w", err)
		}
		uuids = append(uuids, id)
	}
	return uuids, nil
}

// Ensure that storer implements VectorStorer
var _ vectorstore.VectorStorer = (*storer)(nil)
