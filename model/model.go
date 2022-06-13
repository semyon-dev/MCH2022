package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Session struct {
	RefreshToken string             `json:"-" bson:"refreshToken"`
	ExpiresIn    primitive.DateTime `json:"-" bson:"expiresIn"`
	Token        string             `json:"token" bson:"-"`
	UserAgent    string             `json:"userAgent" bson:"userAgent"`
	IP           string             `json:"-" bson:"ip"`
	UserID       primitive.ObjectID `json:"-" bson:"userID"`
}

type Contact struct {
	Type  string `json:"type" bson:"type"` // (telegram, email, other)
	Value string `json:"value" bson:"value"`
}

type User struct {
	ID            primitive.ObjectID `json:"ID" bson:"_id"`
	Password      string             `json:"-" bson:"password"`
	Contact       Contact            `json:"contact" bson:"contact"`
	Locale        string             `json:"locale" bson:"locale"`
	CreatedAt     time.Time          `json:"-" bson:"createdAt"`
	Status        string             `json:"status" bson:"status"`
	Tags          []string           `json:"tags" bson:"tags"`
	Education     string             `json:"education" bson:"education"`
	Age           uint               `json:"age" bson:"age"`
	City          string             `json:"city" bson:"city"`
	Name          string             `json:"name" bson:"name"`
	PhotoURL      string             `json:"photoURL" bson:"photoURL"`
	Description   string             `json:"description" bson:"description"`
	Email         string             `json:"email" bson:"email"`
	EmailVerified bool               `json:"emailVerified" bson:"emailVerified"`
	//RegisteredProjects []primitive.ObjectID `json:"registeredProjects" bson:"registeredProjects"`
	Rating float64 `json:"rating" bson:"rating"`
	Rates  []Rate  `json:"rates" bson:"rates"`
}

type UserWithCredentials struct {
	Id            primitive.ObjectID `json:"ID" bson:"_id"`
	Password      string             `json:"password" bson:"password"`
	Contact       Contact            `json:"contact" bson:"contact"`
	Locale        string             `json:"locale" bson:"locale"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`
	Status        string             `json:"status" bson:"status"`
	Tags          []string           `json:"tags" bson:"tags"`
	Education     string             `json:"education" bson:"education"`
	Age           uint               `json:"age" bson:"age"`
	City          string             `json:"city" bson:"city"`
	FIO           string             `json:"FIO" bson:"FIO"`
	PhotoURL      string             `json:"photoURL" bson:"photoURL"`
	Description   string             `json:"description" bson:"description"`
	Email         string             `json:"email" bson:"email"`
	EmailVerified bool               `json:"emailVerified" bson:"emailVerified"`
	//RegisteredProjects []primitive.ObjectID `json:"registeredProjects" bson:"registeredProjects"`
}

type Project struct {
	Id                   primitive.ObjectID     `json:"ID" bson:"_id"`
	AuthorId             primitive.ObjectID     `json:"authorID" bson:"authorID"`
	PhotoURL             string                 `json:"photoURL" bson:"photoURL"`
	Name                 string                 `json:"name" bson:"name"`
	Type                 string                 `json:"type" bson:"type"`
	ShortDescription     string                 `json:"shortDescription" bson:"shortDescription"`
	Description          string                 `json:"description" bson:"description"`
	Location             string                 `json:"location" bson:"location"`
	Skills               []string               `json:"skills" bson:"skills"`
	AgeStart             int                    `json:"ageStart" bson:"ageStart"`
	AgeEnd               int                    `json:"ageEnd" bson:"ageEnd"`
	Direction            string                 `json:"direction" bson:"direction"` // направление
	Website              string                 `json:"website" bson:"website"`
	Participation        string                 `json:"participation" bson:"participation"` // Способ участия - офлайн/онлайн
	Tags                 []string               `json:"tags" bson:"tags"`
	Roles                []string               `json:"roles" bson:"roles"`
	Requirements         []string               `json:"requirements" bson:"requirements"`
	Services             []string               `json:"services" bson:"services"`
	Deadline             primitive.DateTime     `json:"deadline" bson:"deadline"` // дедлайн подачи заявки
	TimeEnd              primitive.DateTime     `json:"timeEnd" bson:"timeEnd"`
	TimeStart            primitive.DateTime     `json:"timeStart" bson:"timeStart"`
	CreatedAt            time.Time              `json:"-" bson:"createdAt"`
	Participants         []primitive.ObjectID   `json:"participants" bson:"participants"`
	AcceptedParticipants []primitive.ObjectID   `json:"acceptedParticipants" bson:"acceptedParticipants"`
	BannedParticipants   []primitive.ObjectID   `json:"bannedParticipants" bson:"bannedParticipants"`
	CustomFieldsInput    []CustomField          `json:"customFields" bson:"-"`
	CustomFields         map[string]CustomField `json:"customFieldsMap" bson:"customFields"`
	ParticipantsCount    int                    `json:"participantsCount" bson:"-"`
}

type Rate struct {
	Id       primitive.ObjectID `json:"ID" bson:"_id"`
	AuthorId primitive.ObjectID `json:"authorID" bson:"authorID"`
	Value    string             `json:"value" bson:"value"`
	Rate     float64            `json:"rate" bson:"rate"`
}

type CustomField struct {
	Id          string      `json:"ID" bson:"-"`
	Name        string      `json:"name" bson:"name"`
	Description string      `json:"description" bson:"description"`
	Type        string      `json:"type" bson:"type"`
	Payload     bson.M      `json:"payload" bson:"payload"`
	Reply       interface{} `json:"data,omitempty" bson:",omitempty"`
}

type NKO struct {
	PhotoURL           string               `json:"photoURL" bson:"photoURL"`
	Id                 primitive.ObjectID   `json:"ID" bson:"_id"`
	Email              string               `json:"email" bson:"email"`
	EmailVerified      bool                 `json:"emailVerified" bson:"emailVerified"`
	Name               string               `json:"name" bson:"name"`
	Location           string               `json:"location" bson:"location"`
	Time               primitive.DateTime   `json:"time" bson:"time"`
	ShortDescription   string               `json:"shortDescription" bson:"shortDescription"`
	Description        string               `json:"description" bson:"description"`
	Tags               []string             `json:"tags" bson:"tags"`
	INN                string               `json:"INN" bson:"INN"`
	Password           string               `json:"password" bson:"password"`
	Contact            Contact              `json:"contact" bson:"contact"`
	RegisteredProjects []primitive.ObjectID `json:"registeredProjects" bson:"registeredProjects"`
	Rating             float64              `json:"rating" bson:"rating"`
	Rates              []Rate               `json:"rates" bson:"rates"`
}

type NKOReview struct {
	Id       primitive.ObjectID `json:"ID" bson:"_id"`
	Comment  string             `json:"comment" bson:"comment"`
	Rating   int                `json:"rating" bson:"rating"` // от 1 до 5
	AuthorId primitive.ObjectID `json:"authorID" bson:"authorID"`
	NkoID    primitive.ObjectID `json:"nkoID" bson:"nkoID"`
}

type ProjectReview struct {
	Id        primitive.ObjectID `json:"ID" bson:"_id"`
	ProjectID primitive.ObjectID `json:"projectID" bson:"projectID"`
	Comment   string             `json:"comment" bson:"comment"`
	Rating    int                `json:"rating" bson:"rating"` // от 1 до 5
	AuthorId  primitive.ObjectID `json:"authorID" bson:"authorID"`
}

type UserRatingListItem struct {
	Place    int                `json:"place"`
	Name     string             `json:"name"`
	UserID   primitive.ObjectID `json:"userID"`
	Rating   float64            `json:"rating"`
	IsClient bool               `json:"isClient"`
}
