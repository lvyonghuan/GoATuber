package sqlite

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/aldarisbm/memory/pkg/datasource"
	"github.com/aldarisbm/memory/pkg/types"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"os/user"
)

const DomainName = "xyz.memorystore"

type localStorer struct {
	db   *sql.DB
	path string
}

// NewLocalStorer returns a new local storer
// if path is empty, it will default to $HOME/memory/memory.db
func NewLocalStorer(opts ...CallOptions) *localStorer {
	o := applyCallOptions(opts)
	if o.path == "" {
		usr, _ := user.Current()
		dir := usr.HomeDir
		_ = os.Mkdir(fmt.Sprintf("%s/%s", dir, DomainName), os.ModePerm)
		o.path = fmt.Sprintf("%s/%s/memory.db", dir, DomainName)
	}
	db, err := createTable(o.path)
	if err != nil {
		log.Fatal(err)
	}

	ls := &localStorer{
		db:   db,
		path: o.path,
	}
	return ls
}

// Close closes the local storer
func (l *localStorer) Close() error {
	return l.db.Close()
}

// GetDocument returns the document with the given id
func (l *localStorer) GetDocument(id uuid.UUID) (*types.Document, error) {
	var doc types.Document

	stmt, err := l.db.Prepare("SELECT id, user, text, created_at, last_read_at, vector, metadata FROM documents WHERE id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	var metadataBytes, vectorBytes []byte
	err = row.Scan(&doc.ID, &doc.User, &doc.Text, &doc.CreatedAt, &doc.LastReadAt, &vectorBytes, &metadataBytes)
	if err != nil {
		return nil, fmt.Errorf("scanning document: %s", err)
	}
	if err := json.Unmarshal(metadataBytes, &doc.Metadata); err != nil {
		return nil, fmt.Errorf("unmarshaling metadata: %s", err)
	}
	if err := json.Unmarshal(vectorBytes, &doc.Vector); err != nil {
		return nil, fmt.Errorf("unmarshaling vector: %s", err)
	}

	return &doc, nil
}

// GetDocuments returns the documents with the given ids
func (l *localStorer) GetDocuments(ids []uuid.UUID) ([]*types.Document, error) {
	// TODO should probably do this in a single query
	var docs []*types.Document

	for _, id := range ids {
		doc, err := l.GetDocument(id)
		if err != nil {
			return nil, err
		}
		docs = append(docs, doc)
	}
	return docs, nil
}

// StoreDocument stores the given document
// We marshal the vector and metadata to []byte and store them as blobs
func (l *localStorer) StoreDocument(doc *types.Document) error {
	stmt, err := l.db.Prepare("INSERT INTO documents (id, user,  text, created_at, last_read_at, vector, metadata) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("preparing statement: %s", err)
	}

	vector, err := json.Marshal(doc.Vector)
	if err != nil {
		return fmt.Errorf("marshaling vector: %s", err)
	}
	metadata, err := json.Marshal(doc.Metadata)
	if err != nil {
		return fmt.Errorf("marshaling metadata: %s", err)
	}
	res, err := stmt.Exec(doc.ID, doc.User, doc.Text, doc.CreatedAt, doc.LastReadAt, vector, metadata)
	if err != nil {
		return fmt.Errorf("inserting document: %s", err)
	}
	_, err = res.LastInsertId()
	if err != nil {
		return fmt.Errorf("getting last insert id: %s", err)
	}
	return nil
}

// Ensure localStorer implements DataSourcer
var _ datasource.DataSourcer = (*localStorer)(nil)
