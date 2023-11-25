// main.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"net/http"
)

var session *gocql.Session

type Job struct {
	ID       gocql.UUID `json:"id"`
	Title    string     `json:"title"`
	Company  string     `json:"company"`
	Location string     `json:"location"`
}

func main() {
	var err error

	// Connect to Cassandra
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "recruitment"
	session, err = cluster.CreateSession()
	if err != nil {
		fmt.Println("Error connecting to Cassandra:", err)
		return
	}
	defer session.Close()

	// Set up Gin router
	router := gin.Default()

	// Define API routes
	router.POST("/jobs", createJob)
	router.GET("/jobs", getJobs)

	// Run the server
	router.Run(":8080")
}

func createJob(c *gin.Context) {
	var job Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	job.ID = gocql.TimeUUID()
	if err := session.Query(`
        INSERT INTO jobs (id, title, company, location) VALUES (?, ?, ?, ?)`,
		job.ID, job.Title, job.Company, job.Location).Exec(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, job)
}

func getJobs(c *gin.Context) {
	var jobs []Job
	iter := session.Query("SELECT id, title, company, location FROM jobs").Iter()

	var job Job
	for iter.Scan(&job.ID, &job.Title, &job.Company, &job.Location) {
		jobs = append(jobs, job)
	}

	if err := iter.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, jobs)
}
