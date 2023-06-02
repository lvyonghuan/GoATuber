package embeddings

// Embedder is an interface for embedding text
type Embedder interface {
	// EmbedDocumentText returns the embedding of the given text
	EmbedDocumentText(text string) ([]float32, error)
	// EmbedDocumentTexts returns the embeddings of the given texts
	EmbedDocumentTexts(texts []string) ([][]float32, error)
}
