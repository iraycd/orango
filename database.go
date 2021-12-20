package orango

import (
	"errors"
	"regexp"
	"time"

	driver "github.com/arangodb/go-driver"
)

// Database struct
type Database struct {
	Name        string `json:"name"`
	Id          string `json:"id"`
	Path        string `json:"path"`
	System      bool   `json:"isSystem"`
	Collections []Collection
	baseURL     string
	session     *Session
	driverDB    driver.Database
}

// Execute AQL query into server and returns cursor struct
func (d *Database) Execute(q *Query) (*Cursor, error) {
	var cursor Cursor
	t0 := time.Now()
	driverCursor, err := d.driverDB.Query(nil, q.Aql, q.BindVars)
	cursor.driverCursor = driverCursor
	t1 := time.Now()
	cursor.max = int(cursor.driverCursor.Count())
	cursor.Time = t1.Sub(t0)
	return &cursor, err
}

// Col returns Collection attached to current Database
func (db Database) Col(name string) *Collection {
	var col Collection
	driverCol, err := db.driverDB.Collection(nil, name)
	if err != nil {
		if db.session.safe {
			panic("Collection " + name + " not found")
		} else {
			var col CollectionOptions
			col.Name = name
			db.CreateCollection(&col)
			return db.Col(name)
		}
	}
	col.db = &db
	col.Name = driverCol.Name()
	col.driverCol = driverCol
	return &col
}

func validColName(name string) error {
	reg, err := regexp.Compile(`^[A-z]+[0-9\-_]*`)

	if err != nil {
		return err
	}
	if !reg.MatchString(name) {
		return errors.New("Invalid collection name")
	}

	return nil
}

// Create collections
func (d *Database) CreateCollection(c *CollectionOptions) error {

	err := validColName(c.Name)
	if err != nil {
		return err
	}

	_, err = d.driverDB.CreateCollection(nil, c.Name, nil)
	if err != nil {

		return errors.New("Failed to create collection")
	}
	// Collections(d)
	return nil
}
