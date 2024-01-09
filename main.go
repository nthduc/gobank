package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(store Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("new account ->", acc.Number)

	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "duc", "1234", "nthduc1234")
}

func main() {

	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		// seed stuff
		fmt.Println("seeding db")
		seedAccounts(store)

	}

	server := NewApiServer(":3000", store)
	server.Run()

}
