package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"jobtracker/dao"
	"jobtracker/models"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-contrib/cors"
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
	r.Use(cors.Default())
	port := os.Getenv("PORT")
	r.GET(`/job/:id`, handleReadingJob)
	r.GET(`/listAll`, GetAllJobs)
	r.POST("/upload", handleCSVPosting)
	r.POST("/newJob", handleJobPosting)
	r.Run(":" + port)

}
func handleJobPosting(ctx *gin.Context) {
	var jobData models.Job

	// Bind JSON data from request body to jobData
	if err := ctx.ShouldBindJSON(&jobData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if _, err := jobDao.Create(jobData); err != nil {
		ctx.JSON(400, gin.H{
			"error": "couldn't create",
		})
	} else {
		ctx.Status(http.StatusOK)
	}
}

func handleReadingJob(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	job := jobDao.Read(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch job"})
		return
	}

	ctx.JSON(http.StatusOK, job)
}

func handleCSVPosting(ctx *gin.Context) {

	file, err := ctx.FormFile("file")

	// The file cannot be received.
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}
	if !strings.HasSuffix(file.Filename, ".csv") {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "File is not a CSV"})
		return
	}
	uploadedFile, err := file.Open()
	persistFileInDB(uploadedFile, err, ctx)

}

func persistFileInDB(file multipart.File, err error, ctx *gin.Context) error {

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return err
	}
	defer file.Close()
	var jobs []models.Job
	reader := csv.NewReader(file)
	reader.Read()
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break // End of file
			}
			return err
		}
		jobName := record[0]
		jobCompany := record[1]
		status := strings.ToLower(record[2])

		job := models.Job{
			JobName:    jobName,
			JobCompany: jobCompany,
			AppStatus:  status,
		}
		jobs = append(jobs, job)

	}

	if err := jobDao.PerformTransaction(jobs); err != nil {
		return err
	}
	return nil
}

func GetAllJobs(ctx *gin.Context) {
	jobs := jobDao.GetAllJobs()
	ctx.JSON(200, jobs)
}
