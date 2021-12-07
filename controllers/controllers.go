package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/amshashankk/database"
	"github.com/amshashankk/models"
)

//var db = database.DB

func AddEmployee(c *gin.Context) {
	var db = database.DB

	Emp := new(models.Employee)
	c.BindJSON(&Emp)

	db.Create(&Emp)

	c.JSON(http.StatusOK, gin.H{

		"message": "Employee Added Successfully",
		"data":    Emp,
	})
}

func GetEmployees(c *gin.Context) {
	var db = database.DB

	var Emp []models.Employee

	db.Find(&Emp)

	if Emp != nil {
		c.JSON(http.StatusOK, Emp)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "No employee found"})
	}
}

func GetEmployeeById(c *gin.Context) {
	var db = database.DB

	id := c.Param("id")

	var Emp models.Employee

	db.Find(&Emp, id)

	if Emp.Id != "0" {
		c.JSON(http.StatusOK, Emp)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "No employee found"})
	}

}

func DeleteEmployee(c *gin.Context) {
	var db = database.DB

	id := c.Param("id")

	var Emp models.Employee

	db.Find(&Emp, id)

	if Emp.Id != "0" {
		db.Delete(&Emp)
		c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "No employee found"})
	}

}

func UpdateEmpDetails(c *gin.Context) {
	var db = database.DB

	Emp := new(models.Employee)

	c.BindJSON(&Emp)

	id := c.Param("id")

	var Emp1 models.Employee

	db.First(&Emp1, id)

	Emp1.Name = Emp.Name
	Emp1.Email = Emp.Email
	Emp1.Phone = Emp.Phone

	db.Save(&Emp1)

	c.JSON(http.StatusOK, gin.H{
		"message":  "Employee details updated successfully",
		"Employee": Emp1,
	})

}
