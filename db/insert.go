package db

import (
	"context"
	"log"
)

const (
	SessionsCollection = "sessions"
	UsersCollection    = "users"
	NKOCollection      = "nko"
	ProjectsCollection = "projects"
)

func Insert(collection string, document interface{}) (err error) {
	_, err = db.Collection(collection).InsertOne(context.Background(), document)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
