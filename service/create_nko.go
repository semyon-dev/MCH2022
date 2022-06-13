package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
	"mch2022/cErrors"
	"mch2022/config"
	"mch2022/db"
	"mch2022/model"
)

func CreateNKO(nko model.NKO) (model.NKO, error) {

	var err error
	var hashedPassword = []byte{}
	if nko.Password != "" {
		hashedPassword, err = bcrypt.GenerateFromPassword([]byte(nko.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
			return nko, cErrors.InternalServerError
		}
	}

	nko.Id = primitive.NewObjectID()
	nko.Password = string(hashedPassword)
	if nko.Tags == nil {
		nko.Tags = []string{}
	}
	nko.RegisteredProjects = []primitive.ObjectID{}
	if nko.PhotoURL == "" {
		nko.PhotoURL = config.DefaultUserPicture
	}

	err = db.Insert(db.NKOCollection, nko)
	if err != nil {
		log.Println(err)
		return nko, cErrors.InternalServerError
	}

	return nko, err
}
