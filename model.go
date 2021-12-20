package orango

import (
	"errors"
)

// Context to share state between hook and track transaction state
type Context struct {
	Keys map[string]interface{}
	Db   *Database
}

func NewContext(db *Database) (*Context, error) {
	if db == nil {
		return nil, errors.New("Invalid DB")
	}
	var c Context
	c.Db = db
	c.Keys = make(map[string]interface{})
	return &c, nil
}
