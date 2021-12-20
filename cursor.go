package orango

import (
	"time"

	"github.com/arangodb/go-driver"
)

type Cursor struct {
	db *Database `json:"-"`
	Id string    `json:"Id"`

	Index  int           `json:"-"`
	Result []interface{} `json:"result"`
	More   bool          `json:"hasMore"`
	Amount int           `json:"count"`

	Err          bool   `json:"error"`
	ErrMsg       string `json:"errorMessage"`
	Code         int    `json:"code"`
	max          int
	Time         time.Duration `json:"time"`
	driverCursor driver.Cursor
}

func NewCursor(db *Database) *Cursor {
	var c Cursor
	if db == nil {
		return nil
	}
	c.db = db
	return &c
}

// FetchOne iterates over cursor, returns false when no more values into batch, fetch next batch if necesary.
func (c *Cursor) FetchOne(r interface{}) bool {
	for {
		_, err := c.driverCursor.ReadDocument(nil, &r)
		if err != nil {
			return false
			// handle other errors
		}
		return true
	}
	return false
}
