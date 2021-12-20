package orango

import (
	"errors"
	"fmt"
	"log"
	"regexp"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	nap "github.com/diegogub/napping"
)

type Session struct {
	host   string
	safe   bool
	Client driver.Client
	nap    *nap.Session
	dbs    []driver.Database
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
	Extra    string `json:"extra"`
}

type Databases struct {
	List []string `json:"result" `
}

type auxCurrentDB struct {
	Db Database `json:"result"`
}

var Sess Session

// Connects to Database
func Connect(host, user, password string) (*Session, error) {
	var err error
	var conn driver.Connection
	var client driver.Client
	// default unsafe
	conn, err = http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{host},
		//Endpoints: []string{"https://5a812333269f.arangodb.cloud:8529/"},
	})
	if err != nil {
		return nil, errors.New(fmt.Sprint("Failed to create HTTP connection: %v", err))
	}

	log.Println(user, password)

	client, err = driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(user, password),
	})

	if err != nil {
		return nil, errors.New(fmt.Sprint("Invalid auth data to connect: %v", err))
	}
	Sess.host = host
	Sess.Client = client
	return &Sess, nil
}

// List available databases
func (s *Session) AvailableDBs() []string {
	var databaseList []driver.Database
	databaseList, err := s.Client.Databases(nil)
	if err != nil {
		errors.New(fmt.Sprint("Failed to open existing database: %v", err))
	}
	list := []string{}
	for _, element := range databaseList {
		list = append(list, element.Name())
	}
	return list
}

// // Create database
func (s *Session) CreateDB(name string) error {

	var err error
	// validate name
	reg, err := regexp.Compile(`^[A-z]+[0-9\-_]*`)

	if !reg.MatchString(name) {
		return errors.New("Invalid database name")
	}
	if err != nil {
		return err
	}
	var db_exists bool
	db_exists, err = s.Client.DatabaseExists(nil, name)
	if db_exists {
		return errors.New("Database with the specified name already exists")
	}

	var db driver.Database
	db, err = s.Client.CreateDatabase(nil, name, nil)
	log.Println("Created DB")
	log.Println(db.Name())
	if err != nil {
		return errors.New(fmt.Sprint("Failed to create database: %v", err))
	}
	return nil
}

// Drops database
func (s *Session) DropDB(name string) error {
	db := s.DB(name)
	err := db.driverDB.Remove(nil)
	return err
}

// DB returns database
func (s *Session) DB(name string) *Database {
	var db Database
	driverDB, err := s.Client.Database(nil, name)
	if err != nil {
		errors.New(fmt.Sprint("Failed to open existing database: %v", err))
	}
	db.Name = name
	db.session = s
	db.baseURL = s.host + "/_db/" + db.Name + "/_api/"
	db.driverDB = driverDB
	return &db
}

func (s *Session) Safe(safe bool) {
	s.safe = safe
	return
}
