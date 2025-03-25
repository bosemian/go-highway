package main

import (
	"13_defer/db"
	"fmt"
	"log"
	"os"
)

func manageDB(d db.Database) {
	d.Connect()
	// _, err := d.Query()
	// if err != nil {
	// 	// if has an error return
	// 	// but d.Disconnect() not execute
	// 	return
	// }
	defer d.Disconnect()
}

func main() {
	// defer fmt.Println("I will run last!")
	// fmt.Println("Hello, World")

	// stack()
	// fileHandling()

	mysql := db.MySQL{DSN: "root:password@tcp(127.0.0.1:3306)/appdb"}
	pg := db.Postgres{URL: "postgres://user:pass@localhost:5432/appdb"}

	manageDB(mysql)
	manageDB(pg)
}

func stack() {
	defer fmt.Println("first")
	defer fmt.Println("second")
	defer fmt.Println("third")
	fmt.Println("main function end")
}

func fileHandling() {
	f, err := os.Open("file.txt")
	if err != nil {
		// defer f.Close()
		// log.Fatal(err)
		log.Println("Error opening file:", err)
		return
	}
	defer f.Close()
	fmt.Println(f.Name())
}
