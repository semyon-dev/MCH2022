package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mch2022/cErrors"
	"mch2022/config"
	"mch2022/db"
	"mch2022/model"
	"time"
)

func CreateProject(project model.Project) (id primitive.ObjectID, err error) {

	if project.AuthorId.IsZero() {
		return id, cErrors.InvalidAuthorID
	}

	if project.PhotoURL == "" {
		project.PhotoURL = config.DefaultProjectPicture
	}

	if project.Participants == nil {
		project.Participants = []primitive.ObjectID{}
	}

	if project.Skills == nil {
		project.Skills = []string{}
	}

	if project.Roles == nil {
		project.Skills = []string{}
	}

	if project.Services == nil {
		project.Services = []string{}
	}

	if project.Requirements == nil {
		project.Requirements = []string{}
	}

	if project.BannedParticipants == nil {
		project.Participants = []primitive.ObjectID{}
	}

	if project.CustomFieldsInput == nil {
		project.CustomFields = map[string]model.CustomField{}
	}

	if project.CustomFields == nil {
		project.CustomFields = map[string]model.CustomField{}
	}

	for _, v := range project.CustomFieldsInput {
		project.CustomFields[primitive.NewObjectID().Hex()] = v
	}

	project.CreatedAt = time.Now()
	project.Id = primitive.NewObjectID()

	err = db.Insert(db.ProjectsCollection, project)
	if err != nil {
		return project.Id, cErrors.InternalServerError
	}
	return project.Id, nil
}
