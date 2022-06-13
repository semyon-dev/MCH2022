package service

import (
	_ "github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mch2022/db"
	"mch2022/model"
)

func GetMyProjects(userID primitive.ObjectID) (pr []model.Project, err error) {

	//user, ok := db.GetUserById(userID)
	//if !ok {
	//	return memberProjects, createdProjects, cErrors.InvalidUserID
	//}
	//
	//memberProjects = db.GetProjectsByIds(user.RegisteredProjects)
	//for memberIndex, _ := range memberProjects {
	//	memberProjects[memberIndex].ParticipantsCount = len(memberProjects[memberIndex].Participants) - 1
	//}
	//
	//if memberProjects == nil {
	//	memberProjects = []model.Project{}
	//}

	createdProjects := db.GetProjectsByAuthorId(userID)
	for i := range createdProjects {
		var cs = []model.CustomField{}
		for i, v := range createdProjects[i].CustomFields {
			v.Id = i
			cs = append(cs, v)
		}
		//createdProjects[i].CustomFieldsInput = cs
		createdProjects[i].ParticipantsCount = len(createdProjects[i].Participants)
	}
	if createdProjects == nil {
		createdProjects = []model.Project{}
	}

	return createdProjects, nil
}
