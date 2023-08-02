package main

import (
	"fmt"
	"jobtracker/dao"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var jobDao *dao.JobDAO

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	jobDao = dao.NewJobDAO()
}
func main() {

	r := gin.Default()
	port := os.Getenv("PORT")

	r.GET(`/job/:id`, func(ctx *gin.Context) {

		id, _ := strconv.Atoi(ctx.Param("id"))

		job := jobDao.Read(id)
		fmt.Println(job)
		ctx.JSON(200, job)

	})

	r.Run(":" + port)
}
