package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

type Person struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
  Address string  `json:"address"`
	Age int    `json:"age"`
}

var persons = []Person{
	{ID: "1", Name  : "Queen", Address: "Calaca", Age: 50},
	{ID: "2", Name  : "King", Address: "Lemery", Age: 52},
	{ID: "3", Name  : "Pawn", Address: "Taal", Age: 30},
}

func getPersons(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, persons)
}

func PersonById(c *gin.Context) {
	id := c.Param("id")
	Person, err := getPersonById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, person)
}


func createPerson(c *gin.Context) {
	var newPerson person

	if err := c.BindJSON(&newPersons); err != nil {
		return
	}

	persons = append(personss, newPerson)
	c.IndentedJSON(http.StatusCreated, newPerson)
}

func DeletePerson(c *gin.Context) {

	id := c.Param("id")
	
	err := DeletePerson(&person, id)
	
}




func main() {
	router := gin.Default()
	router.GET("/persons", getPersons)
	router.GET("/persons/:id", PersonById)
	router.POST("/persons", createPerson)
	router.DELETE("/persons", DeletePerson)
	router.Run("localhost:8080")
}
