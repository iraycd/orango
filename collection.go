package orango

import (
	"errors"

	"github.com/arangodb/go-driver"
)

type CollectionOptions struct {
	Name        string                 `json:"name"`
	Type        uint                   `json:"type"`
	Sync        bool                   `json:"waitForSync,omitempty"`
	Compact     bool                   `json:"doCompact,omitempty"`
	JournalSize int                    `json:"journalSize,omitempty"`
	System      bool                   `json:"isSystem,omitempty"`
	Volatile    bool                   `json:"isVolatile,omitempty"`
	Keys        map[string]interface{} `json:"keyOptions,omitempty"`
	// Count
	Count int64 `json:"count"`
	// Cluster
	Shards    int      `json:"numberOfShards,omitempty"`
	ShardKeys []string `json:"shardKeys,omitempty"`
}

type Collection struct {
	db     *Database `json:"db"`
	Name   string    `json:"name"`
	System bool      `json:"isSystem"`
	Status int       `json:"status"`
	// 3 = Edges , 2 =  Documents
	Type      int    `json:"type"`
	policy    string `json:"-"`
	revision  bool   `json:"-"`
	driverCol driver.Collection
}

type DocumentMeta struct {
	driver.DocumentMeta
}

// Save saves doc into collection, doc should have Document Embedded to retrieve error and Key later.
func (col *Collection) Save(doc interface{}) (DocumentMeta, error) {
	var err error
	collectionProperties, err := col.driverCol.Properties(nil)
	if err != nil {
		return DocumentMeta{}, errors.New("Collection does not exist")
	}

	if collectionProperties.Type == driver.CollectionTypeDocument {
		meta, err := col.driverCol.CreateDocument(nil, doc)
		if err != nil {
			return DocumentMeta{}, errors.New("Invalid document json")
		}
		return DocumentMeta{meta}, nil
	} else {
		return DocumentMeta{}, errors.New("Trying to save doc into EdgeCollection")
	}
}

// Replace document
func (col *Collection) Replace(key string, doc interface{}) (DocumentMeta, error) {
	var err error

	collectionProperties, err := col.driverCol.Properties(nil)
	if err != nil {
		return DocumentMeta{}, errors.New("Collection does not exist")
	}

	if key == "" {
		return DocumentMeta{}, errors.New("Key must not be empty")
	}

	if collectionProperties.Type == driver.CollectionTypeDocument {
		meta, err := col.driverCol.UpdateDocument(nil, key, doc)
		if err != nil {
			return DocumentMeta{}, errors.New("Invalid document json")
		}
		return DocumentMeta{meta}, nil
	} else {
		meta, err := col.driverCol.UpdateDocument(nil, key, doc)
		if err != nil {
			return DocumentMeta{}, errors.New("Invalid document json")
		}
		return DocumentMeta{meta}, nil
	}
}

type Index struct {
	Id        string   `json:"id"`
	Type      string   `json:"type"`
	Unique    bool     `json:"unique"`
	MinLength int      `json:"minLength"`
	Fields    []string `json:"fields"`
	Size      int64    `json:"size"`
}
