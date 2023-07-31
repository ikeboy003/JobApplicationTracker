package main

import (
	"fmt"
	"jobtracker/dao"
	"jobtracker/models"
)

func main() {

	newJob := models.Job{
		JobName:    "Fake Job",
		JobCompany: "Fake Company",
		AppStatus:  "applied",
	}
	jdao := dao.NewJobDAO()

	res, _ := jdao.Create(newJob)
	if res {
		fmt.Println("Job Created")
	}

}
