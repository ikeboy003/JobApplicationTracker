package models

type Job struct {
	JobID      int    `gorm:"column:job_id;primaryKey" json:"jobID"`
	JobName    string `gorm:"column:job_name" json:"jobName"`
	JobCompany string `gorm:"column:job_company" json:"jobCompany"`
	AppStatus  string `gorm:"column:app_status" json:"appStatus"`
}

type Dao[T any] interface {
	Create(T) (bool, error)
	Read(int) (T, error)
	Update(T) (bool, error)
	Delete(T) (bool, error)
}

func (Job) TableName() string {
	return "jobsapplied"
}
