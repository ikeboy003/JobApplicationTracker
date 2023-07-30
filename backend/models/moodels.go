package models

type Job struct {
	JobID      int    `json:"jobID"`
	JobName    string `json:"jobName"`
	JobCompany string `json:"jobCompany"`
	AppStatus  string `json:"appStatus"`
}
