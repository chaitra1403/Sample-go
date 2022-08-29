package main
//A standalone program is always in package main.

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// postDetail adds an Student details from JSON received in the request body.
func postDetail(c *gin.Context) {
    var newStudent studentDB

    // Call BindJSON to bind the received JSON to newStudent.
    if err := c.BindJSON(&newStudent); err != nil {
        return
    }

    // Add the new Student details to the slice.
    Student = append(Student, newStudent)
    c.IndentedJSON(http.StatusCreated, newStudent)
}

// getStudentByUSN locates the Student whose USN value matches the USN
//parameter sent by the client, then returns that student as a response.
func getStudentByUSN(c *gin.Context) {
    USN := c.Param("USN")

    // Loop over the list of StudentDB, looking for
    // an USN whose USN value matches the parameter.
    for _, a := range Student {
        if a.USN == USN {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Student not found"})
}


// StudentDB represents data of Students.
type studentDB struct {
    USN  string  `json:"usn"`
    Name  string  `json:"name"`
    Course string  `json:"course"`
    CGPA float64 `json:"cgpa"`
}

// Student details to seed studentDB.
var Student = []studentDB{
    {USN: "1RV19", Name: "Chaitra", Course: "MCA", CGPA: 9.2},
    {USN: "1RV20", Name: "Charana", Course: "BE", CGPA: 8.7},
    {USN: "1RV21", Name: "Manoj", Course: "BCA", CGPA: 7.2},
}

//following sets up an association in which getDetails handles requests to the /Student endpoint path.
func main() {
    router := gin.Default()
    router.GET("/Student", getDetails)
    router.GET("/Student/:USN", getStudentByUSN)
    router.POST("/Student", postDetail)

    router.Run("localhost:8080")
}

// getDetails responds with the list of all details of Students as JSON.
func getDetails(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, Student)
}

