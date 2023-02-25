package config

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

func GetSession() *mgo.Session {
	s, err := mgo.Dial({connectionLink})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database is connected successfully")
	return s
}
