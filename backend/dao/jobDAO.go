package dao

import (
	"jobtracker/models"
	"jobtracker/utils"
)

type JobDAO struct {
}

var db, _ = utils.GetGormConnection()

func (j *JobDAO) Create(job models.Job) (bool, error) {
	res := db.Create(&job)
	if res.Error != nil {
		return false, res.Error
	}

	return true, nil
}

func NewJobDAO() *JobDAO {

	return &JobDAO{}
}
