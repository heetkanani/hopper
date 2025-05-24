package main

import (
	"fmt"
	"log"

	"github.com/heetkanani/hopper/hopper"
)

func main() {
	db, err := hopper.New()
	if err != nil {
		log.Fatal(err)
	}
	user := map[string]string{
		"name": "Heet",
		"age":  "26",
	}
	id, err := db.Insert("users", user)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", id)

}
