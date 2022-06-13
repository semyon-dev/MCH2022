package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
	"mch2022/cErrors"
	"mch2022/config"
	"mch2022/db"
	"mch2022/model"
	"time"
)

func CreateUser(user model.UserWithCredentials) (newUser model.UserWithCredentials, err error) {

	if user.Email != "" {
		_, exist := db.GetUserWithCredentials("email", user.Email)
		if exist {
			return user, cErrors.EmailAlreadyUsed
		}
	}

	var hashedPassword = []byte{}
	if user.Password != "" {
		hashedPassword, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
			return user, cErrors.InternalServerError
		}
	}

	user.Id = primitive.NewObjectID()
	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	if user.Tags == nil {
		user.Tags = []string{}
	}
	user.RegisteredProjects = []primitive.ObjectID{}
	if user.PhotoURL == "" {
		user.PhotoURL = config.DefaultUserPicture
	}

	err = db.Insert(db.UsersCollection, user)
	if err != nil {
		log.Println(err)
		return user, cErrors.InternalServerError
	}

	return user, err
}
