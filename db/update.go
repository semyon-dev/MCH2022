package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"mch2022/model"
)

func AddUserToProject(projectID, userID primitive.ObjectID) {
	filter := bson.M{"_id": projectID}
	update := bson.M{"$push": bson.M{"participants": userID}}
	_, err := db.Collection(ProjectsCollection).UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println(err)
	}
}

func DeleteUserFromProject(projectID, userID primitive.ObjectID) {
	filter := bson.M{"_id": projectID}
	update := bson.M{"$pull": bson.M{"participants": userID}}
	_, err := db.Collection(ProjectsCollection).UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println(err)
	}
}

func AddRateToUser(userID primitive.ObjectID, rate model.Rate) error {
	filter := bson.M{"_id": userID}
	update := bson.M{"$push": bson.M{"rates": rate}}
	_, err := db.Collection(UsersCollection).UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func AddRateToNKO(nkoID primitive.ObjectID, rate model.Rate) error {
	filter := bson.M{"_id": nkoID}
	update := bson.M{"$push": bson.M{"rates": rate}}
	_, err := db.Collection(NKOCollection).UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println(err)
	}
	return nil
}
