package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"mch2022/model"
)

var saveProjection interface{} = bson.M{"password": 0, "thirdPartyAuth": 0, "session": 0}

func GetProjectByID(id primitive.ObjectID) (project model.Project, isExist bool) {
	filter := bson.M{"_id": id}
	err := db.Collection(ProjectsCollection).FindOne(context.Background(), filter).Decode(&project)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Project{}, false
		}
		log.Println(err)
		return
	}
	return project, true
}

func SearchProjects(query string) (projects []model.Project) {
	filter := bson.M{"name": bson.M{"$regex": query}}
	cursor, err := db.Collection(ProjectsCollection).Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	if err = cursor.All(context.Background(), &projects); err != nil {
		log.Println(err)
	}
	return projects
}

func GetProjects() (projects []model.Project) {
	cursor, err := db.Collection(ProjectsCollection).Find(context.Background(), bson.M{})
	if err != nil {
		log.Println(err)
	}
	if err = cursor.All(context.Background(), &projects); err != nil {
		log.Println(err)
	}
	return
}

func GetProjectsByFilters(ageStart, ageEnd int, participation, location, direction, name, searchQuery string, tags, roles, skills []string, authorID primitive.ObjectID) (projects []model.Project) {
	var filters = []bson.M{}
	if participation != "" {
		filters = append(filters, bson.M{"participation": participation})
	}
	if location != "" {
		filters = append(filters, bson.M{"location": location})
	}
	if direction != "" {
		filters = append(filters, bson.M{"direction": direction})
	}
	if !authorID.IsZero() {
		filters = append(filters, bson.M{"authorID": authorID})
	}
	if len(tags) != 0 {
		filters = append(filters, bson.M{"tags": bson.M{"$in": tags}})
	}
	if len(roles) != 0 {
		filters = append(filters, bson.M{"roles": bson.M{"$in": roles}})
	}
	if len(skills) != 0 {
		filters = append(filters, bson.M{"skills": bson.M{"$in": skills}})
	}
	if name != "" {
		filters = append(filters, bson.M{"$text": bson.M{"$search": name}})
	}
	if searchQuery != "" {
		filters = append(filters, bson.M{"name": bson.M{"$regex": searchQuery, "$options": "$i"}})
	}
	if ageStart != 0 && ageEnd != 0 {
		filters = append(filters, bson.M{"$or": bson.A{
			bson.M{"ageStart": bson.M{"$gte": ageStart, "$lte": ageEnd}},
			bson.M{"ageEnd": bson.M{"$gte": ageStart, "$lte": ageEnd}},
		}})
	}
	filter := bson.M{"$and": filters}
	if len(filters) == 0 {
		filter = bson.M{}
	}
	o := options.Find().SetBatchSize(100).SetAllowDiskUse(true).SetAllowPartialResults(false)
	cursor, err := db.Collection(ProjectsCollection).Find(context.Background(), filter, o)
	if err != nil {
		log.Println(err)
	}
	if err = cursor.All(context.Background(), &projects); err != nil {
		log.Println(err)
	}
	return
}

func GetNKOsByFilters(searchQuery string, tags []string) (nkos []model.NKO) {
	var filters = []bson.M{}
	if len(tags) != 0 {
		filters = append(filters, bson.M{"tags": bson.M{"$in": tags}})
	}
	if searchQuery != "" {
		filters = append(filters, bson.M{"name": bson.M{"$regex": searchQuery, "$options": "$i"}})
	}
	filter := bson.M{"$and": filters}
	if len(filters) == 0 {
		filter = bson.M{}
	}
	o := options.Find().SetBatchSize(100).SetAllowDiskUse(true).SetAllowPartialResults(false)
	cursor, err := db.Collection(NKOCollection).Find(context.Background(), filter, o)
	if err != nil {
		log.Println(err)
	}
	if err = cursor.All(context.Background(), &nkos); err != nil {
		log.Println(err)
	}
	return
}

func GetUsers() (users []model.User) {
	cursor, err := db.Collection(UsersCollection).Find(context.Background(), bson.M{})
	if err != nil {
		log.Println(err)
	}
	if err = cursor.All(context.Background(), &users); err != nil {
		log.Println(err)
	}
	return users
}

func GetProjectsByAuthorId(authorId primitive.ObjectID) (Projects []model.Project) {
	filter := bson.M{"authorID": authorId}
	cursor, err := db.Collection(ProjectsCollection).Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	if err = cursor.All(context.Background(), &Projects); err != nil {
		log.Println(err)
	}
	return
}

func GetUsersByIdsWithPagination(ids []primitive.ObjectID, limit, skip int64) (users []model.User) {
	filter := bson.M{"_id": bson.M{"$in": ids}}
	opts := options.Find()
	opts.SetLimit(limit)
	opts.SetSkip(skip)
	opts.SetProjection(bson.M{"registeredProjects": 0, "password": 0, "thirdPartyAuth": 0, "session": 0})
	cursor, err := db.Collection(UsersCollection).Find(context.Background(), filter, opts)
	if err != nil {
		log.Println(err)
	}
	if err = cursor.All(context.Background(), &users); err != nil {
		log.Println(err)
	}
	return
}

// GetUserWithCredentials without projection
func GetUserWithCredentials(field, value string) (user model.UserWithCredentials, isExist bool) {
	filter := bson.M{field: value}
	err := db.Collection(UsersCollection).FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.UserWithCredentials{}, false
		}
		log.Println(err)
		return
	}
	return user, true
}

func GetNKOByField(field, value string) (nko model.NKO, isExist bool) {
	filter := bson.M{field: value}
	err := db.Collection(NKOCollection).FindOne(context.Background(), filter).Decode(&nko)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.NKO{}, false
		}
		log.Println(err)
		return
	}
	return nko, true
}

func GetNKO(id primitive.ObjectID) (nko model.NKO, isExist bool) {
	filter := bson.M{"_id": id}
	err := db.Collection(NKOCollection).FindOne(context.Background(), filter).Decode(&nko)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.NKO{}, false
		}
		log.Println(err)
		return
	}
	return nko, true
}

func GetUserById(id primitive.ObjectID) (user model.User, isExist bool) {
	filter := bson.M{"_id": id}
	opts := options.FindOne()
	opts.SetProjection(saveProjection)
	err := db.Collection(UsersCollection).FindOne(context.Background(), filter, opts).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, false
		}
		log.Println(err)
		return
	}
	return user, true
}

func CountRating(rates []model.Rate) (result float64) {
	for _, rate := range rates {
		result += rate.Rate / float64(len(rates))
	}
	return result
}
