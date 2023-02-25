package config

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

func GetSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	fmt.Println("Database is connected successfully")
	return s
}
