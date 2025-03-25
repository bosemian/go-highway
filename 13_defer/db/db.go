package db

import (
	"errors"
	"fmt"
)

// Interface
type Database interface {
	Connect() error
	Disconnect() error
	Query() (int, error)
}

// MySQL struct
type MySQL struct {
	DSN string
}

func (m MySQL) Connect() error {
	fmt.Println("Connecting to MySQL with DSN:", m.DSN)
	return nil
}

func (m MySQL) Disconnect() error {
	fmt.Println("Disconnecting from MySQL")
	return nil
}

func (m MySQL) Query() (int, error) {
	fmt.Println("Query MySQL database...")
	// if true {
	// 	return 0, errors.New("Not found record")
	// }

	return 1, nil
}

// Postgres struct
type Postgres struct {
	URL string
}

func (p Postgres) Connect() error {
	fmt.Println("Connecting to Postgres with URL:", p.URL)
	return nil
}

func (p Postgres) Disconnect() error {
	fmt.Println("Disconnecting from Postgres")
	return nil
}

func (p Postgres) Query() (int, error) {
	fmt.Println("Query Postgres database...")
	if true {
		return 0, errors.New("Not found record")
	}

	return 1, nil
}
