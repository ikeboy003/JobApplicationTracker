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
func (j *JobDAO) Read(jobId int) models.Job {
	var job models.Job
	db.Where("job_id= ?", jobId).First(&job)
	return job
}
func NewJobDAO() *JobDAO {

	return &JobDAO{}
}

func (j *JobDAO) PerformTransaction(jobs []models.Job) error {

	tx := db.Begin()

	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, job := range jobs {
		if err := tx.Create(&job).Error; err != nil {
			tx.Rollback()
			return err

		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
