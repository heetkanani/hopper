package main

import (
	"fmt"
	"log"

	"github.com/heetkanani/hopper/hopper"
)

func main() {
	user := map[string]string{
		"name": "Heet",
		"age":  "26",
	}
	_ = user

	db, err := hopper.New()
	if err != nil {
		log.Fatal(err)
	}
	coll, err := db.CreateColledtion("users")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", coll)

}
