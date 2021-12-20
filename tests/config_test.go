package tests_test

import (
	"orango.jhink.com/orango"
)

// Configure to start testing
var (
	TestCollection = "TestCollection"
	TestDoc        DocTest
	TestDbName     = "orangodb"
	TestUsername   = "root"
	TestPassword   = "orango-db"
	TestString     = "test string"
	verbose        = false
	TestServer     = "http://localhost:8560"
)

// document to test
type DocTest struct {
	orango.Document // arango Document to save id, key, rev
	Text            string
}
