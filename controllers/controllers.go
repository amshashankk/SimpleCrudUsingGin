package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/amshashankk/database"
	"github.com/amshashankk/models"
)

func AddEmployee(c *gin.Context) {

	var Emp models.Employee
	c.BindJSON(&Emp)

	addEmployee := models.Employee{Id: Emp.Id, Phone: Emp.Phone, Name: Emp.Name, Address: Emp.Address}
	database.ConnectDB.Create(&addEmployee)

	c.JSON(http.StatusOK, gin.H{
		"message": "Employee added successfully",
	})

}

func DeleteEmployee(c *gin.Context) {

	var Emp models.Employee
	c.BindJSON(&Emp)

	if err := database.ConnectDB.Where("Id = ?", Emp.Id).Delete(&models.Employee{}).Error; err != nil {
		fmt.Printf("error deleting employee", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func GetEmployeeById(c *gin.Context) {

	var Emp []models.Employee

	if err := database.ConnectDB.Raw("SELECT * from employee where Id = ?").Scan(&Emp).Error; err != nil {
		fmt.Printf("error showing employee details", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if Emp != nil {
		c.JSON(http.StatusOK, Emp)
	} else {
		c.JSON(http.StatusOK, json.RawMessage(`[]`))
	}
}

func UpdateEmpDetails(c *gin.Context) {

	var Emp models.Employee
	c.BindJSON(&Emp)

	if err := database.ConnectDB.Model(&Emp).Where("Id = ?").Updates(models.Employee{Name: Emp.Name, Phone: Emp.Phone, Address: Emp.Address}).Error; err != nil {
		fmt.Printf("error updating details", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
