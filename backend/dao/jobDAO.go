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
func (j *JobDAO) Read(jobId int) (job models.Job) {

	db.Where("job_id= ?", jobId).First(&job)
	return job
}
func NewJobDAO() *JobDAO {

	return &JobDAO{}
}
