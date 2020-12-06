/*
 * Horton
 * John Shields
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func dbConn() (db *sql.DB) {
	//dbDriver := "mysql"
	//dbUser := "root"
	//dbPass := ""
	//dbName := "repotadb"
	//db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	db, err := sql.Open("mysql", "john:local@tcp(127.0.0.1:3306)/repotadb") // local

	if err != nil {
		panic(err.Error())
	}
	return db
}

// CreateReport - Create a report
func CreateReport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteReport - Delete a Job Report
func DeleteReport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetReport - Get a job report
func GetReport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetReports - List All report
func GetReports(c *gin.Context) {
	db := dbConn()

	selDB, err :=
		db.Query("SELECT * FROM jobReports")
	if err != nil {
		panic(err.Error())
	}

	res := []JobReport{}

	for selDB.Next() {

		var report JobReport

		err = selDB.Scan(&report.JobReportId, &report.WorkerId, &report.WorkDoneId, &report.TimeDate, &report.VehicleModel,
			&report.VehicleReg, &report.VehicleLocation, &report.MilesOnVehicle, &report.Warranty, &report.Breakdown)

		if err != nil {
			panic(err.Error())
		}
		res = append(res, report)
		log.Printf(string(report.JobReportId))
	}

	defer db.Close()

	c.JSON(http.StatusOK, res)
}

// UpdateReport - Update a job report
func UpdateReport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetCustomerReports - Gets all reports on a customer by customer name
func GetCustomerReports(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{})
}