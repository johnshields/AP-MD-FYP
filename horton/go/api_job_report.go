/*
 * John Shields
 * Horton
 * API version: 1.0.0
 *
 * Job Report API
 * Handles all Job Reports activities - Create, Update, Delete & Getting Reports.
 *
 * References
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 * https://www.golangprograms.com/example-of-golang-crud-using-mysql-from-scratch.html
 * https://levelup.gitconnected.com/build-a-rest-api-using-go-mysql-gorm-and-mux-a02e9a2865ee
 * https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin
 */

package openapi

import (
	"errors"
	"fmt"
	"github.com/GIT_USER_ID/GIT_REPO_ID/go/config"
	"github.com/GIT_USER_ID/GIT_REPO_ID/go/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

// Function that takes tin the input data and calls the function insertJobReport to create a new report.
func CreateReport(c *gin.Context) {
	var report models.JobReport

	// Blind data to object, else throw error
	if err := c.BindJSON(&report); err != nil {
		fmt.Println(err.Error())
	}

	// Insert Job Report details
	if err := insertJobReport(report, wa.Username); err == nil {
		c.JSON(201, "Report created successfully")
	} else {
		c.JSON(401, models.Error{Code: 401, Messages: "Not able to create Report"})
		log.Printf("\n[ALERT] Not completing request.")
	}
}

// Function that creates a new report by starting and committing a MySQL transaction
// to insert into the tables, jobreports and customers.
func insertJobReport(report models.JobReport, username string) error {

	db := config.DbConn()
	fmt.Println("\n[INFO] Processing Report Details...",
		"\nReport Date:", report.Date, "\nCustomer Name:", report.CustomerName)

	// insert into the table jobreports
	insertReport, err := db.Prepare(
		"INSERT INTO jobreports(worker_id, date_stamp, vehicle_model, vehicle_reg, vehicle_location, " +
			"miles_on_vehicle, warranty, breakdown, cause, correction, parts, work_hours, job_report_complete) " +
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

	// insert into the table customers
	insertCustomer, err := db.Prepare("INSERT INTO customers (job_report_id, customer_name, customer_complaint)" +
		" VALUES (LAST_INSERT_ID(), ?, ?)")

	if err != nil {
		log.Println("\n[ALERT] MySQL Error: Error Creating new Report:\n", err)
		return errors.New("error creating Report")
	}

	// Check logged in user
	if !isValidAccount(username) {
		log.Println("\n[ALERT] User has not logged in!")
		return errors.New("error creating Report")
	}

	// begin MySQL transition
	db.Query("BEGIN")
	// execute insert into the table jobreports
	reportResult, err := insertReport.Exec(wa.Id, report.Date, report.VehicleModel, report.VehicleReg, report.VehicleLocation,
		report.MilesOnVehicle, report.Warranty, report.Breakdown, report.Cause, report.Correction, report.Parts,
		report.WorkHours, report.JobComplete)
	// execute insert into the table customers
	customerResult, err := insertCustomer.Exec(report.CustomerName, report.Complaint)
	db.Query("COMMIT") // commit MySQL transition

	if err != nil {
		log.Println("\n[ALERT] MySQL Error: Error Inserting Report Details.\n", err)
		return errors.New("error creating Report")
	}
	fmt.Println("\n[INFO] Print MySQL Results for Report:\n", reportResult, customerResult)

	defer db.Close()
	// Everything is good
	return nil
}

// Function that gets a report in the database from a JOIN QUERY by its requested ID and logged in worker's username.
func GetReportById(c *gin.Context) {

	db := config.DbConn()
	var res []models.JobReport
	worker := wa.Username

	// check for logged in user's cookie
	CheckForCookie(c)

	// Get id from request
	reportId := c.Params.ByName("jobReportId")
	fmt.Printf("Get Report with ID: " + reportId)

	if !isValidAccount(worker) {
		log.Println("\n[ALERT] User has not logged in!")
		c.JSON(401, models.Error{Code: 401, Messages: "User has not logged in!"})
	}

	// JOIN Query to get report by id
	selDB, err := db.Query("SELECT DISTINCT jr.job_report_id, jr.date_stamp, jr.vehicle_model, "+
		"jr.vehicle_reg, jr.miles_on_vehicle, jr.vehicle_location, jr.warranty, jr.breakdown, "+
		"cust.customer_name, cust.customer_complaint, jr.cause, jr.correction, jr.parts, jr.work_hours, "+
		"wkr.worker_name, jr.job_report_complete FROM jobreports jr INNER JOIN customers cust "+
		"ON jr.job_report_id = cust.job_report_id "+
		"INNER JOIN workers wkr ON jr.worker_id = wkr.worker_id "+
		"WHERE jr.job_report_id = ? AND wkr.username = ?", reportId, worker)

	fmt.Println("\n[INFO] Processing Report...")

	if err != nil {
		log.Println("\n[ALERT] Failed to process reports! \n500 Internal Server Error.")
		c.JSON(500, nil)
	}

	fmt.Println("\n[INFO] Loading model...")

	// Run through each record and read values
	for selDB.Next() {
		var report models.JobReport

		err = selDB.Scan(&report.JobReportId, &report.Date, &report.VehicleModel, &report.VehicleReg, &report.MilesOnVehicle,
			&report.VehicleLocation, &report.Warranty, &report.Breakdown, &report.CustomerName, &report.Complaint, &report.Cause,
			&report.Correction, &report.Parts, &report.WorkHours, &report.WorkerName, &report.JobComplete)

		if err != nil {
			// return user friendly message to client
			log.Println("\n[ALERT] Failed to load model! \n500 Internal Server Error.")
			c.JSON(500, nil)
		}
		// Add each record to array
		res = append(res, report)
		log.Printf(string(report.JobReportId))
	}
	// Return values, Status OK
	c.JSON(http.StatusOK, res)

	fmt.Println("\n[INFO] Report Processed.", res)
	defer db.Close()
}

// Function that gets all report in the database from a JOIN QUERY by a logged in worker's username.
func GetReports(c *gin.Context) {

	db := config.DbConn()

	worker := wa.Username
	var res []models.JobReport
	var report models.JobReport

	if !isValidAccount(worker) {
		log.Println("\n[ALERT] User is not logged in!")
		c.JSON(401, models.Error{Code: 401, Messages: "User is not logged in!"})
	}

	CheckForCookie(c)

	// JOIN Query to get worker's job reports
	selDB, err := db.Query("SELECT DISTINCT jr.job_report_id, jr.date_stamp, jr.vehicle_model, "+
		"jr.vehicle_reg, jr.miles_on_vehicle, jr.vehicle_location, jr.warranty, jr.breakdown, "+
		"cust.customer_name, cust.customer_complaint, jr.cause, jr.correction, jr.parts, jr.work_hours, "+
		"wkr.worker_name, jr.job_report_complete FROM jobreports jr INNER JOIN customers cust "+
		"ON jr.job_report_id = cust.job_report_id "+
		"INNER JOIN workers wkr ON jr.worker_id = wkr.worker_id WHERE wkr.username = ?", worker)

	fmt.Println("\n[INFO] Processing Reports...")

	if err != nil {
		log.Println("\n[ALERT] Failed to process reports! \n500 Internal Server Error.")
		c.JSON(500, nil)
	}

	// Run through each record and read values
	for selDB.Next() {

		err = selDB.Scan(&report.JobReportId, &report.Date, &report.VehicleModel, &report.VehicleReg, &report.MilesOnVehicle,
			&report.VehicleLocation, &report.Warranty, &report.Breakdown, &report.CustomerName, &report.Complaint, &report.Cause,
			&report.Correction, &report.Parts, &report.WorkHours, &report.WorkerName, &report.JobComplete)

		if err != nil {
			log.Println("\n[ALERT] Failed to load model! \n500 Internal Server Error.")
			c.JSON(500, nil)
		}
		// Add each record to array
		res = append(res, report)
		log.Printf(string(report.JobReportId))
	}
	// Return values, Status OK
	c.JSON(http.StatusOK, res)

	fmt.Println("\n[INFO] Reports Processed.", res)
	defer db.Close()
}

// Function to update/edit report by its requested ID.
func UpdateReport(c *gin.Context) {
	db := config.DbConn()
	var report models.JobReport

	// Get id from request
	reportId := c.Params.ByName("jobReportId")
	fmt.Printf("Get Report with ID: " + reportId)

	CheckForCookie(c)

	// Blind data to object, else throw error
	if err := c.BindJSON(&report); err != nil {
		fmt.Println(err.Error())
	}

	// Read in values from client request and build object
	update, err := db.Exec("UPDATE jobreports jr SET jr.date_stamp = ?, jr.vehicle_model = ?, "+
		"jr.vehicle_reg = ?, jr.vehicle_location = ?, jr.miles_on_vehicle = ?, jr.warranty = ?, "+
		"jr.breakdown = ?, jr.cause = ?, jr.correction = ?, jr.parts = ?, jr.work_hours = ?, "+
		"jr.job_report_complete = ? WHERE jr.job_report_id = ?", report.Date, report.VehicleModel, report.VehicleReg, report.VehicleLocation,
		report.MilesOnVehicle, report.Warranty, report.Breakdown, report.Cause, report.Correction, report.Parts,
		report.WorkHours, report.JobComplete, reportId)

	if err != nil {
		log.Println("\n[ALERT] MySQL Error: Error Updating Report:\n", err)
		c.JSON(503, models.Error{Code: 503, Messages: "Error Updating Report"})
	} else {
		fmt.Println("\n[INFO] Processing Job Report Details...",
			"\nReport Number:", report.JobReportId, "\nReport Date:", report.Date)

		// Return 202 response for "Updated"
		c.JSON(202, gin.H{})
		fmt.Println("\n[INFO] Print MySQL Results for Report:\n", update)
		defer db.Close()
	}
}

// Function to delete a report by its requested ID.
func DeleteReport(c *gin.Context) {
	db := config.DbConn()

	// Get id from request
	reportId := c.Params.ByName("jobReportId")
	fmt.Printf("Get Report with ID: " + reportId)

	CheckForCookie(c)

	// Create query
	res, err := db.Exec("DELETE FROM jobreports WHERE job_report_id=?", reportId)

	if err != nil {
		fmt.Printf("500 Internal Server Error.")
		c.JSON(500, nil)
	}

	affectedRows, err := res.RowsAffected()

	if err != nil {
		fmt.Printf("[ALERT] Report not deleted.")
		// return user friendly message to client
		c.JSON(500, models.Error{Code: 500, Messages: "Report not deleted"})
	}

	fmt.Printf("The statement affected %d rows\n", affectedRows)

	c.JSON(204, nil)
}
