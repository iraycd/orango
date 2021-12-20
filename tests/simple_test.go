package tests_test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"orango.jhink.com/orango"
)

func TestSimple(t *testing.T) {
	// connect
	session, err := orango.Connect(TestServer, TestUsername, TestPassword)
	assert.Nil(t, err)
	log.Println("Client")
	log.Println(session.Client)
	defer session.DropDB(TestDbName)
	// db := session.DB("_system")
	dbs := session.AvailableDBs()
	log.Println(dbs)

	// log.Print("Session")
	// log.Print(s)
	// // Create the db
	session.CreateDB(TestDbName)

	db := session.DB(TestDbName)
	log.Println("session.DB")
	log.Println("db.Name()")
	log.Println(db.Name)
	assert.NotNil(t, db.Name)

	c := db.Col(TestCollection)
	assert.NotNil(t, c.Name)

	// Save
	var saveTestDoc DocTest
	saveTestDoc.Text = TestString
	meta, err := c.Save(saveTestDoc)
	assert.Nil(t, err)
	log.Println("meta.ID")
	log.Println(meta.ID)
	assert.NotNil(t, meta.ID)

	saveTestDoc.Text = "Hello"
	meta, err = c.Replace(meta.Key, saveTestDoc)
	assert.Nil(t, err)
	log.Println("meta.ID")
	log.Println(meta.ID)
	assert.NotNil(t, meta.ID)

	q := orango.NewQuery("FOR i in TestCollection RETURN i")
	curr, err := db.Execute(q)
	if err != nil {
		panic(err)
	}
	var doc DocTest
	for curr.FetchOne(&doc) {
		assert.NotNil(t, doc.Text)
	}

}
